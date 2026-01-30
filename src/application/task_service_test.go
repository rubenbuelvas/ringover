package application

import (
	"errors"
	"testing"

	"github.com/rubenbuelvas/ringover/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetTasks_Success(t *testing.T) {
	mockRepo := domain.NewMockTaskRepository(t)
	id := int64(1)
	title := "Test Task"
	tasks := []domain.Task{
		{
			Id:    &id,
			Title: &title,
		},
	}
	mockRepo.EXPECT().GetTasks().Return(tasks, nil)
	service := NewTaskService(mockRepo)
	result, err := service.GetTasks()

	assert.NoError(t, err)
	assert.Equal(t, tasks, result)
}

func TestGetTasks_Error(t *testing.T) {
	mockRepo := domain.NewMockTaskRepository(t)
	mockRepo.EXPECT().GetTasks().Return(nil, errors.New("database error"))
	service := NewTaskService(mockRepo)
	result, err := service.GetTasks()

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "database error", err.Error())
}
