package main

type Task struct {
	ID       int
	Title    string
	Done     bool
	Deadline string
}

func NewTask(ID int, Title string, Deadline string) Task {
	task := Task{
		ID:       ID,
		Title:    Title,
		Done:     false,
		Deadline: Deadline,
	}
	return task
}
