package services

import (
	"context"
	"exampleTodo/interfaces"
	"exampleTodo/models"
	"exampleTodo/services/mocks"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTaskService_GetTasks(t *testing.T) {
	id1 := primitive.NewObjectID()
	id2 := primitive.NewObjectID()

	ctrl := gomock.NewController(t)
	mockTaskRepo := mocks.NewMockTaskRepository(ctrl)

	type fields struct {
		taskRepo interfaces.TaskRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		setup   func()
		want    []*models.Task
		wantErr bool
	}{
		{
			name: "GetTasks",
			fields: fields{
				taskRepo: mockTaskRepo,
			},
			args: args{
				ctx: context.Background(),
			},
			setup: func() {
				tasks := []*models.Task{
					{
						ID:        id1,
						Text:      "Task 1",
						Completed: false,
					},
					{
						ID:        id2,
						Text:      "Task 2",
						Completed: false,
					},
				}
				mockTaskRepo.EXPECT().GetTasks(gomock.Any()).Return(tasks, nil)
			},
			want: []*models.Task{
				{
					ID:        id1,
					Text:      "Task 1",
					Completed: false,
				},
				{
					ID:        id2,
					Text:      "Task 2",
					Completed: false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			s := &TaskService{
				taskRepo: tt.fields.taskRepo,
			}
			got, err := s.GetTasks(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskService.GetTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskService.GetTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskService_CreateTask(t *testing.T) {
	id1 := primitive.NewObjectID()

	ctrl := gomock.NewController(t)
	mockTaskRepo := mocks.NewMockTaskRepository(ctrl)
	type fields struct {
		taskRepo interfaces.TaskRepository
	}
	type args struct {
		ctx      context.Context
		taskName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		setup   func()
		want    *models.Task
		wantErr bool
	}{
		{
			name: "CreateTask",
			fields: fields{
				taskRepo: mockTaskRepo,
			},
			args: args{
				ctx:      context.Background(),
				taskName: "Task 1",
			},
			setup: func() {
				task := &models.Task{
					ID:        id1,
					Text:      "Task 1",
					Completed: false,
				}
				mockTaskRepo.EXPECT().CreateTask(gomock.Any(), "Task 1").Return(task, nil)
			},
			want: &models.Task{
				ID:        id1,
				Text:      "Task 1",
				Completed: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			s := &TaskService{
				taskRepo: tt.fields.taskRepo,
			}
			got, err := s.CreateTask(tt.args.ctx, tt.args.taskName)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskService.CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskService.CreateTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskService_CompleteTask(t *testing.T) {
	id1 := primitive.NewObjectID()

	ctrl := gomock.NewController(t)
	mockTaskRepo := mocks.NewMockTaskRepository(ctrl)

	type fields struct {
		taskRepo interfaces.TaskRepository
	}
	type args struct {
		ctx  context.Context
		task models.Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		setup   func()
		want    *models.Task
		wantErr bool
	}{
		{
			name: "CompleteTask",
			fields: fields{
				taskRepo: mockTaskRepo,
			},
			args: args{
				ctx: context.Background(),
				task: models.Task{
					ID:        id1,
					Text:      "Task 1",
					Completed: false,
				},
			},
			setup: func() {
				task := &models.Task{
					ID:        id1,
					Text:      "Task 1",
					Completed: true,
				}
				mockTaskRepo.EXPECT().CompleteTask(gomock.Any(), gomock.Any()).Return(task, nil)
			},
			want: &models.Task{
				ID:        id1,
				Text:      "Task 1",
				Completed: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			s := &TaskService{
				taskRepo: tt.fields.taskRepo,
			}
			got, err := s.CompleteTask(tt.args.ctx, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskService.CompleteTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskService.CompleteTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskService_GetTask(t *testing.T) {
	id1 := primitive.NewObjectID()

	ctrl := gomock.NewController(t)
	mockTaskRepo := mocks.NewMockTaskRepository(ctrl)
	type fields struct {
		taskRepo interfaces.TaskRepository
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		setup   func()
		want    *models.Task
		wantErr bool
	}{
		{
			name: "GetTask",
			fields: fields{
				taskRepo: mockTaskRepo,
			},
			args: args{
				ctx: context.Background(),
				id:  id1.Hex(),
			},
			setup: func() {
				task := &models.Task{
					ID:        id1,
					Text:      "Task 1",
					Completed: false,
				}
				mockTaskRepo.EXPECT().GetTask(gomock.Any(), id1.Hex()).Return(task, nil)
			},
			want: &models.Task{
				ID:        id1,
				Text:      "Task 1",
				Completed: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			s := &TaskService{
				taskRepo: tt.fields.taskRepo,
			}
			got, err := s.GetTask(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskService.GetTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskService.GetTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
