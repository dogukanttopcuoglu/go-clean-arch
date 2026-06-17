package request

import "github.com/dogukanttopcuoglu/clean-lab/internal/entity"

type CreateTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TranstitionTask struct {
	Status entity.TaskStatus `json:"status"`
}
