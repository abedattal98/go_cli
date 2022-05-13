package interfaces

import (
	"context"
	"exampleTodo/models"
)

type TaskRepository interface {
	GetTasks(ctx context.Context) ([]*models.Task, error)
	CreateTask(ctx context.Context, taskName string) (*models.Task, error)
	CompleteTask(ctx context.Context, id string) error
	GetTask(ctx context.Context, id string) (*models.Task, error)
}
