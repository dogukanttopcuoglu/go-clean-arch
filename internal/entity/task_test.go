package entity_test

import (
	"testing"

	"github.com/dogukanttopcuoglu/go-clean-arch/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTask_Transition(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		from      entity.TaskStatus
		to        entity.TaskStatus
		wantErr   bool
		wantState entity.TaskStatus
	}{
		{"todo_to_in_progress", entity.TaskStatusTodo, entity.TaskStatusInProgress, false, entity.TaskStatusInProgress},
		{"in_progress_to_done", entity.TaskStatusInProgress, entity.TaskStatusDone, false, entity.TaskStatusDone},
		{"in_progress_to_todo", entity.TaskStatusInProgress, entity.TaskStatusTodo, false, entity.TaskStatusTodo},
		{"todo_to_done_invalid", entity.TaskStatusTodo, entity.TaskStatusDone, true, entity.TaskStatusTodo},
		{"done_to_todo_invalid", entity.TaskStatusDone, entity.TaskStatusTodo, true, entity.TaskStatusDone},
		{"invalid_status_to_todo", entity.TaskStatus("invalid"), entity.TaskStatusTodo, true, entity.TaskStatus("invalid")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			task := entity.Task{Status: tt.from}
			err := task.Transition(tt.to)
			if tt.wantErr {
				require.ErrorIs(t, err, entity.ErrInvalidTransition)
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, tt.wantState, task.Status)
		})
	}
}
