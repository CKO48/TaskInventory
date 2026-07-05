package control

import "taskmanager/domain"

type TaskController struct {
	taskmap map[string]domain.Task
}

func NewTaskController(taskmap map[string]domain.Task) TaskController {
	return TaskController{taskmap}
}

func (tc *TaskController) AddTask(args []string) (string, error) {
	// Implementation for adding a task
	return "", nil
}

func (tc *TaskController) RemoveTask(args []string) (string, error) {
	// Implementation for removing a task
	return "", nil
}

func (tc TaskController) ListTasks(args []string) (string, error) {
	// Implementation for listing all the tasks
	return "", nil
}

func (tc TaskController) CompleteTask(args []string) (string, error) {
	// Implementation for completing a task
	return "", nil
}
