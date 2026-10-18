package request

import "github.com/dogukanttopcuoglu/go-clean-arch/internal/entity"

type CreateTask struct {
	Title       string `json:"title" validate:"required,max=255" example:"My task"`
	Description string `json:"description" validate:"max=1000" example:"Task description"`
}

type UpdateTask struct {
	Title       string `json:"title" validate:"required,max=255" example:"Updated task"`
	Description string `json:"description" validate:"max=1000" example:"Updated description"`
}

type TransitionTask struct {
	Status entity.TaskStatus `json:"status" validate:"required,oneof=todo in_progress done" example:"in_progress"`
}
