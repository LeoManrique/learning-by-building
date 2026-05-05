package main

import (
	"cli-task-tracker/model"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./db/tasks.json")
	if err != nil {
		fmt.Printf("layer 0: %T = %v\n", err, err)
		inner := errors.Unwrap(err)
		fmt.Printf("layer 1: %T = %v\n", inner, inner)
		fmt.Printf("layer 2: %T = %v\n", errors.Unwrap(inner), errors.Unwrap(inner))
	}
	fmt.Printf("data=%q \nerr=%v\n", data, err)

	var tasks []model.Task
	uerr := json.Unmarshal(data, &tasks)
	fmt.Printf("tasks=%+v uerr=%v\n", tasks, uerr)
}
