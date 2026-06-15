package entity

import (
	"errors"
	"testing"
)

func TestTask_Transition(t *testing.T) {
	tests := []struct {
		name      string
		from      TaskStatus
		to        TaskStatus
		wantErr   bool
		wantState TaskStatus
	}{
		{name: "todo to in_progress", from: TaskStatusTodo, to: TaskStatusInProgress, wantErr: false, wantState: TaskStatusInProgress},
		{name: "in_progress to done", from: TaskStatusInProgress, to: TaskStatusDone, wantErr: false, wantState: TaskStatusDone},
		{name: "in_progress to todo", from: TaskStatusInProgress, to: TaskStatusTodo, wantErr: false, wantState: TaskStatusTodo},
		{name: "todo to done (invalid)", from: TaskStatusTodo, to: TaskStatusDone, wantErr: true, wantState: TaskStatusTodo},
		{name: "done to in_progress (invalid)", from: TaskStatusDone, to: TaskStatusInProgress, wantErr: true, wantState: TaskStatusDone},
		{name: "done to todo (invalid)", from: TaskStatusDone, to: TaskStatusTodo, wantErr: true, wantState: TaskStatusDone},
		{name: "unknown status (invalid)", from: TaskStatus("unknown"), to: TaskStatusTodo, wantErr: true, wantState: TaskStatus("unknown")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			task := Task{Status: tt.from}
			err := task.Transition(tt.to)

			if tt.wantErr {
				if !errors.Is(err, ErrInvalidTransition) {
					t.Fatalf("expected %v, got %v", ErrInvalidTransition, err)
				}

			} else {
				if err != nil {
					t.Fatalf("expected nil error, got %v", err)
				}
			}

			if task.Status != tt.wantState {
				t.Fatalf("expected %q, got %q", tt.wantState, task.Status)
			}

		})
	}
}
