package application

import (
	"github.com/rubenbuelvas/ringover/src/domain"
	"go.uber.org/zap"
)

type TaskService struct {
	taskRepository domain.TaskRepository
}

func NewTaskService(taskRepository domain.TaskRepository) ITaskService {
	return &TaskService{
		taskRepository: taskRepository,
	}
}

func (service *TaskService) GetTasks() ([]domain.Task, error) {
	tasks, err := service.taskRepository.GetTasks()
	if err != nil {
		zap.L().Error("TaskService: Failed to get tasks", zap.Error(err))
	}
	return tasks, err
}

func (service *TaskService) GetTaskById() ([]domain.Task, error) {
	tasks, err := service.taskRepository.GetTasks()
	if err != nil {
		zap.L().Error("TaskService: Failed to get tasks", zap.Error(err))
	}
	return tasks, err
}

func (service *TaskService) GetSubtasks(taskId int64) ([]domain.Task, error) {
	subtasks, err := service.taskRepository.GetSubtasks(taskId)
	if err != nil {
		zap.L().Error("TaskService: Failed to get subtasks", zap.Error(err))
	}
	return subtasks, err
}

func (service *TaskService) CreateTask(taskDTO CreateTaskDTO) error {
	task := taskDTO.ToTask()
	err := service.taskRepository.CreateTask(*task)
	if err != nil {
		zap.L().Error("TaskService: Failed to create task", zap.Error(err))
	}
	return err
}

func (service *TaskService) PatchTask(taskId int64, data PatchTaskDTO) error {
	task, err := service.taskRepository.GetTaskById(taskId)
	if err != nil {
		zap.L().Error("TaskService: Failed to get task", zap.Error(err))
		return err
	}
	data.ApplyToTask(&task)
	err = service.taskRepository.PatchTask(taskId, task)
	if err != nil {
		zap.L().Error("TaskService: Failed to patch task", zap.Error(err))
	}
	return err
}

func (service *TaskService) DeleteTask(taskId int64) error {
	err := service.taskRepository.DeleteTask(taskId)
	if err != nil {
		zap.L().Error("TaskService: Failed to delete task", zap.Error(err))
	}
	return err
}
