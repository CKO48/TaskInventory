package control

import (
	"context"
	"errors"
	"strings"
	"taskmanager/domain"
	"taskmanager/repo"
	sqlite "taskmanager/repo/sqlite"
	"time"
)

/*
TaskController is responsible for managing tasks in the task map.
*/
type TaskController struct {
	taskMap map[string]domain.Task
	db      repo.Database
}

/*
Create a new instance of a TaskController
*/
func NewTaskController() *TaskController {
	db, err := sqlite.InitSQLiteDB()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	taskMap, err := db.Load(ctx)
	if err != nil {
		panic(err)
	}

	return &TaskController{taskMap: taskMap, db: db}
}

/*
Adds a new task to the task map based on the provided title and optional description in the args slice.
Returns a success message or an error if the title is not provided or if there is an issue saving the task to the database.
*/
func (tc *TaskController) AddTask(args []string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("All tasks must have a title")
	}

	var newTask domain.Task
	var description string
	var title string = args[1]

	if len(args) > 2 {
		description = args[2]
	} else {
		description = ""
	}

	newTask = domain.NewStandartTask(title, description, false)

	tc.taskMap[title] = newTask
	_, err := tc.db.SaveTask(title, description, false)
	if err != nil {
		return "", err
	}

	return "Task added successfully", nil
}

/*
Removes a task from the task map based on the provided title in the args slice.
Returns a success message or an error if the task is not found or if no title is specified.
*/
func (tc *TaskController) RemoveTask(args []string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("No task specified")
	}

	var title string = args[1]
	_, exists := tc.taskMap[title]

	if !exists {
		return "", errors.New("Task not found")
	}
	delete(tc.taskMap, title)
	_, err := tc.db.DeleteTask(title)
	if err != nil {
		return "", err
	}

	return "Task removed successfully", nil
}

/*
Lists all tasks in the task map.

Returns a formatted string with task details or an error if no tasks are found.
*/
func (tc *TaskController) ListTasks(args []string) (string, error) {
	sep := "-----------------------------------------------------------\n"
	var sb strings.Builder
	sb.WriteString(sep)

	for _, v := range tc.taskMap {
		sb.WriteString(v.ShowDetails())
		sb.WriteString("\n")
		sb.WriteString(sep)
	}
	return sb.String(), nil
}

/*
marks a task as completed based on the provided title in the args slice.
Returns a success message or an error if the task is not found or if no title is specified.
*/
func (tc *TaskController) CompleteTask(args []string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("No task specified")
	}

	task, exists := tc.taskMap[args[1]]

	if exists {
		task.Complete()
		_, err := tc.db.SaveTask(task.GetTitle(), task.GetDescription(), task.GetStatus())
		if err != nil {
			return "", err
		}
		return "Task completed successfully", nil
	}

	return "", errors.New("Task not found")
}

// Close closes the database connection associated with the TaskController.
func (tc *TaskController) Close() error {
	return tc.db.Close()
}
