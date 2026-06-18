package v1

import (
	"github.com/dogukanttopcuoglu/clean-lab/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func NewRoutes(apiV1Group fiber.Router, u usecase.User, tk usecase.Task) {
	r := &V1{
		taskUseCase: tk,
		userUseCase: u,
	}

	taskGroup := apiV1Group.Group("/tasks")
	{
		taskGroup.Post("/", r.createTask)
		taskGroup.Get("/", r.listTasks)
		taskGroup.Get("/:id", r.getTask)
		taskGroup.Put("/:id", r.updateTask)
		taskGroup.Patch("/:id/status", r.transitionTask)
		taskGroup.Delete("/:id", r.deleteTask)

	}

	authGroup := apiV1Group.Group("/auth")
	{
		authGroup.Post("/register", r.register)
		authGroup.Post("/login", r.login)
	}

}
