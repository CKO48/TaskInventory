package domain

type Task interface {
	ShowDetails() string
	isDone() bool
}
