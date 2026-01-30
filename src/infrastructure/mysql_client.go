package infrastructure

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rubenbuelvas/ringover/src/domain"
)

type MySQLClient struct {
	db sqlx.DB
}

func NewMySQLClient(config string) (domain.TaskRepository, error) {
	db, err := sqlx.Connect("mysql", config)
	if err != nil {
		return nil, err
	}
	return &MySQLClient{
		db: *db,
	}, nil
}

func (client *MySQLClient) GetTasks() ([]domain.Task, error) {
	result := []domain.Task{}
	query := `
		SELECT t.id, t.title, t.description, t.status, t.priority, t.due_date, 
               t.completed_at, t.parent_task_id, t.category_id, t.created_at, t.updated_at,
               c.id, c.name
        FROM tasks t
        JOIN categories c 
        ON t.category_id = c.id
	`
	err := client.db.Select(&result, query)
	return result, err
}

func (client *MySQLClient) GetTaskById(taskId int64) (domain.Task, error) {
	result := domain.Task{}
	query := `
		SELECT *
        FROM tasks t
        JOIN categories c 
        ON t.category_id = c.id
		WHERE t.id = ?
	`
	err := client.db.Get(&result, query, taskId)
	return result, err
}

func (client *MySQLClient) GetSubtasks(taskId int64) ([]domain.Task, error) {
	result := []domain.Task{}
	query := `
		SELECT * FROM tasks
		WHERE parent_task_id = ?
	`
	err := client.db.Select(&result, query, taskId)
	return result, err
}

func (client *MySQLClient) CreateTask(task domain.Task) error {
	query := `
		INSERT INTO tasks (title, description, status, priority, due_date, completed_at, parent_task_id, category_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := client.db.Exec(query, task.Title, task.Description, task.Status, task.Priority, task.DueDate, task.CompletedAt, task.ParentTaskId, task.CategoryId)
	return err
}

func (client *MySQLClient) PatchTask(taskId int64, data domain.Task) error {
	query := `
		UPDATE tasks
		SET title = ?, description = ?, status = ?, priority = ?, due_date = ?, completed_at = ?, parent_task_id = ?, category_id = ?
		WHERE id = ?
	`
	_, err := client.db.Exec(query, data.Title, data.Description, data.Status, data.Priority, data.DueDate, data.CompletedAt, data.ParentTaskId, data.CategoryId, taskId)
	return err
}

func (client *MySQLClient) DeleteTask(taskId int64) error {
	query := `
		DELETE FROM tasks
		WHERE id = ?
	`
	_, err := client.db.Exec(query, taskId)
	return err
}
