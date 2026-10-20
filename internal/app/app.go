package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dogukanttopcuoglu/go-clean-arch/config"
	"github.com/dogukanttopcuoglu/go-clean-arch/internal/controller/restapi"
	persistTaskRepo "github.com/dogukanttopcuoglu/go-clean-arch/internal/repo/persistent/task"
	"github.com/dogukanttopcuoglu/go-clean-arch/internal/usecase"
	taskUseCase "github.com/dogukanttopcuoglu/go-clean-arch/internal/usecase/task"
	"github.com/dogukanttopcuoglu/go-clean-arch/pkg/httpserver"
	"github.com/dogukanttopcuoglu/go-clean-arch/pkg/jwt"
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

func waitForShutdown(server *httpserver.Server) error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	var serverErr error

	select {
	case <-interrupt:
	case serverErr = <-server.Notify():
	}

	shutdownErr := server.Shutdown()

	if serverErr != nil {
		return fmt.Errorf("app - waitForShutdown - server.Notify: %w", serverErr)

	}

	if shutdownErr != nil {
		return fmt.Errorf("app - waitForShutdown - server.Shutdown: %w", shutdownErr)
	}
	return nil
}

func Run(cfg *config.Config) error {
	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		return fmt.Errorf("app - Run - postgres.New: %w", err)
	}
	defer pg.Close()

	// JWT

	jwtManager := jwt.New(cfg.JWT.Secret, cfg.JWT.TokenExpiry)

	uc := initUseCases(pg)
	server := httpserver.New(httpserver.Port(cfg.HTTP.Port))
	restapi.NewRouter(server.App, uc.task, jwtManager)

	server.Start()

	return waitForShutdown(server)

}
