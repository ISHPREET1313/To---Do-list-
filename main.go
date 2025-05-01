package main

import (
	"fmt"

	"toDoList.ishpreet.com/data"
)

func printToDOList(list []data.Task) {
	fmt.Println("Your Today's To-Do List")
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
}

func delTask(list []data.Task, id int) []data.Task {
	return append(list[:id-1], list[id:]...)
}

func main() {
	list := []data.Task{}

	printToDOList(list)

	exitloop := false

	for !exitloop {
		fmt.Println("Select below options")
		fmt.Println("0 to view TASK")
		fmt.Println("1 to add TASK ")
		fmt.Println("2 to delete TASK")
		fmt.Println("3 for exit ")

		var input int
		fmt.Scanln(&input)
		switch input {
		case 0:
			printToDOList(list)

		case 1:
			var name string
			fmt.Println("Enter Your Task Name:")
			fmt.Scanln(&name)
			list = append(list, data.NewTask(len(list)+1, name))

		case 2:
			fmt.Println("Enter Id to delete a task")
			var id int
			fmt.Scanln(&id)
			if id <= len(list) {
				list = delTask(list, id)
			} else {
				fmt.Println("Invalid Id")
			}

		case 3:
			fmt.Println("Bye!!")
			exitloop = true

		default:
			fmt.Println("Invalid innput")
		}
	}
}
