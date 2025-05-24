package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/goyalaadi12071998/app/boot"
	"github.com/goyalaadi12071998/app/constants"
	"github.com/goyalaadi12071998/app/controllers"
	"github.com/goyalaadi12071998/app/models"
)

func main() {
	fmt.Println("**************************************")
	fmt.Println("basic validations are not added for now \nInterfaces are created and added comments as why not used")
	fmt.Println("**************************************")
	fmt.Println("")
	fmt.Println("")
	boot.Init()
	user1 := controllers.UserController.CreateUser("Aadesh Goyal", "test@gmail.com", "Test@1234")
	user2 := controllers.UserController.CreateUser("Rajat Garg", "test2@gmail.com", "Test@1234")

	for user := range models.Users {
		fmt.Println(user.Name + ":" + user.Email)
	}

	task1 := controllers.TaksController.CreateTask(
		"Task - A", constants.TASK_TYPE_TASK, user1, time.Now().Add(time.Hour*36).Unix(), 10, nil,
	)
	task2 := controllers.TaksController.CreateTask(
		"Task - B", constants.TASK_TYPE_TASK, user2, time.Now().Add(time.Hour*36).Unix(), 10, nil,
	)
	controllers.TaksController.CreateTask(
		"Task - C", constants.TASK_TYPE_SUBTASK, user1, time.Now().Add(time.Hour*36).Unix(), 10, task1,
	)
	controllers.TaksController.CreateTask(
		"Task - D", constants.TASK_TYPE_SUBTASK, user1, time.Now().Add(time.Hour*36).Unix(), 10, task1,
	)

	// controllers.TaksController.UpdateTask(task1, "Task-AA", 0, 3, "PENDING")
	// // tasks := controllers.TaksController.GetAllTasksForUser(user1, "")
	// fmt.Println(controllers.TaksController.GetTaskProgressPercentage(task1))
	controllers.TaksController.UpdateTask(task2, "Task-AA", 0, 3, "COMPLETED")
	fmt.Println(controllers.TaksController.GetTaskProgressPercentage(task2))

	tasks := controllers.TaksController.GetAllTasksForUser(user1, "")
	for _, task := range tasks {
		fmt.Println("Name:", task.Name+", Priority:"+strconv.Itoa(int(task.Priority)))
	}
}
