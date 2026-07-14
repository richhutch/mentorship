package main

import "fmt"

func addTask(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: todo add <title> <deadline>")
		return
	}
	title := args[1]
	deadline := args[2]

	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	task := NewTask(1, title, deadline)
	tasks = append(tasks, task)

	err = SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving task:", err)
		return
	}

	fmt.Println("Added task:", task.Title, "| Deadline:", task.Deadline)
}

func listTasks() {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	if len(tasks) == 0 {
		fmt.Println("No tasks yet.")
		return
	}
	for _, task := range tasks {
		status := " "
		if task.Done {
			status = "x"
		}
		fmt.Printf("[%s] %d: %s (due %s)\n", status, task.ID, task.Title, task.Deadline)
	}
}
