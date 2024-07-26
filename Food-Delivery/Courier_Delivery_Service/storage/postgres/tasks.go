package postgres

import (
	"context"
	"courier_delivery/genproto"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Task struct {
	Db *pgx.Conn
}

func NewTask(db *pgx.Conn) *Task {
	return &Task{Db: db}
}

// CreateTask yangi vazifani yaratadi
func (t *Task) CreateTask(ctx context.Context, req *genproto.CreateTaskRequest) (*genproto.TaskResponse, error) {
	query := `INSERT INTO tasks (title, description, status, assigned_to, due_date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) RETURNING id, created_at, updated_at`
	var task genproto.TaskResponse
	err := t.Db.QueryRow(ctx, query, req.Title, req.Description, req.Status, req.AssignedTo, req.DueDate).Scan(&task.Id, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}
	task.Title = req.Title
	task.Description = req.Description
	task.Status = req.Status
	task.AssignedTo = req.AssignedTo
	task.DueDate = req.DueDate
	return &task, nil
}

// GetTask vazifani ID orqali qaytaradi
func (t *Task) GetTask(ctx context.Context, req *genproto.GetTaskRequest) (*genproto.TaskResponse, error) {
	query := `SELECT id, title, description, status, assigned_to, due_date, created_at, updated_at FROM tasks WHERE id=$1`
	var task genproto.TaskResponse
	err := t.Db.QueryRow(ctx, query, req.Id).Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.AssignedTo, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}
	return &task, nil
}

// UpdateTask vazifani yangilaydi
func (t *Task) UpdateTask(ctx context.Context, req *genproto.UpdateTaskRequest) (*genproto.TaskResponse, error) {
	query := `UPDATE tasks SET title=$1, description=$2, status=$3, assigned_to=$4, due_date=$5, updated_at=NOW() WHERE id=$6 RETURNING id, title, description, status, assigned_to, due_date, created_at, updated_at`
	var task genproto.TaskResponse
	err := t.Db.QueryRow(ctx, query, req.Title, req.Description, req.Status, req.AssignedTo, req.DueDate, req.Id).Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.AssignedTo, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}
	return &task, nil
}

// DeleteTask vazifani o'chiradi
func (t *Task) DeleteTask(ctx context.Context, req *genproto.DeleteTaskRequest) (*genproto.DeleteTaskResponse, error) {
	query := `DELETE FROM tasks WHERE id=$1`
	_, err := t.Db.Exec(ctx, query, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete task: %w", err)
	}
	return &genproto.DeleteTaskResponse{
		Success: true,
		Message: "Task deleted successfully",
	}, nil
}

// GetAllTasks berilgan foydalanuvchiga tayinlangan barcha vazifalarni qaytaradi
func (t *Task) GetAllTasks(ctx context.Context, req *genproto.GetAllTasksRequest) (*genproto.GetAllTasksResponse, error) {
	query := `SELECT id, title, description, status, assigned_to, due_date, created_at, updated_at FROM tasks WHERE assigned_to=$1`
	rows, err := t.Db.Query(ctx, query, req.AssignedTo)
	if err != nil {
		return nil, fmt.Errorf("failed to get all tasks: %w", err)
	}
	defer rows.Close()

	var tasks []*genproto.TaskResponse
	for rows.Next() {
		var task genproto.TaskResponse
		err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.AssignedTo, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, &task)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to iterate over tasks: %w", rows.Err())
	}

	return &genproto.GetAllTasksResponse{
		Tasks: tasks,
	}, nil
}
