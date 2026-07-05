package control

import "taskmanager/domain"

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
