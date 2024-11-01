package main

import (
	"bufio"
	"fmt"
	"os"
)

func addTask() {
	fmt.Println("Add a task : ")
	reader := bufio.NewReader(os.Stdin)
	task, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading task : ", err)
		return
	}
	f, err := os.OpenFile("todos.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.WriteString(task); err != nil {
		panic(err)
	}
	fmt.Printf("Added task to todolist\n")

}
