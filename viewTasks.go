package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func viewTasks() {
	f, err := os.Open("todos.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	var index int = 0

	// Loop to read each task
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading task. %v\n", err)
			return
		}
		index = index + 1
		fmt.Printf("%d. %s", index, line)
	}
}
