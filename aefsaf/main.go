package main

import (
	"fmt"
	"strconv"
)


type Task struct {

	UserID int
	TaskID int
}

func main(){
	task := Task{TaskID:2}
	fmt.Println(task)
	fmt.Println(task.UserID)
	fmt.Println(strconv.Itoa(task.UserID))
}