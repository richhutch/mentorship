package main

import (
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	deadline, err := time.Parse(dateLayout, "2026-06-10")
	if err != nil {
		t.Fatalf("failed to parse test deadline: %v", err)
	}

	task := NewTask(1, "Buy groceries", deadline)

	if task.ID != 1 {
		t.Errorf("expected ID 1, got %d", task.ID)
	}
	if task.Title != "Buy groceries" {
		t.Errorf("expected title 'Buy groceries', got %s", task.Title)
	}
	if task.Done != false {
		t.Errorf("expected Done to be false, got %v", task.Done)
	}
	if !task.Deadline.Equal(deadline) {
		t.Errorf("expected deadline %v, got %v", deadline, task.Deadline)
	}
}
