package application

type CreateTaskDTO struct {
	Title        string  `json:"title"`
	Description  *string `json:"description,omitempty"`
	Status       *string `json:"status,omitempty"`
	Priority     *int8   `json:"priority,omitempty"`
	DueDate      *string `json:"due_date,omitempty"`
	ParentTaskId *int64  `json:"parent_task_id,omitempty"`
	CategoryID   *int64  `json:"category_id,omitempty"`
}

type PatchTaskDTO struct {
	Title        *string `json:"title,omitempty"`
	Description  *string `json:"description,omitempty"`
	Status       *string `json:"status,omitempty"`
	Priority     *int8   `json:"priority,omitempty"`
	DueDate      *string `json:"due_date,omitempty"`
	ParentTaskId *int64  `json:"parent_task_id,omitempty"`
	CategoryID   *int64  `json:"category_id,omitempty"`
}
