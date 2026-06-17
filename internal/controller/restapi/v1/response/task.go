package response

import "github.com/dogukanttopcuoglu/clean-lab/internal/entity"

type TaskList struct {
	Tasks []entity.Task `json:"tasks"`
	Total int           `json:"total" example:"42"`
}
