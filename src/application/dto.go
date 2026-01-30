package application

import "github.com/rubenbuelvas/ringover/src/domain"

type CreateTaskDTO struct {
	Title        string  `json:"title"`
	Description  *string `json:"description,omitempty"`
	Status       *string `json:"status,omitempty"`
	Priority     *int8   `json:"priority,omitempty"`
	DueDate      *string `json:"due_date,omitempty"`
	ParentTaskId *int64  `json:"parent_task_id,omitempty"`
	CategoryId   *int64  `json:"category_id,omitempty"`
}

func (dto *CreateTaskDTO) ToTask() *domain.Task {
	return &domain.Task{
		Title:        &dto.Title,
		Description:  dto.Description,
		Status:       dto.Status,
		Priority:     dto.Priority,
		DueDate:      dto.DueDate,
		ParentTaskId: dto.ParentTaskId,
		CategoryId:   dto.CategoryId,
	}
}

type PatchTaskDTO struct {
	Title        *string `json:"title,omitempty"`
	Description  *string `json:"description,omitempty"`
	Status       *string `json:"status,omitempty"`
	Priority     *int8   `json:"priority,omitempty"`
	DueDate      *string `json:"due_date,omitempty"`
	ParentTaskId *int64  `json:"parent_task_id,omitempty"`
	CategoryId   *int64  `json:"category_id,omitempty"`
}

func (dto *PatchTaskDTO) ApplyToTask(task *domain.Task) {
	if dto.Title != nil {
		task.Title = dto.Title
	}
	if dto.Description != nil {
		task.Description = dto.Description
	}
	if dto.Status != nil {
		task.Status = dto.Status
	}
	if dto.Priority != nil {
		task.Priority = dto.Priority
	}
	if dto.DueDate != nil {
		task.DueDate = dto.DueDate
	}
	if dto.ParentTaskId != nil {
		task.ParentTaskId = dto.ParentTaskId
	}
	if dto.CategoryId != nil {
		task.CategoryId = dto.CategoryId
	}
}
