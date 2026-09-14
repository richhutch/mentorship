package main

import (
	"fmt"
	"os"
)

const dataFile = "tasks.json"

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage: todo <command>")
		fmt.Println("Available commands:")
		fmt.Println("  add    - add a new task")
		fmt.Println("  list   - list all tasks")
		fmt.Println("  done   - mark a task as done")
		fmt.Println("  delete - delete a task")
		os.Exit(1)
	}
	command := args[0]

	store := NewFileStore(dataFile)

	switch command {
	case "add":
		addTask(args, store)
	case "list":
		if err := listTasks(store); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "done":
		doneTask(args, store)
	case "delete":
		deleteTask(args, store)
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}