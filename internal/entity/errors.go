package entity

import "errors"

var (
	ErrUserNotFound       = errors.New("user note found")
	ErrUserAlreadyExist   = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")

	ErrTaskNotFound      = errors.New("task not found")
	ErrTaskForbidden     = errors.New("task forbidden")
	ErrInvalidTransation = errors.New("invalid task transition")
)
