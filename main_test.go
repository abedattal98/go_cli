package main

import (
	"exampleTodo/models"
	"testing"
)

func Test_waitForTask(t *testing.T) {
	type args struct {
		task    *models.Task
		minutes int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "waitForTask",
			args: args{
				task: &models.Task{
					Text: "test",
				},
				minutes: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := waitForTask(tt.args.task, tt.args.minutes); (err != nil) != tt.wantErr {
				t.Errorf("waitForTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
