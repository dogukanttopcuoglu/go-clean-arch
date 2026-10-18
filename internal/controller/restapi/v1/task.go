package v1

import (
	"net/http"

	"github.com/dogukanttopcuoglu/go-clean-arch/internal/controller/restapi/v1/request"
	"github.com/gofiber/fiber/v2"
)

func (r *V1) createTask(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(string)

	if !ok {
		return errorResponse(ctx, http.StatusUnauthorized, "unauthorized")
	}

	var body request.CreateTask

	if err := ctx.BodyParser(&body); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	if err := r.v.Struct(body); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	task, err := r.tk.Create(ctx.UserContext(), userID, body.Title, body.Description)
	if err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "internal server error")
	}

	return ctx.Status(http.StatusCreated).JSON(task)
}
