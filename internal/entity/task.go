package entity

import (
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

// Valid reports whether s is a known task status.
func (s TaskStatus) Valid() bool {
	switch s {
	case TaskStatusDone, TaskStatusInProgress, TaskStatusTodo:
		return true
	default:
		return false
	}
}

// Transition validates and applies a status transition
func (t *Task) Transition(newStatus TaskStatus) error {
	validTransation := map[TaskStatus][]TaskStatus{
		TaskStatusTodo:       {TaskStatusInProgress},
		TaskStatusInProgress: {TaskStatusDone, TaskStatusTodo},
		TaskStatusDone:       {},
	}

	allowed, ok := validTransation[t.Status]
	if !ok {
		return ErrInvalidTransition
	}

	for _, status := range allowed {
		if status == newStatus {
			t.Status = newStatus
			return nil
		}
	}

	return ErrInvalidTransition
}
