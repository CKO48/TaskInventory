package domain

// Task is a data structure representing a task with basic operations.
type Task interface {
	ShowDetails() string
	IsDone() bool
	Complete()
}
