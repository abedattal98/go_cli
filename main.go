package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("tasker").Collection("tasks")
}

func main() {
	for _, arg := range os.Args[1:] {
		switch arg {
			
		case "add":
			flag.Parse()
			fmt.Printf("Hello from add %s\n", os.Args[2])
			task, err := createTask(os.Args[2])
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			fmt.Printf("Task %s is added\n", task.Text)
			os.Exit(1)

		case "list":
			flag.Parse()
			fmt.Printf("Hello from list")
			tasks, err := getTasks()
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			printTasks(tasks)
			os.Exit(1)

		case "do":
			if len(os.Args) < 2 {
				fmt.Printf("Please provide a task number to mark as complete\n")
				return
			}
			fmt.Printf("Hello we are doing %s\n", os.Args[2])

			task, err := getTask(os.Args[2])
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			if task.Completed {
				fmt.Printf("Task %s is already complete\n", task.Text)
				os.Exit(1)
			}
			completeTask(os.Args[2])
			fmt.Printf("Task %s is complete\n", task.Text)
			os.Exit(1)

		case "wait":
			fmt.Printf("Hello from wait %s\n", os.Args[2])
			if len(os.Args) < 3 {
				fmt.Printf("Please provide a task number to mark as complete\n")
				os.Exit(1)
			}
			minutes, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			task, err := getTask(os.Args[3])
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			if task.Completed {
				fmt.Printf("Task %s is already complete\n", task.Text)
				os.Exit(1)
			}
			waitForTask(task, minutes)
			completeTask(os.Args[3])
			fmt.Printf("Task %s is complete\n", task.Text)
			os.Exit(1)

		default:
			fmt.Printf("No valid command %s\n", arg)
		}
	}
}

func printTasks(tasks []*Task) {
	for i, v := range tasks {
		if v.Completed {
			fmt.Printf("Completed task %d: %s %s\n", i+1, v.ID.String(), v.Text)
		} else {
			fmt.Printf("Pending task %d: %s %s \n", i+1, v.ID.String(), v.Text)
		}
	}
}

func waitForTask(task *Task, minutes int) error {
	fmt.Printf("Waiting for task %s to complete...\n", task.Text)

	// creates a new Timer that will send the current time on its channel after at least duration d.
	timer1 := time.NewTimer(time.Duration(minutes) * time.Minute)
	<-timer1.C

	fmt.Printf("Task %s is complete!\n", task.Text)
	return nil
}
