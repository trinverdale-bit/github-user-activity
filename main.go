package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type GitHubEvent struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Actor struct {
		Login string `json:"login"`
	} `json:"actor"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	CreatedAt string `json:"created_at"`
}

var username string = os.Args[1]

func main() {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
        fmt.Println("Error: Received status code:", response.StatusCode, "check username")
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var events []GitHubEvent
	err = json.Unmarshal(body, &events)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	for _, event := range events {
        if event.Type == "PushEvent" {
            fmt.Printf("- Pushed commit to %s", event.Repo.Name)
            fmt.Println()
        } else if event.Type == "CreateEvent" {
            fmt.Printf("- Created the repository at %s", event.CreatedAt)
            fmt.Println()
        } else if event.Type == "WatchEvent" {
            fmt.Printf("- Starred the repository %s at %s", event.Repo.Name, event.CreatedAt)
            fmt.Println()
        } else if event.Type == "DeleteEvent" {
            fmt.Printf("- Deleted branch/tag %s at %s", event.Repo.Name, event.CreatedAt)
            fmt.Println()
        }
	}
}
