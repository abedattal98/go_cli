package interfaces

import (
	"context"
	"exampleTodo/models"
)

type TaskService interface {
	GetTasks(ctx context.Context) ([]*models.Task, error)
	CreateTask(ctx context.Context, taskName string) (*models.Task, error)
	CompleteTask(ctx context.Context, task models.Task) (*models.Task, error)
	WaitForTask(ctx context.Context, task models.Task, minutes int) (*models.Task, error)
	GetTask(ctx context.Context, id string) (*models.Task, error)
}
