package controllers

import (
	"github.com/goyalaadi12071998/app/models"
	"github.com/goyalaadi12071998/app/services"
)

var TaksController taskcontroller

type taskcontroller struct {
	taskService services.ITaskService
}

func InitaliseTaskController(taskService services.ITaskService) {
	TaksController = taskcontroller{
		taskService: taskService,
	}
}

func (t taskcontroller) CreateTask(name string, taskType string, user *models.User, dueDate int64, priority int16, sourceTask *models.Task) *models.Task {
	task := t.taskService.CreateTask(name, taskType, user, dueDate, priority, sourceTask)
	return task
}

func (t taskcontroller) DeleteTask(task *models.Task, user *models.User) bool {
	success := t.taskService.DeleteTask(task, user)
	return success
}

func (t taskcontroller) GetAllTasksForUser(user *models.User, status string) []models.Task {
	result := t.taskService.GetTasksForUser(user, status)
	return result
}

func (t taskcontroller) GetTaskProgressPercentage(task *models.Task) int {
	return t.taskService.GetTaskProgressPercentage(task)
}

func (t taskcontroller) UpdateTask(task *models.Task, name string, dueDate int64, priority int, status string, user *models.User) bool {
	return t.taskService.UpdateTask(task, name, dueDate, priority, status, user)
}
