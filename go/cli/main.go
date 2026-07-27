package main

import (
	"fmt"
	"os"
)

func main() {
	// 1: means to give me everything from index 1 onwards
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage: todo <command>")
		return
	}
	command := args[0]

	switch command {
	case "add":
		addTask(args)
	case "list":
		listTasks()
	case "done":
		doneTask(args)
	case "delete":
		deleteTask(args)
	default:
		fmt.Println("Unknown command:", command)
	}
}
