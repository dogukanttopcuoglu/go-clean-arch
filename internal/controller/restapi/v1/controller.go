package v1

import (
	"github.com/dogukanttopcuoglu/go-clean-arch/internal/usecase"
	"github.com/go-playground/validator/v10"
)

type V1 struct {
	tk usecase.Task
	v  *validator.Validate
}
