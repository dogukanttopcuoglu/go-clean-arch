package restapi

import (
	"net/http"

	v1 "github.com/dogukanttopcuoglu/clean-lab/internal/controller/restapi/v1"
	"github.com/dogukanttopcuoglu/clean-lab/internal/usecase"
	"github.com/dogukanttopcuoglu/clean-lab/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, userUseCase usecase.User, taskUseCase usecase.Task, log logger.Logger) {
	app.Get("healthz", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(http.StatusOK)
	})

	apiV1Group := app.Group("v1")
	{
		v1.NewRoutes(apiV1Group, userUseCase, taskUseCase, log)
	}
}
