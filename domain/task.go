package domain

import "fmt"

type Task interface {
	ShowDetails() string
	isDone() bool
}

type standartTask struct {
	title       string
	description string
	done        bool
}

func (t standartTask) ShowDetails() string {
	return fmt.Sprintf("Title: %s | Description: %s | Done: %s", t.title, t.description, t.done)
}

func (t standartTask) isDone() bool { return t.done }
