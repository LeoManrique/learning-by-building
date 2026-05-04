package main

import (
	"cli-task-tracker/model"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./db/tasks.json")
	fmt.Printf("data=%q \nerr=%v\n", data, err)

	var tasks []model.Task
	uerr := json.Unmarshal(data, &tasks)
	fmt.Printf("tasks=%+v uerr=%v\n", tasks, uerr)
}
