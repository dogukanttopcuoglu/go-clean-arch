package app

import (
	"github.com/dogukanttopcuoglu/clean-lab/config"
	"github.com/dogukanttopcuoglu/clean-lab/internal/controller/restapi"
	"github.com/dogukanttopcuoglu/clean-lab/internal/repo/memory"
	"github.com/dogukanttopcuoglu/clean-lab/internal/usecase/task"
	"github.com/dogukanttopcuoglu/clean-lab/internal/usecase/user"
	"github.com/dogukanttopcuoglu/clean-lab/pkg/httpserver"
	"github.com/dogukanttopcuoglu/clean-lab/pkg/logger"
)

func Run(cfg *config.Config) error {
	appLogger, err := logger.New(logger.Options{
		Environment: cfg.App.Environment,
		Level:       cfg.Logger.Level,
		ServiceName: cfg.Logger.ServiceName,
	})
	if err != nil {
		return err
	}

	defer func() {
		_ = appLogger.Sync()
	}()

	userRepo := memory.NewUserRepo()
	taskRepo := memory.NewTaskRepo()

	userUseCase := user.New(userRepo)
	taskUseCase := task.New(taskRepo)

	httpServer := httpserver.New(
		appLogger,
		httpserver.Port(cfg.HTTP.Port),
	)

	restapi.NewRouter(httpServer.App, userUseCase, taskUseCase, appLogger)
	return httpServer.Start()
}
