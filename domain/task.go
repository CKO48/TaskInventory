package domain

import "fmt"

type Task interface {
	ShowDetails() string
	isDone() bool
}

type StandartTask struct {
	title       string
	description string
	done        bool
}

func (t StandartTask) ShowDetails() string {
	return fmt.Sprintf("Title: %s | Description: %s | Done: %s", t.title, t.description, t.done)
}

func (t StandartTask) isDone() bool { return t.done }
