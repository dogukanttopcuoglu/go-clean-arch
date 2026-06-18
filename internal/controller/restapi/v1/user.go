package v1

import (
	"errors"
	"net/http"

	"github.com/dogukanttopcuoglu/clean-lab/internal/controller/restapi/v1/request"
	"github.com/dogukanttopcuoglu/clean-lab/internal/controller/restapi/v1/response"
	"github.com/dogukanttopcuoglu/clean-lab/internal/entity"
	"github.com/gofiber/fiber/v2"
)

func (r *V1) register(ctx *fiber.Ctx) error {
	var body request.Register

	if err := ctx.BodyParser(&body); err != nil {
		r.log.Error(err, "restapi - v1 - register - invalid request body")
		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}
	user, err := r.userUseCase.Register(
		ctx.UserContext(),
		body.Username,
		body.Email,
		body.Password,
	)
	if err != nil {
		if errors.Is(err, entity.ErrInvalidInput) {
			return errorResponse(ctx, http.StatusBadRequest, "invalid input")
		}

		if errors.Is(err, entity.ErrUserAlreadyExists) {
			return errorResponse(ctx, http.StatusBadRequest, "could not register with provided credentials")
		}

		return errorResponse(ctx, http.StatusInternalServerError, "internal server error")
	}

	return ctx.Status(http.StatusCreated).JSON(user)
}

func (r *V1) login(ctx *fiber.Ctx) error {
	var body request.Login

	if err := ctx.BodyParser(&body); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid request")
	}

	token, err := r.userUseCase.Login(
		ctx.UserContext(),
		body.Email,
		body.Password,
	)

	if err != nil {
		if errors.Is(err, entity.ErrInvalidCredentials) {
			return errorResponse(ctx, http.StatusUnauthorized, "invalid credentials")
		}

		return errorResponse(ctx, http.StatusInternalServerError, "internal server error")
	}

	return ctx.Status(http.StatusOK).JSON(response.Token{
		Token: token,
	})
}
