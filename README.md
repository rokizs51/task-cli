Here's a simple, natural-looking README template for your task CLI:

  # task-cli

  A simple command-line task manager written in Go. https://github.com/rokizs51/task-cli

  ## Install

  ```bash
  go build -o task .

  Usage

  # Add a task
  task-cli add "Buy groceries"

  # List all tasks
  task-cli list

  # List tasks by status
  task-cli list done
  task-cli list todo
  task-cli list in-progress

  # Update a task
  task-cli update 0 "Buy groceries and cook dinner"

  # Mark task as in-progress
  task-cli mark-in-progress 0

  # Mark task as done
  task-cli mark-done 0

  # Delete a task
  task-cli delete 0

  Tasks stored in

  Tasks are saved in task.json in the current directory.

  Requirements

  - Go 1.21+