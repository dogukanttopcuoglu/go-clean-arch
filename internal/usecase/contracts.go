package usecase

import (
	"context"

	"github.com/dogukanttopcuoglu/go-clean-arch/internal/entity"
)

type Task interface {
	Create(ctx context.Context, userID, title, description string) (entity.Task, error)
	Get(ctx context.Context, userID, taskID string) (entity.Task, error)
	List(ctx context.Context, userID string, status *entity.TaskStatus, limit, offset int) ([]entity.Task, int, error)
	Update(ctx context.Context, userID, taskID, title, description string) (entity.Task, error)
	Transition(ctx context.Context, userID, taskID string, newStatus entity.TaskStatus) (entity.Task, error)
	Delete(ctx context.Context, userID, taskID string) error
}
