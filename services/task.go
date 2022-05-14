package services

import (
	"context"
	"exampleTodo/interfaces"
	"exampleTodo/models"
	"fmt"
	"log"
	"os"
	"time"
)

type TaskService struct {
	taskRepo interfaces.TaskRepository
}

func NewTaskService(taskRepo interfaces.TaskRepository) interfaces.TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) GetTasks(ctx context.Context) ([]*models.Task, error) {
	return s.taskRepo.GetTasks(ctx)
}
func (s *TaskService) CreateTask(ctx context.Context, taskName string) (*models.Task, error) {
	return s.taskRepo.CreateTask(ctx, taskName)
}
func (s *TaskService) CompleteTask(ctx context.Context, id string) (*models.Task, error) {
	task, err := s.taskRepo.GetTask(ctx, os.Args[2])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if task.Completed {
		fmt.Printf("Task %s is already complete\n", task.Text)
		os.Exit(1)
	}
	return s.taskRepo.CompleteTask(ctx, id)
}

func (s *TaskService) WaitForTask(ctx context.Context, id string, minutes int) (*models.Task, error) {
	// check if the task exists
	task, err := s.taskRepo.GetTask(ctx, os.Args[2])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// check if the task is complete
	if task.Completed {
		fmt.Printf("Task %s is already complete\n", task.Text)
		os.Exit(1)
	}
	waitForTask(task, minutes)

	return s.taskRepo.CompleteTask(ctx, id)
}

func waitForTask(task *models.Task, minutes int) error {
	fmt.Printf("Waiting for task %s to complete...\n", task.Text)
	// creates a new Timer that will send the current time on its channel after at least duration d.
	timer1 := time.NewTimer(time.Duration(minutes) * time.Minute)
	// checks if the timer has expired
	<-timer1.C
	fmt.Printf("Task %s is complete!\n", task.Text)
	return nil
}
