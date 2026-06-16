package task_test

import (
	"context"
	"testing"

	"github.com/dogukanttopcuoglu/clean-lab/internal/entity"
	"github.com/dogukanttopcuoglu/clean-lab/internal/repo/memory"
	"github.com/dogukanttopcuoglu/clean-lab/internal/usecase/task"
)

func TestUseCase_Create(t *testing.T) {
	ctx := context.Background()
	repo := memory.NewTaskRepo()
	uc := task.New(repo)

	created, err := uc.Create(ctx, "user-1", "Learn Go", "Write usecase test")
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}

	if created.ID == "" {
		t.Fatalf("expected some uuid val, got %q", created.ID)
	}

	if created.UserID != "user-1" {
		t.Fatalf("expected user-1 , got %q", created.UserID)
	}

	if created.Title != "Learn Go" {
		t.Fatalf("expected Learn Go, got %q", created.Title)
	}

	if created.Description != "Write usecase test" {
		t.Fatalf("expected Write usecase test, got %q", created.Description)
	}

	if created.Status != entity.TaskStatusTodo {
		t.Fatalf("expected TaskStatusTodo, got %q", created.Status)
	}

	if created.CreatedAt.IsZero() {
		t.Fatal("expected time now, got zero")
	}

	if created.UpdatedAt.IsZero() {
		t.Fatalf("expected time now, got zero")
	}

	saved, err := repo.GetByID(ctx, "user-1", created.ID)
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
	if saved.ID != created.ID {
		t.Fatalf("expected saved taskID to match created taskID ")
	}

}

func TestUseCase_Transition(t *testing.T) {
	ctx := context.Background()

	repo := memory.NewTaskRepo()

	uc := task.New(repo)

	created, err := uc.Create(ctx, "user-1", "Learn Go", "Write usecase test")
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}

	if created.Status != entity.TaskStatusTodo {
		t.Fatalf("expected TaskStatusTodo, got %q", created.Status)
	}

	transitionedTask, err := uc.Transition(ctx, created.UserID, created.ID, entity.TaskStatusInProgress)
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}

	if transitionedTask.Status != entity.TaskStatusInProgress {
		t.Fatalf("expected TaskStatusInProgress, got %q", transitionedTask.Status)
	}

	if transitionedTask.UpdatedAt.IsZero() {
		t.Fatal("expected updated time to be set")
	}

	saved, err := repo.GetByID(ctx, transitionedTask.UserID, transitionedTask.ID)
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}

	if saved.Status != entity.TaskStatusInProgress {
		t.Fatalf("expected TaskStatusInProgress, got %q", saved.Status)
	}

}
