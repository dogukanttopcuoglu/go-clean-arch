package memory

import (
	"context"

	"github.com/dogukanttopcuoglu/clean-lab/internal/entity"
	"github.com/dogukanttopcuoglu/clean-lab/internal/repo"
)

type TaskRepo struct {
	tasksByID map[string]entity.Task
}

func NewTaskRepo() *TaskRepo {
	return &TaskRepo{
		tasksByID: make(map[string]entity.Task),
	}
}

func (r *TaskRepo) Store(ctx context.Context, task *entity.Task) error {
	_ = ctx
	if _, ok := r.tasksByID[task.ID]; ok {
		return entity.ErrTaskAlreadyExists
	}

	r.tasksByID[task.ID] = *task
	return nil
}

func (r *TaskRepo) GetByID(ctx context.Context, userID, taskID string) (entity.Task, error) {
	_ = ctx
	task, ok := r.tasksByID[taskID]
	if !ok {
		return entity.Task{}, entity.ErrTaskNotFound
	}

	if task.UserID != userID {
		return entity.Task{}, entity.ErrTaskForbidden
	}

	return task, nil

}

func (r *TaskRepo) List(ctx context.Context, userID string, filter repo.TaskFilter) ([]entity.Task, int, error) {
	_ = ctx
	matchedTasks := []entity.Task{}

	for _, task := range r.tasksByID {
		if task.UserID != userID {
			continue
		}
		if filter.Status != nil {
			if task.Status != *filter.Status {
				continue
			}
		}

		matchedTasks = append(matchedTasks, task)

	}

	total := len(matchedTasks)
	start := int(filter.Offset)

	if start > total {
		return []entity.Task{}, total, nil
	}

	end := start + int(filter.Limit)

	if filter.Limit == 0 {
		end = total
	}

	if end > total {
		end = total
	}

	return matchedTasks[start:end], total, nil
}

func (r *TaskRepo) Update(ctx context.Context, task *entity.Task) error {
	_ = ctx
	existingTask, ok := r.tasksByID[task.ID]
	if !ok {
		return entity.ErrTaskNotFound
	}

	if existingTask.UserID != task.UserID {
		return entity.ErrTaskForbidden
	}

	r.tasksByID[task.ID] = *task
	return nil

}

func (r *TaskRepo) Delete(ctx context.Context, userID, taskID string) error {
	_ = ctx

	existingTask, ok := r.tasksByID[taskID]
	if !ok {
		return entity.ErrTaskNotFound
	}

	if existingTask.UserID != userID {
		return entity.ErrTaskForbidden
	}

	delete(r.tasksByID, taskID)
	return nil
}
