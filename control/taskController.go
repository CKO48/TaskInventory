package control

import (
	"errors"
	"strings"
	"taskmanager/domain"
)

/*
TaskController is responsible for managing tasks in the task map.
*/
type TaskController struct {
	taskMap map[string]domain.Task
}

/*
Create a new TaskController with the provided task map.

If the task map is nil, a new empty map will be created
*/
func NewTaskController(taskMap map[string]domain.Task) *TaskController {
	if taskMap == nil {
		taskMap = make(map[string]domain.Task)
	}
	return &TaskController{taskMap: taskMap}
}

/*
Adds a new task to the task map.

The args slice should contain the command, title, and optionally a description.

Returns a success message or an error if the title is missing.
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

	_, exists := tc.taskMap[args[1]]

	if !exists {
		return "", errors.New("Task not found")
	}
	delete(tc.taskMap, args[1])

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

// Completes a task in the task map based on the provided title in the args slice.
//
// Returns a success message or an error if the task is not found or if no title is specified.
func (tc *TaskController) CompleteTask(args []string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("No task specified")
	}

	task, exists := tc.taskMap[args[1]]

	if exists {
		task.Complete()
		return "Task completed successfully", nil
	}

	return "", errors.New("Task not found")
}
