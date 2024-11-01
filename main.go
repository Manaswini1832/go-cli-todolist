package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Select an option : \n[A] Add Task\n[B] View Tasks\n[C] Delete a Task\n")

	reader := bufio.NewReader(os.Stdin)
	option, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	option = strings.TrimSpace(option)
	switch option {
	case "A":
		addTask()
	case "B":
		viewTasks()
	case "C":
		deleteTask()
	default:
		fmt.Println("Invalid option. Try again!")
	}

}
