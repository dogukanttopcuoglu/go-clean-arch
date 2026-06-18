package app

import (
	"github.com/dogukanttopcuoglu/clean-lab/config"
	"github.com/dogukanttopcuoglu/clean-lab/internal/controller/restapi"
	"github.com/dogukanttopcuoglu/clean-lab/internal/repo/memory"
	"github.com/dogukanttopcuoglu/clean-lab/internal/usecase/task"
	"github.com/dogukanttopcuoglu/clean-lab/internal/usecase/user"
	"github.com/dogukanttopcuoglu/clean-lab/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func Run(cfg *config.Config) error {
	fiberApp := fiber.New()

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

	restapi.NewRouter(fiberApp, userUseCase, taskUseCase, appLogger)

	addr := ":" + cfg.HTTP.Port

	appLogger.Info("http server listening on " + addr)
	return fiberApp.Listen(addr)
}
