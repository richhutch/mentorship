package main

import (
	"fmt"
	"strconv"
	"time"
)

const dateLayout = "2006-01-02"

func addTask(args []string, storage Storage) {
	if len(args) < 3 {
		fmt.Println("Usage: todo add <title> <deadline>")
		return
	}
	title := args[1]

	deadline, err := time.Parse(dateLayout, args[2])
	if err != nil {
		fmt.Println("Invalid deadline format, expected YYYY-MM-DD:", err)
		return
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	task := NewTask(1, title, deadline)
	tasks = append(tasks, task)

	err = storage.SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving task:", err)
		return
	}

	fmt.Println("Added task:", task.Title, "| Deadline:", task.Deadline.Format(dateLayout))
}

func listTasks(storage Storage) error {
	tasks, err := storage.LoadTasks()
	if err != nil {
		return fmt.Errorf("error loading tasks: %w", err)
	}
	if len(tasks) == 0 {
		fmt.Println("No tasks yet.")
		return nil
	}
	for _, task := range tasks {
		status := " "
		if task.Done {
			status = "x"
		}
		fmt.Printf("[%s] %d: %s (due %s)\n", status, task.ID, task.Title, task.Deadline.Format(dateLayout))
	}
	return nil
}

func doneTask(args []string, storage Storage) {
	if len(args) < 2 {
		fmt.Println("Usage: todo done <id>")
		return
	}

	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid ID:", args[1])
		return
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			found = true
			break
		}
	}

	if !found {
		fmt.Println("No task found with ID:", id)
		return
	}

	err = storage.SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("Marked task", id, "as done")
}

func deleteTask(args []string, storage Storage) {
	if len(args) < 2 {
		fmt.Println("Usage: todo delete <id>")
		return
	}

	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid ID:", args[1])
		return
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	found := false
	newTasks := []Task{}
	for _, task := range tasks {
		if task.ID == id {
			found = true
			continue
		}
		newTasks = append(newTasks, task)
	}

	if !found {
		fmt.Println("No task found with ID:", id)
		return
	}

	err = storage.SaveTasks(newTasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("Deleted task", id)
}
