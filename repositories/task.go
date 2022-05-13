package repositories

import (
	"context"
	"exampleTodo/interfaces"
	"exampleTodo/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	db *mongo.Collection
}

func NewTaskRepo(db *mongo.Database) interfaces.TaskRepository {
	return &TaskRepository{db: db.Collection("tasks")}
}

func (client *TaskRepository) GetTasks(ctx context.Context) ([]*models.Task, error) {
	// A slice of tasks for storing the decoded documents
	var tasks []*models.Task

	cursor, err := client.db.Find(ctx, bson.D{})
	if err != nil {
		defer cursor.Close(ctx)
		return nil, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cursor.Next(ctx) {
		var t models.Task
		err := cursor.Decode(&t)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, &t)
	}

	if err := cursor.Err(); err != nil {
		return tasks, err
	}

	// once exhausted, close the cursor
	cursor.Close(ctx)

	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}

	return tasks, nil
}

func (client *TaskRepository) GetTask(ctx context.Context, id string) (*models.Task, error) {
	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": mongoID}

	t := &models.Task{}
	err = client.db.FindOne(ctx, filter).Decode(t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (client *TaskRepository) CreateTask(ctx context.Context, taskName string) (*models.Task, error) {
	task := &models.Task{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Text:      taskName,
		Completed: false,
	}
	_, err := client.db.InsertOne(ctx, task)
	return task, err
}

func (client *TaskRepository) CompleteTask(ctx context.Context, id string) error {
	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": mongoID}

	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "completed", Value: true},
	}}}

	t := &models.Task{}
	return client.db.FindOneAndUpdate(ctx, filter, update).Decode(t)
}
