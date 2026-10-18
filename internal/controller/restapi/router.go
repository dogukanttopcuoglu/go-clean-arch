package restapi

import (
	v1 "github.com/dogukanttopcuoglu/go-clean-arch/internal/controller/restapi/v1"
	"github.com/dogukanttopcuoglu/go-clean-arch/internal/usecase"
	"github.com/dogukanttopcuoglu/go-clean-arch/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, tk usecase.Task, jwtManager *jwt.Manager) {
	apiv1Group := app.Group("/v1")

	v1.NewRoutes(apiv1Group, tk, jwtManager)
}
