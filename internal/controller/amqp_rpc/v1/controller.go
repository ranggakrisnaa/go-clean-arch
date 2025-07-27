package v1

import (
	"github.com/ranggakrisna/go-clean-arch/internal/usecase"
	"github.com/ranggakrisna/go-clean-arch/pkg/logger"
	"github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
	t usecase.Translation
	l logger.Interface
	v *validator.Validate
}
