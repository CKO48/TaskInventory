package domain

type Task interface {
	ShowDetails() string
	IsDone() bool
	Complete()
}
