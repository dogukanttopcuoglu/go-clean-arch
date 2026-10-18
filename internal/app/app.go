package app

import (
	persistTaskRepo "github.com/dogukanttopcuoglu/go-clean-arch/internal/repo/persistent/task"
	"github.com/dogukanttopcuoglu/go-clean-arch/internal/usecase"
	taskUseCase "github.com/dogukanttopcuoglu/go-clean-arch/internal/usecase/task"
	"github.com/dogukanttopcuoglu/go-clean-arch/pkg/postgres"
)

type useCases struct {
	task usecase.Task
}

func initUseCases(pg *postgres.Postgres) useCases {
	taskRepo := persistTaskRepo.New(pg)

	return useCases{
		task: taskUseCase.New(taskRepo),
	}
}
