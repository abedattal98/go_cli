package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getTasks() ([]*Task, error) {
	// A slice of tasks for storing the decoded documents
	var tasks []*Task

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		defer cursor.Close(ctx)
		return nil, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cursor.Next(ctx) {
		var t Task
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
func getTask(id string) (*Task, error) {
	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": mongoID}

	t := &Task{}
	err = collection.FindOne(ctx, filter).Decode(t)
	if err != nil {
		return nil, err
	}

	return t, nil
}
func createTask(taskName string) (*Task, error) {
	task := &Task{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Text:      taskName,
		Completed: false,
	}
	_, err := collection.InsertOne(ctx, task)
	return task, err
}
func completeTask(id string) error {
	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": mongoID}

	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "completed", Value: true},
	}}}

	t := &Task{}
	return collection.FindOneAndUpdate(ctx, filter, update).Decode(t)
}
