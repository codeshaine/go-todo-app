# CLI TODO Application

A simple command-line TODO application written in Go. This application allows you to manage a list of tasks with functionalities to add, complete, incomplete, delete, and list tasks.

## Features

- **Add a new TODO**
- **Mark TODOs as complete or incomplete**
- **Delete TODOs**
- **List all TODOs**

## Installation

To run this application, you need to have Go installed. You can download and install Go from the [official Go website](https://golang.org/dl/).

Clone the repository and navigate into the project directory:

```bash
git clone https://github.com/codeshaine/go-todo-app.git
cd go-todo-app
```

## Building

To build the application, run the following command:

```bash
go build -o todo
```

This will compile the source code into an executable named `todo` (or `todo.exe` on Windows) in the current directory.

## Usage

1. **Add a new TODO**

   To add a new TODO, use the `-add` flag followed by the task description. For example:

   ```bash
   ./todo -add "Buy groceries"
   ```

2. **Mark a TODO as complete**

   To mark a TODO as complete, use the `-complete` flag followed by the task index. For example:

   ```bash
   ./todo -complete 1
   ```

3. **Mark a TODO as incomplete**

   To mark a TODO as incomplete, use the `-incomplete` flag followed by the task index. For example:

   ```bash
   ./todo -incomplete 1
   ```

4. **Delete a TODO**

   To delete a TODO, use the `-delete` flag followed by the task index. For example:

   ```bash
   ./todo -delete 1
   ```

5. **List all TODOs**

   To list all TODOs, use the `-list` flag:

   ```bash
   ./todo -list
   ```

## Example

### Adding a TODO

```bash
$ ./todo -add "Complete Go project"
```

### Listing TODOs

```bash
$ ./todo -list
╔═══╤═══════════════════════════════╤═══════╤═══════════════════╤═══════════════════╗
║ # │             Task              │ Done? │     CreatedAt     │    CompletedAt    ║
╟━━━┼━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━╢
║ 1 │ Complete Go project           │  No   │ 07 Sep 2024 21:34 │         -         ║
╟━━━┼━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━╢
║                       You’ve got 1 task left. Keep going!                        ║
╚═══╧═══════════════════════════════╧═══════╧═══════════════════╧═══════════════════╝
```
