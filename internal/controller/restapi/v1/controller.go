package v1

import (
	"github.com/dogukanttopcuoglu/clean-lab/internal/usecase"
	"github.com/dogukanttopcuoglu/clean-lab/pkg/logger"
)

type V1 struct {
	taskUseCase usecase.Task
	userUseCase usecase.User
	log         logger.Logger
}
