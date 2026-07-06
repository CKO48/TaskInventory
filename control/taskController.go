package control

import (
	"errors"
	"strings"
	"taskmanager/domain"
)

type TaskController struct {
	taskMap map[string]domain.Task
}

func NewTaskController(taskMap map[string]domain.Task) *TaskController {
	if taskMap == nil {
		taskMap = make(map[string]domain.Task)
	}
	return &TaskController{taskMap: taskMap}
}

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
