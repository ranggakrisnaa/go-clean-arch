package v1

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ranggakrisna/go-clean-arch/internal/usecase"
	"github.com/ranggakrisna/go-clean-arch/pkg/logger"
)

// NewTranslationRoutes -.
func NewTranslationRoutes(apiV1Group fiber.Router, t usecase.Translation, l logger.Interface) {
	r := &V1{t: t, l: l, v: validator.New(validator.WithRequiredStructEnabled())}

	translationGroup := apiV1Group.Group("/translation")

	{
		translationGroup.Get("/history", r.history)
		translationGroup.Post("/do-translate", r.doTranslate)
	}
}

func NewAuthRoutes(apiV1Group fiber.Router, au *usecase.Auth, l logger.Interface) {
	r := &V1{au: *au, l: l, v: validator.New(validator.WithRequiredStructEnabled())}

	authGroup := apiV1Group.Group("/auth")

	{
		authGroup.Post("/register", r.Register)
	}
}
