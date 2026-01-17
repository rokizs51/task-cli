package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Task struct  {
	ID int `json:"id"`
	Description string `json:"description"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func WriteTask(filename string,  task string) error{
	tasks, err := ReadFile(filename)
	if err != nil {
		return err
	}
	var newTask Task
	tasklLargestId := -1
	for _, task := range tasks {
		if task.ID > tasklLargestId {
			tasklLargestId = task.ID
		}
	}
	newTask.ID = tasklLargestId +1
	newTask.Description = task
	newTask.Status = "todo"
	newTask.CreatedAt = time.Now()
	newTask.UpdatedAt = time.Now()

	tasks = append(tasks, newTask)

	data ,err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
func UpdateTask(filename string, taskID int, description string) error {
	tasks, err := ReadFile(filename)
	if err != nil {
		return err
	}
	foundTask := false
	for idx, task := range tasks {
		if task.ID == taskID {
			tasks[idx].Description = description
			tasks[idx].UpdatedAt = time.Now()
			foundTask = true
			break
		}
	}
	if !foundTask {
		return errors.New("task does not exist")
	}
	
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(filename, data, 0644)

}

func UpdateTaskStatus(filename string, taskID int, status string) error {
	tasks, err := ReadFile(filename)
	if err != nil {
		return err
	}
	foundTask := false
	for idx, task := range tasks {
		if task.ID == taskID {
			tasks[idx].Status = status
			tasks[idx].UpdatedAt = time.Now()
			foundTask = true
			break
		}
	}
	if !foundTask{
		return errors.New("task doesnt exist")
	}
	
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(filename, data, 0644)

}

func RemoveTask(filename string, taskID int) error {
	tasks, err := ReadFile(filename)
	if err != nil {
		return err
	}
	foundTask := false
	for idx, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:idx], tasks[idx+1:]... )
			foundTask = true
			break
		}
	}
	
	if !foundTask {
		return	errors.New("task does not exist")
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(filename, data, 0644)
}

func ReadFile(filename string) ([]Task, error){
	data, err := os.ReadFile(filename)	
	if err != nil {
		if err.Error() == "open task.json: no such file or directory"{
			return []Task{}, nil
		}
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, nil
}

func PrintTask(status string) error {
	tasks, err := ReadFile("task.json")
	if err != nil {
		return err
	}
	var filteredTask []Task
	if status != ""{
		for _, task := range tasks {
			if task.Status == status {
				filteredTask =append(filteredTask, task)
			}
		}
	}else {
		filteredTask = tasks
	}
	fmt.Printf("%-4s %-20s %-10s %-20s %-20s\n", "ID", "Description", "Status", "Created", "Updated")
	fmt.Println("-----------------------------------------------------------------------------------")
	for _, task := range filteredTask {
		fmt.Printf("%-4d %-20s %-10s %-20s %-20s\n",
			task.ID,
			task.Description,
			task.Status,
			task.CreatedAt.Format("2006-01-02 15:04"),
			task.UpdatedAt.Format("2006-01-02 15:04"),
		)
	}
	return nil
}
