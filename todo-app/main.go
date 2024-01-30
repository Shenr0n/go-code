package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Text string
	Done bool
}

func main() {
	tasks := []Task{}

	for {
		showMenu()
		option := getUserInput("Enter your choice: ")

		switch option {
		case "1":
			showTasks(tasks)
		case "2":
			addTask(&tasks)
		case "3":
			markTaskDone(&tasks)
		case "4":
			saveTaskToFile(tasks)
		case "5":
			fmt.Println("Exiting the To-do Appplication")
			return
		default:
			fmt.Println("Invalid choice. Try again")
		}
	}
}

func showMenu() {
	fmt.Println("\nMenu:")
	fmt.Println("1. Show Tasks")
	fmt.Println("2. Add a Task")
	fmt.Println("3. Mark Task as Done")
	fmt.Println("4. Save Tasks to a File")
	fmt.Println("5. Exit")
}

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func showTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No Tasks available")
		return
	}

	fmt.Println("Tasks: ")
	for i, task := range tasks {
		status := " "
		if task.Done {
			status = "x"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Text)
	}
}

func addTask(tasks *[]Task) {
	taskText := getUserInput("Enter task: ")
	*tasks = append(*tasks, Task{Text: taskText})
	fmt.Println("Task added.")
}

func markTaskDone(tasks *[]Task) {
	showTasks(*tasks)
	taskIndexStr := getUserInput("Enter the task number to mark as done: ")
	taskIndex, err := parseTaskIndex(taskIndexStr)
	if err != nil || taskIndex < 1 || taskIndex > len(*tasks) {
		fmt.Println("Invalid task number. Try again.")
	}
	(*tasks)[taskIndex-1].Done = true
	fmt.Println("Task marked as done")
}

func parseTaskIndex(input string) (int, error) {
	return strconv.Atoi(input)
}

func saveTaskToFile(tasks []Task) {
	file, err := os.Create("tasks.txt")
	if err != nil {
		fmt.Println("Error creating tasks file: ", err)
		return
	}
	defer file.Close()

	for _, task := range tasks {
		status := " "
		if task.Done {
			status = "x"
		}
		file.WriteString(fmt.Sprintf("[%s] %s\n", status, task.Text))
	}
	fmt.Println("Tasks saved to file 'tasks.txt'")
}
