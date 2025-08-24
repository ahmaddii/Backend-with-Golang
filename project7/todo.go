package main

import "fmt"

var tasks [100]string // fixed array for tasks
var taskCount int = 0 // number of current tasks

func main() {
	for {
		var userInput int

		fmt.Println("-------Welcome to To Do List App------")
		fmt.Println("1 . Add Task")
		fmt.Println("2 . Remove Task")
		fmt.Println("3 . View Tasks")
		fmt.Println("4 . Exit")

		fmt.Print("Enter choice: ")
		fmt.Scanln(&userInput)

		if userInput == 1 {
			addTask()
		} else if userInput == 2 {
			removeTask()
		} else if userInput == 3 {
			viewTasks()
		} else if userInput == 4 {
			fmt.Println("Exiting... Goodbye!")
			break
		} else {
			fmt.Println("Invalid choice! Please try again.")
		}
	}
}

// Add Task
func addTask() {
	if taskCount >= 100 {
		fmt.Println("Task list is full!")
		return
	}

	fmt.Print("Enter the Task to add: ")
	fmt.Scanln(&tasks[taskCount])
	taskCount++
	fmt.Println("Task Added Successfully ✅")
}

// Remove Task
func removeTask() {
	if taskCount == 0 {
		fmt.Println("No tasks to remove!")
		return
	}

	viewTasks()

	var choice int
	fmt.Print("Enter the task number to remove: ")
	fmt.Scanln(&choice)

	if choice < 1 || choice > taskCount {
		fmt.Println("Invalid task number!")
		return
	}

	// Shift tasks to left
	for i := choice - 1; i < taskCount-1; i++ {
		tasks[i] = tasks[i+1]
	}

	taskCount--
	fmt.Println("Task Removed Successfully ❌")
}

// View Tasks
func viewTasks() {
	if taskCount == 0 {
		fmt.Println("No tasks found! Add some tasks first.")
		return
	}

	fmt.Println("Your Tasks:")
	for i := 0; i < taskCount; i++ {
		fmt.Printf("%d. %s\n", i+1, tasks[i])
	}
}
