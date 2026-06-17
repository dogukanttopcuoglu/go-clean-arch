package app

import (
	"github.com/dogukanttopcuoglu/clean-lab/internal/controller/restapi"
	"github.com/dogukanttopcuoglu/clean-lab/internal/repo/memory"
	"github.com/dogukanttopcuoglu/clean-lab/internal/usecase/task"
	"github.com/dogukanttopcuoglu/clean-lab/internal/usecase/user"
	"github.com/gofiber/fiber/v2"
)

func Run() error {
	fiberApp := fiber.New()

	userRepo := memory.NewUserRepo()
	taskRepo := memory.NewTaskRepo()

	userUseCase := user.New(userRepo)
	taskUseCase := task.New(taskRepo)

	restapi.NewRouter(fiberApp, userUseCase, taskUseCase)

	return fiberApp.Listen(":8080")
}
