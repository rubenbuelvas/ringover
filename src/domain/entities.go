package domain

type Task struct {
	Id           *int64  `db:"id" json:"id"`
	Title        *string `db:"title" json:"title"`
	Description  *string `db:"description" json:"description,omitempty"`
	Status       *string `db:"status" json:"status,omitempty"`
	Priority     *int8   `db:"priority" json:"priority,omitempty"`
	DueDate      *string `db:"due_date" json:"due_date,omitempty"`
	CompletedAt  *string `db:"completed_at" json:"completed_at,omitempty"`
	ParentTaskId *int64  `db:"parent_task_id" json:"parent_task_id,omitempty"`
	CategoryId   *int64  `db:"category_id" json:"category_id,omitempty"`
	CategoryName *string `db:"category_name" json:"category_name,omitempty"`
	CreatedAt    *string `db:"created_at" json:"created_at"`
	UpdatedAt    *string `db:"updated_at" json:"updated_at"`
}
