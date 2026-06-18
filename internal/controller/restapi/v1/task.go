package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dogukanttopcuoglu/clean-lab/internal/controller/restapi/v1/request"
	"github.com/dogukanttopcuoglu/clean-lab/internal/controller/restapi/v1/response"
	"github.com/dogukanttopcuoglu/clean-lab/internal/entity"
	"github.com/gofiber/fiber/v2"
)

const temporaryUserID = "user-1"

func (r *V1) createTask(ctx *fiber.Ctx) error {
	var body request.CreateTask

	if err := ctx.BodyParser(&body); err != nil {
		r.log.Warn(err, "restapi - v1 - create task - invalid request body")
		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	task, err := r.taskUseCase.Create(ctx.UserContext(), temporaryUserID, body.Title, body.Description)
	if err != nil {
		r.log.Error(err, "restapi - v1 - create task - unexpected error")
		return errorResponse(ctx, http.StatusInternalServerError, "internal server error")
	}

	return ctx.Status(http.StatusCreated).JSON(task)
}

func (r *V1) listTasks(ctx *fiber.Ctx) error {
	var status *entity.TaskStatus

	if rawStatus := ctx.Query("status"); rawStatus != "" {
		taskStatus := entity.TaskStatus(rawStatus)
		if !taskStatus.Valid() {
			return errorResponse(ctx, http.StatusBadRequest, "invalid task status")
		}

		status = &taskStatus
	}

	limit, err := strconv.Atoi(ctx.Query("limit", "10"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(ctx.Query("offset", "0"))
	if err != nil {
		offset = 0
	}

	tasks, total, err := r.taskUseCase.List(
		ctx.UserContext(),
		temporaryUserID,
		status,
		limit,
		offset,
	)
	if err != nil {
		r.log.Error(err, "restapi - v1 - list tasks - unexpected error")
		return errorResponse(ctx, http.StatusInternalServerError, "internal server error")
	}

	return ctx.Status(http.StatusOK).JSON(response.TaskList{
		Tasks: tasks,
		Total: total,
	})
}

func (r *V1) getTask(ctx *fiber.Ctx) error {
	taskID := ctx.Params("id")

	task, err := r.taskUseCase.Get(ctx.UserContext(), temporaryUserID, taskID)
	if err != nil {
		if errors.Is(err, entity.ErrTaskNotFound) {
			return errorResponse(ctx, http.StatusNotFound, "task not found")
		}

		if errors.Is(err, entity.ErrTaskForbidden) {
			return errorResponse(ctx, http.StatusForbidden, "forbidden")
		}
		r.log.Error(err, "restapi - v1 - get task - unexpected error")
		return errorResponse(ctx, http.StatusInternalServerError, "internal server error")
	}

	return ctx.Status(http.StatusOK).JSON(task)
}

func (r *V1) updateTask(ctx *fiber.Ctx) error {
	taskID := ctx.Params("id")

	var body request.UpdateTask

	if err := ctx.BodyParser(&body); err != nil {
		r.log.Warn(err, "restapi - v1 - update task - invalid request body")
		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	task, err := r.taskUseCase.Update(
		ctx.UserContext(),
		temporaryUserID,
		taskID,
		body.Title,
		body.Description,
	)
	if err != nil {
		if errors.Is(err, entity.ErrTaskNotFound) {
			return errorResponse(ctx, http.StatusNotFound, "task not found")
		}

		if errors.Is(err, entity.ErrTaskForbidden) {
			return errorResponse(ctx, http.StatusForbidden, "forbidden")
		}

		return errorResponse(ctx, http.StatusInternalServerError, "internal server error")
	}

	return ctx.Status(http.StatusOK).JSON(task)
}

func (r *V1) transitionTask(ctx *fiber.Ctx) error {
	taskID := ctx.Params("id")

	var body request.TranstitionTask

	if err := ctx.BodyParser(&body); err != nil {
		r.log.Warn(err, "restapi - v1 - transition task - invalid request body")
		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	if !body.Status.Valid() {
		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	task, err := r.taskUseCase.Transition(
		ctx.UserContext(),
		temporaryUserID,
		taskID,
		body.Status,
	)
	if err != nil {
		if errors.Is(err, entity.ErrTaskNotFound) {
			return errorResponse(ctx, http.StatusNotFound, "task not found")
		}

		if errors.Is(err, entity.ErrTaskForbidden) {
			return errorResponse(ctx, http.StatusForbidden, "forbidden")
		}

		if errors.Is(err, entity.ErrInvalidTransition) {
			return errorResponse(ctx, http.StatusBadRequest, "invalid status transition")
		}
		r.log.Error(err, "restapi - v1 - transition task - unexpected error")
		return errorResponse(ctx, http.StatusInternalServerError, "internal server error")
	}

	return ctx.Status(http.StatusOK).JSON(task)
}

func (r *V1) deleteTask(ctx *fiber.Ctx) error {
	taskID := ctx.Params("id")

	err := r.taskUseCase.Delete(ctx.UserContext(), temporaryUserID, taskID)
	if err != nil {
		if errors.Is(err, entity.ErrTaskNotFound) {
			return errorResponse(ctx, http.StatusNotFound, "task not found")
		}

		if errors.Is(err, entity.ErrTaskForbidden) {
			return errorResponse(ctx, http.StatusForbidden, "forbidden")
		}
		r.log.Error(err, "restapi - v1 - delete task - unexpected error")
		return errorResponse(ctx, http.StatusInternalServerError, "internal server error")
	}

	return ctx.SendStatus(http.StatusNoContent)
}
