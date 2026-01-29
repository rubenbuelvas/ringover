package application

import "github.com/rubenbuelvas/ringover/src/domain"

type ITaskService interface {
	GetTasks() ([]domain.Task, error)
	GetSubtasks(taskId int64) ([]domain.Task, error)
	CreateTask(data CreateTaskDTO) error
	PatchTask(taskId int64, data UpdateTaskDTO) error
	DeleteTask(taskId int64) error
}
