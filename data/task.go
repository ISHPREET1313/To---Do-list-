package data

import "fmt"

type Task struct {
	Id   int
	Name string
}

func (task Task) String() string {
	return fmt.Sprintf("%d ---  %v", task.Id, task.Name)
}

func NewTask(id int, name string) Task {
	return Task{Id: id, Name: name}
}
