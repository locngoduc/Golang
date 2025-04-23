package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadTasks() {
	file, err := os.ReadFile("./tasks.json")
	if err != nil {
		fmt.Println("Error reading tasks file:", err)
		return
	}
	json.Unmarshal(file, &Tasks)
}

func SaveTasks() {
	data, err := json.Marshal(Tasks)
	if err != nil {
		fmt.Println("Error marshalling tasks:", err)
		return
	}
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing tasks file:", err)
		return
	}
	fmt.Println("Tasks saved successfully.")
}
