package repo

import (
	"context"

	"github.com/dogukanttopcuoglu/go-clean-arch/internal/entity"
)

type TaskRepo interface {
	Store(ctx context.Context, task *entity.Task) error
	GetByID(ctx context.Context, userID, taskID string) (entity.Task, error)
	List(ctx context.Context, userID string, filter TaskFilter) ([]entity.Task, int, error)
	Update(ctx context.Context, task *entity.Task) error
	Delete(ctx context.Context, userID, taskID string) error
}

type TaskFilter struct {
	Status *entity.TaskStatus
	Limit  uint64
	Offset uint64
}
