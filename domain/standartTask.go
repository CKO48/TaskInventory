package domain

import "fmt"

/*
StandartTask represents a basic task with a title, description, and completion status.
*/
type StandartTask struct {
	title       string
	description string
	done        bool
}

// Creates a new StandartTask with the provided title, description, and completion status.
func NewStandartTask(title, description string, done bool) *StandartTask {
	return &StandartTask{
		title:       title,
		description: description,
		done:        done,
	}
}

// ShowDetails returns a formatted string with the task's details.
func (t *StandartTask) ShowDetails() string {
	return fmt.Sprintf("Title: %s | Description: %s | Done: %t", t.title, t.description, t.done)
}

func (t *StandartTask) IsDone() bool { return t.done }

func (t *StandartTask) Complete() { t.done = true }
