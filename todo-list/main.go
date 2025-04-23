package main

import (
	"fmt"
	"os"

	"github.com/locngoduc/Golang/todo-list/todo"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: todo add|list|done <task>")
		return
	}

	switch args[1] {
	case "add":
		todo.AddTask(args[2])
		break
	case "list":
		todo.ListTask()
		break
	case "done":
		todo.CompleteTask(args[2])
		break
	}
}
