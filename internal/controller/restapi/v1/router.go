package v1

import (
	"github.com/dogukanttopcuoglu/go-clean-arch/internal/controller/restapi/middleware"
	"github.com/dogukanttopcuoglu/go-clean-arch/internal/usecase"
	"github.com/dogukanttopcuoglu/go-clean-arch/pkg/jwt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewRoutes(apiV1Group fiber.Router, tk usecase.Task, jwtManager *jwt.Manager) {
	r := &V1{
		tk: tk,
		v:  validator.New(validator.WithRequiredStructEnabled()),
	}

	protected := apiV1Group.Group("", middleware.Auth(jwtManager))

	taskGroup := protected.Group("tasks")

	taskGroup.Post("/", r.createTask)

}
