package domain

import "fmt"

type StandartTask struct {
	title       string
	description string
	done        bool
}

func NewStandartTask(title, description string, done bool) StandartTask {
	return StandartTask{
		title:       title,
		description: description,
		done:        done,
	}
}

func (t StandartTask) ShowDetails() string {
	return fmt.Sprintf("Title: %s | Description: %s | Done: %t", t.title, t.description, t.done)
}

func (t StandartTask) isDone() bool { return t.done }
