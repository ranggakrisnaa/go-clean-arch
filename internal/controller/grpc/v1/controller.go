package v1

import (
	v1 "github.com/ranggakrisna/go-clean-arch/docs/proto/v1"
	"github.com/ranggakrisna/go-clean-arch/internal/usecase"
	"github.com/ranggakrisna/go-clean-arch/pkg/logger"
	"github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
	v1.TranslationServer

	t usecase.Translation
	l logger.Interface
	v *validator.Validate
}
