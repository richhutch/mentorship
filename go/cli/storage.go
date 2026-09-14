package main

import (
	"encoding/json"
	"os"
)

// Storage defines how tasks are loaded and saved, so the
// task logic doesn't need to know whether tasks live in a
// file, memory, or a database.
type Storage interface {
	LoadTasks() ([]Task, error)
	SaveTasks(tasks []Task) error
}

// FileStore implements Storage using a JSON file on disk.
type FileStore struct {
	Path string
}

func NewFileStore(path string) *FileStore {
	return &FileStore{Path: path}
}

func (f *FileStore) LoadTasks() ([]Task, error) {
	data, err := os.ReadFile(f.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (f *FileStore) SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(f.Path, data, 0644)
}