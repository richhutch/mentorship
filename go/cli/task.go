package main

import "time"

type Task struct {
	ID       int
	Title    string
	Done     bool
	Deadline time.Time
}

func NewTask(ID int, Title string, Deadline time.Time) Task {
	task := Task{
		ID:       ID,
		Title:    Title,
		Done:     false,
		Deadline: Deadline,
	}
	return task
}
