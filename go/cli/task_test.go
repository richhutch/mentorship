package main

import "testing"

func TestNewTask(t *testing.T) {
	task := NewTask(1, "Buy groceries", "2026-06-10")

	if task.ID != 1 {
		t.Errorf("expected ID 1, got %d", task.ID)
	}
	if task.Title != "Buy groceries" {
		t.Errorf("expected title 'Buy groceries', got %s", task.Title)
	}
	if task.Done != false {
		t.Errorf("expected Done to be false, got %v", task.Done)
	}
}
