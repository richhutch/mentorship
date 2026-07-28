package main

import (
	"encoding/json"
	"os"
)

const dataFile = "tasks.json"

// "LoadTasks = grab what's on disk, empty list if nothing there. Save = write current list to disk. Always load before modifying

func LoadTasks() ([]Task, error) {
	data, err := os.ReadFile(dataFile)
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

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dataFile, data, 0644)
}
