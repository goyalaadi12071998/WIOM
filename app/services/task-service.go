package services

import (
	"errors"
	"log"

	"github.com/goyalaadi12071998/app/models"
)

var taskService TaskService

type TaskService struct {
}

func InitialiseTaskService() ITaskService {
	taskService = TaskService{}
	return taskService
}

type ITaskService interface {
	CreateTask(name string, taskType string, user *models.User, dueDate int64, priority int16, sourceTask *models.Task) *models.Task
	DeleteTask(task *models.Task, user *models.User) bool
	GetTasksForUser(user *models.User, status string) []models.Task
	UpdateTask(task *models.Task, name string, dueDate int64, priority int, status string, user *models.User) bool
	GetTaskProgressPercentage(task *models.Task) int
}

func (t TaskService) CreateTask(name string, taskType string, user *models.User, dueDate int64, priority int16, sourceTask *models.Task) *models.Task {
	task := models.NewTask(name, taskType, user, dueDate, int16(priority), sourceTask)
	task.SaveTaskToDB()
	return task
}

func (t TaskService) DeleteTask(task *models.Task, user *models.User) bool {
	taskExist := task.CheckIfExist()
	if taskExist {
		return task.DeleteTask(user)
	}

	return false
}

func (t TaskService) GetTasksForUser(user *models.User, status string) []models.Task {
	allTasks := models.Tasks
	result := []models.Task{}
	for task := range allTasks {
		if task.User == user {
			if status != "" && task.Status != status {
				continue
			}
			result = append(result, *task)
		}
	}
	return result
}

func (t TaskService) UpdateTask(task *models.Task, name string, dueDate int64, priority int, status string, user *models.User) bool {
	exist := task.CheckIfExist()
	if !exist {
		log.Fatal(errors.New("task does not exist"))
	}

	success := task.UpdateTask(name, dueDate, priority, status, user)
	return success
}

func (t TaskService) GetTaskProgressPercentage(task *models.Task) int {
	return task.GetTaskProgressPercentage()
}
