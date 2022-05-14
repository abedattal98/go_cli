package main

import (
	"context"
	"exampleTodo/models"
	"exampleTodo/repositories"
	"exampleTodo/services"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()
var db *mongo.Database

func init() {
	// set up connection to mongodb options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	// connect to mongodb
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// check connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// set db to use "tasker" database
	db = client.Database("tasker")
}

func main() {
	// create a new task repository
	repo := repositories.NewTaskRepo(db)
	// create a new task service
	taskService := services.NewTaskService(repo)
	// check command line arguments
	for _, arg := range os.Args[1:] {
		switch arg {
		case "add":
			flag.Parse()
			fmt.Printf("Hello from add %s\n", os.Args[2])
			//args validation
			if os.Args[2] == "" {
				fmt.Println("Please enter a task name")
				os.Exit(1)
			}

			task, err := taskService.CreateTask(ctx, os.Args[2])

			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			fmt.Printf("Task %s is added\n", task.Text)
			os.Exit(1)

		case "list":
			flag.Parse()
			fmt.Printf("Hello from list")

			tasks, err := taskService.GetTasks(ctx)

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

			task, err := taskService.CompleteTask(ctx, os.Args[2])

			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			fmt.Printf("Task %s is complete\n", task.Text)
			os.Exit(1)

		case "wait":
			fmt.Printf("Hello from wait %s\n", os.Args[2])
			if len(os.Args) < 3 {
				fmt.Printf("Please provide a task number to mark as complete\n")
				os.Exit(1)
			}
			// convert task minutes to int
			minutes, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			// check string valid
			if os.Args[3] == "" {
				fmt.Printf("Please provide a task number to mark as complete\n")
				os.Exit(1)
			}

			task, err := taskService.WaitForTask(ctx, os.Args[3], minutes)

			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			fmt.Printf("Task %s is complete\n", task.Text)
			os.Exit(1)

		default:
			fmt.Printf("No valid command %s\n", arg)
		}
	}
}

func printTasks(tasks []*models.Task) {
	for i, v := range tasks {
		if v.Completed {
			fmt.Printf("Completed task %d: %s %s\n", i+1, v.ID.String(), v.Text)
		} else {
			fmt.Printf("Pending task %d: %s %s \n", i+1, v.ID.String(), v.Text)
		}
	}
}
