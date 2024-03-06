package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	name        string
	description string
	complete    bool
}

type ListToDo struct {
	tasks []Task
}

// Add task
func (l *ListToDo) addTask(t Task) {

	l.tasks = append(l.tasks, t)
}

// Mark if was done

func (l *ListToDo) MarkDone(index int) {

	l.tasks[index].complete = true

}

//Update task

func (l *ListToDo) UpdateTask(index int, t Task) {
	l.tasks[index] = t

}

// Delete task
func (l *ListToDo) deleteTask(index int) {

	l.tasks = append(l.tasks[:index], l.tasks[index+1:]...)
}
func main() {
	Read := bufio.NewReader(os.Stdin)
	List := ListToDo{}
	for {
		var option int
		fmt.Println("Choose one option: \n",
			"1. Add a new Task\n",
			"2. Mark a task DONE\n",
			"3. Update a task\n",
			"4. Delete a task\n",
			"5. Exit")
		fmt.Printf("Put one option:")
		fmt.Scan(&option)

		switch option {
		case 1:
			var t Task
			fmt.Printf("Add a task's name: \n")
			fmt.Printf("Add a description: \n")
			t.name, _ = Read.ReadString('\n')
			t.description, _ = Read.ReadString('\n')
			List.addTask(t)
			fmt.Println("The task was added")
		case 2:
			var index int
			fmt.Print("Add index to mark a task like complete:")
			fmt.Scan(&index)
			List.MarkDone(index)
			fmt.Println("This task was marked as complete.")
		case 3:
			var index int
			var t Task
			fmt.Print("Add index to update the task : \n")
			fmt.Scan(&index)
			fmt.Print("Add a task's name: \n")
			t.name, _ = Read.ReadString('\n')
			fmt.Print("Add a description: ")
			t.description, _ = Read.ReadString('\n')
			List.UpdateTask(index, t)
			fmt.Println("This task was updated.")
		case 4:
			var index int
			fmt.Print("Add index to delete the task :")
			fmt.Scan(&index)
			List.deleteTask(index)
			fmt.Println("This task was deleted.")
		case 5:
			fmt.Println("Exit...")
			return
		default:

			fmt.Println("Invalid option.")
		}

		fmt.Println("List to DO:")
		fmt.Println("============================================")
		for i, t := range List.tasks {
			fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.name, t.description, t.complete)
		}
		fmt.Println("============================================")

	}

}
