package v1

import (
	v1 "github.com/ranggakrisna/go-clean-arch/docs/proto/v1"
	"github.com/ranggakrisna/go-clean-arch/internal/usecase"
	"github.com/ranggakrisna/go-clean-arch/pkg/logger"
	"github.com/go-playground/validator/v10"
	pbgrpc "google.golang.org/grpc"
)

// NewTranslationRoutes -.
func NewTranslationRoutes(app *pbgrpc.Server, t usecase.Translation, l logger.Interface) {
	r := &V1{t: t, l: l, v: validator.New(validator.WithRequiredStructEnabled())}

	{
		v1.RegisterTranslationServer(app, r)
	}
}
