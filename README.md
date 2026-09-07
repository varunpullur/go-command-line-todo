# 📝 Todo CLI

A clean, modular command-line todo manager written in Go — built around a swappable storage layer, so today it's in-memory and tomorrow it could be a file, a database, or anything else, with zero changes to business logic.

---

## ✨ Features

- ➕ Add tasks
- 📋 List all tasks
- ✅ Mark tasks as done
- ✏️ Update task titles
- 🗑️ Delete tasks
- 💾 Pluggable storage backend (in-memory now, file-based coming soon)

---

## 🏗️ Architecture

This project follows a layered architecture with strict separation of concerns — each package has exactly one job.

```
User Input (stdin)
      │
      ▼
┌──────────────┐      ┌─────────────────┐      ┌─────────────┐
│   cli.App    │ ───▶ │   commands      │ ───▶ │    store    │
│ (orchestrate)│      │ (business logic)│      │(persistence)│
└──────────────┘      └─────────────────┘      └─────────────┘
      │                                            │
      ▼                                            ▼
┌─────────────┐                            ┌───────────────┐
│      ui     │                            │      task     │
│  (display)  │                            │ (domain model)│
└─────────────┘                            └───────────────┘
```

| Layer        | Responsibility                                          | Does NOT do                             |
| ------------ | ------------------------------------------------------- | --------------------------------------- |
| `cli.App`  | Owns the input loop, reads user choices, calls commands | Business logic, persistence, formatting |
| `commands` | Use-case logic — get/mutate/save flows                 | Printing, reading stdin                 |
| `store`    | Persistence via the`Store` interface                  | Business rules, formatting              |
| `task`     | Domain model — what a`Task` *is*                   | I/O, storage, display                   |
| `ui`       | All printed output — menus, results, banners           | Reading input, orchestration            |

---

## 📂 Project Structure

```
todo-cli/
├── cmd/
│   └── main.go              # Entry point — wiring only
│
├── internal/
│   ├── task/
│   │   └── task.go          # Task struct + domain rules
│   │
│   ├── store/
│   │   ├── store.go         # Store interface (the storage contract)
│   │   └── memory_store.go  # In-memory implementation
│   │
│   ├── cli/
│   │   ├── app.go           # Main loop + action flows
│   │   ├── reader.go        # Shared stdin reader (ReadLine / ReadInt)
│   │   └── commands/
│   │       └── commands.go  # Business logic connecting cli ↔ store
│   │
│   └── ui/
│       ├── printer.go       # Formats Task / []Task for display
│       └── prompts.go       # Static menu banners
│
└── go.mod
```

---

## 🔌 The Storage Contract

All storage backends implement a single interface, so swapping `MemoryStore` for a future `FileStore` requires **no changes** anywhere else in the codebase:

```go
type Store interface {
    Add(t task.Task) (task.Task, error)
    Get(id int) (task.Task, error)
    List() ([]task.Task, error)
    Update(t task.Task) error
    Delete(id int) error
}
```

> 🔜 **Coming soon:** `FileStore` — persists tasks to disk using the same interface, selectable from the startup menu.

---

## 🚀 Getting Started

### Prerequisites

- Go 1.21+ installed

### Run it

```bash
git clone <your-repo-url>
cd todo-cli
go run ./cmd
```

### Build a binary

```bash
go build -o todo ./cmd
./todo
```

---

## 🎮 Usage

On startup, choose a storage backend:

```
=====================================
        Welcome to Todo CLI
=====================================
Choose storage type:
  1. In-Memory (data lost on exit)
  2. File-based (coming soon)
Enter your choice:
```

Then manage your tasks:

```
=====================================
              Todo CLI
=====================================
  1. Add Task
  2. List Tasks
  3. Mark Task as Done
  4. Update Task
  5. Delete Task
  6. Exit
Enter your choice:
```

---

## 🧪 Testing

```bash
go test ./...
```

Because `commands` and `store` never touch stdin/stdout directly, both layers are fully unit-testable without mocking terminal I/O.
