package v1

import (
	"github.com/go-playground/validator/v10"
	"github.com/ranggakrisna/go-clean-arch/internal/usecase"
	"github.com/ranggakrisna/go-clean-arch/pkg/logger"
)

// V1 -.
type V1 struct {
	t  usecase.Translation
	au usecase.Auth
	l  logger.Interface
	v  *validator.Validate
}
