
# 📝 Go To-Do List CLI App

A simple command-line To-Do List application written in Go. This app allows you to **view**, **add**, and **delete** tasks interactively via the terminal.

---

## 📁 Project Structure

```
.
├── main.go               # Main application logic
└── data/
    └── task.go           # Task struct definition and helper methods
```

---

## 🔧 Features

- View current tasks
- Add new tasks
- Delete tasks by ID
- Clean CLI interface using standard Go packages

---

## ▶️ Getting Started

### Prerequisites

- Go 1.18 or later installed

### Run the App

```bash
go run main.go
```

---

## 🖥️ Sample Usage

```
Your Today's To-Do List
Select below options
0 to view TASK
1 to add TASK
2 to delete TASK
3 for exit
```

Add a task:

```
Enter Your Task Name:
Buy groceries
```

Delete a task:

```
Enter Id to delete a task
1
```

---

## 📦 Code Overview

### `main.go`

Handles user interaction:

- `printToDOList(list)` — displays current tasks
- `delTask(list, id)` — deletes a task from the list
- Loop with options to view, add, delete, or exit

### `data/task.go`

Defines the `Task` type:

```go
package data

import "fmt"

type Task struct {
    Id   int
    Name string
}

func (task Task) String() string {
    return fmt.Sprintf("%d ---  %v", task.Id, task.Name)
}

func NewTask(id int, name string) Task {
    return Task{Id: id, Name: name}
}
```

- Implements a `String()` method for clean display
- `NewTask(id, name)` is a constructor for new tasks

---

## ❗ Notes

- The to-do list is **stored in memory** only. Closing the app will erase all tasks.
- No concurrency or persistence (e.g., file/DB storage) is included in this version.

---
