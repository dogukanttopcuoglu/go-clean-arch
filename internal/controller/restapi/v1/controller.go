package v1

import "github.com/dogukanttopcuoglu/clean-lab/internal/usecase"

type V1 struct {
	taskUseCase usecase.Task
	userUseCase usecase.User
}
