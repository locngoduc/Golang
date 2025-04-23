package todo

import (
	"fmt"
	"strconv"
)

type Task struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

var Tasks []Task

func AddTask(desc string) {
	Tasks = append(Tasks, Task{Description: desc, Done: false})
	fmt.Println("Added task:", desc)
}

func ListTask() {
	LoadTasks()
	for i, t := range Tasks {
		status := "Success"
		if !t.Done {
			status = "Fail"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, t.Description, status)
	}
}

func CompleteTask(id string) {
	LoadTasks()
	idx, _ := strconv.Atoi(id)
	if idx < 1 || idx > len(Tasks) {
		fmt.Println("Invalid task ID")
		return
	}
	Tasks[idx-1].Done = true
	SaveTasks()
	fmt.Println("Task completed:", Tasks[idx-1].Description)
}
