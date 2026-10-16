package entity

import (
	"slices"
	"time"
)

type TaskStatus string

const (
	TaskStatusTodo       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusDone       TaskStatus = "done"
)

type Task struct {
	ID          string
	UserID      string
	Title       string
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (s TaskStatus) Valid() bool {
	switch s {
	case TaskStatusTodo, TaskStatusInProgress, TaskStatusDone:
		return true
	default:
		return false
	}
}

func (t *Task) Transition(newStatus TaskStatus) error {
	validTransitions := map[TaskStatus][]TaskStatus{
		TaskStatusTodo:       {TaskStatusInProgress},
		TaskStatusInProgress: {TaskStatusDone, TaskStatusTodo},
		TaskStatusDone:       {},
	}

	allowed, ok := validTransitions[t.Status]

	if !ok {
		return ErrInvalidTransition
	}

	if slices.Contains(allowed, newStatus) {
		t.Status = newStatus

		return nil
	}

	return ErrInvalidTransition
}
