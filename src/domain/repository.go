package domain

type TaskRepository interface {
	GetTasks() ([]Task, error)
	GetSubtasks(taskId int64) ([]Task, error)
	CreateTask(task Task) error
	PatchTask(taskId int64, data Task) error
	DeleteTask(taskId int64) error
}
