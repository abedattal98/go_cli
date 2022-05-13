package repositories

import (
	"context"
	"exampleTodo/models"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)


func TestTaskRepository_CreateTask(t *testing.T) {
	type fields struct {
		db *mongo.Collection
	}
	type args struct {
		ctx      context.Context
		taskName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Task
		wantErr bool
	}{
		{
			name: "CreateTask",
			fields: fields{
				db: nil,
			},
			args: args{
				ctx:      context.TODO(),
				taskName: "test",
			},
			want: &models.Task{
				Text: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &TaskRepository{
				db: tt.fields.db,
			}
			got, err := client.CreateTask(tt.args.ctx, tt.args.taskName)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskRepository.CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskRepository.CreateTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
