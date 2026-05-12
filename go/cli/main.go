package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage: todo <command>")
		return
	}
	command := args[0]

	switch command {
	case "add":
		fmt.Println("add command")
	case "list":
		fmt.Println("list command")
	case "done":
		fmt.Println("done command")
	case "delete":
		fmt.Println("delete command")
	default:
		fmt.Println("Unknown command:", command)
	}
}
