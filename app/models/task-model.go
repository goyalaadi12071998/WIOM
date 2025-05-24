package models

import (
	"time"
)

var Tasks map[*Task]bool = map[*Task]bool{}

type Task struct {
	Name       string
	TaskType   string
	SourceTask *Task
	Status     string
	User       *User
	DueDate    int64
	Priority   int16
	CreatedAt  int64
	UpdatedAt  int64
}

type ITask interface {
	SaveTaskToDB()
	CheckIfExist() bool
	DeleteTask() bool
	UpdateTask() bool
	GetTaskProgressPercentage() int
}

func (t *Task) GetTaskProgressPercentage() int {
	allSubTasks := []*Task{}
	totalSubTask := 0
	for val := range Tasks {
		if val.SourceTask == t && val.Status == "COMPLETED" {
			totalSubTask += 1
			allSubTasks = append(allSubTasks, val)
		}
	}

	if totalSubTask == 0 {
		if t.Status != "COMPLETED" {
			return 0
		}
		return 100
	} else {
		return (len(allSubTasks) / totalSubTask) * 100
	}

}

func (t *Task) UpdateTask(name string, dueDate int64, priority int, status string, user *User) bool {
	if t.User != user {
		return false
	}

	if name != "" {
		t.Name = name
	}

	if dueDate != 0 {
		t.DueDate = dueDate
	}

	if priority != 0 {
		t.Priority = int16(priority)
	}

	if status != "" {
		t.Status = status
	}

	return true
}

func (t *Task) CheckIfExist() bool {
	for task := range Tasks {
		if task == t {
			return true
		}
	}

	return false
}

func (t *Task) DeleteTask(user *User) bool {
	if t.User != user {
		return false
	}

	_, exists := Tasks[t]
	if !exists {
		return false
	}

	delete(Tasks, t)
	return true
}

func (t *Task) SaveTaskToDB() {
	Tasks[t] = true
}

// I have created interfaces but not using here require little more code to but the functionality will going to be same
// we just have to remove *Task -> ITask
func NewTask(name string, taskType string, user *User, dueDate int64, priority int16, sourceTask *Task) *Task {
	return &Task{
		Name:       name,
		TaskType:   taskType,
		User:       user,
		SourceTask: sourceTask,
		DueDate:    dueDate,
		Priority:   priority,
		Status:     "CREATED",
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}
}
