# GitHub User Activity

A simple Go CLI tool to fetch and display GitHub user events such as pushes, repository creations, stars, and deletions.
Project url: https://roadmap.sh/projects/github-user-activity

---

## Overview

This tool fetches recent GitHub events for a given username using the [GitHub API](https://docs.github.com/en/rest/activity/events) and displays them in a readable format. It is written in Go and demonstrates basic HTTP requests, JSON parsing, and CLI interaction.

---

## Features

- Fetch and display GitHub user events.
- Supports the following event types:
  - **PushEvent**: Pushed commits to a repository.
  - **CreateEvent**: Created a repository.
  - **WatchEvent**: Starred a repository.
  - **DeleteEvent**: Deleted a branch or tag.
---

## Installation

### Prerequisites
- Go 1.21 or higher installed on your machine.

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/trinverdale-bit/github-user-activity.git
   cd github-user-activity
2. Build the project:
   `go build -o github-user-activity main.go`
3. Run ./github-user-activity and specify a user
  `./github-user-activity trinverdale-bit`
  Output:
  ```
- Pushed commit to trinverdale-bit/github-user-activity
- Pushed commit to trinverdale-bit/github-user-activity
- Starred the repository neovim/neovim at 2025-01-20T10:23:31Z
- Pushed commit to trinverdale-bit/github-user-activity
- Pushed commit to trinverdale-bit/task-tracker
- Created the repository at 2025-01-20T01:24:39Z
- Created the repository at 2025-01-20T01:24:39Z
- Pushed commit to trinverdale-bit/task-tracker
- Pushed commit to trinverdale-bit/task-tracker
- Pushed commit to trinverdale-bit/task-tracker
- Pushed commit to trinverdale-bit/task-tracker
- Pushed commit to trinverdale-bit/task-tracker
- Pushed commit to trinverdale-bit/task-tracker
- Created the repository at 2025-01-19T04:27:58Z
- Created the repository at 2025-01-19T04:27:58Z
- Pushed commit to trinverdale-bit/task-tracker
- Pushed commit to trinverdale-bit/task-tracker
- Created the repository at 2024-12-31T04:17:15Z
- Created the repository at 2024-12-31T04:17:15Z
- Created the repository at 2024-12-31T04:16:39Z
- Deleted branch/tag trinverdale-bit/task-tracker at 2024-12-31T04:08:45Z
- Created the repository at 2024-12-31T04:08:44Z
- Created the repository at 2024-12-31T03:54:54Z
- Created the repository at 2024-12-31T03:54:54Z
```
