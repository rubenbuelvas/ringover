package domain

type Category struct {
	ID   *int64  `db:"id"`
	Name *string `db:"name"`
}

type Task struct {
	ID           *int64  `db:"id"`
	Title        *string `db:"title"`
	Description  *string `db:"description"`
	Status       *string `db:"status"`
	Priority     *int8   `db:"priority"`
	DueDate      *string `db:"due_date"`
	CompletedAt  *string `db:"completed_at"`
	ParentTaskId *int64  `db:"parent_task_id"`
	CategoryID   *int64  `db:"category_id"`
	CreatedAt    *string `db:"created_at"`
	UpdatedAt    *string `db:"updated_at"`
	Category     *Category
}
