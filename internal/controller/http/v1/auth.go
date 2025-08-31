package v1

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ranggakrisna/go-clean-arch/internal/controller/http/v1/request"
	"github.com/ranggakrisna/go-clean-arch/internal/entity"
)

func (r *V1) Register(ctx *fiber.Ctx) error {
	var body request.Register

	if err := ctx.BodyParser(&body); err != nil {
		r.l.Error(err, "http - v1 - Register")

		return errorResponse(ctx, fiber.StatusBadRequest, fmt.Sprintf("invalid request body: %v", err.Error()))
	}

	if err := r.v.Struct(body); err != nil {
		r.l.Error(err, "http - v1 - Register")

		return errorResponse(ctx, http.StatusBadRequest, fmt.Sprintf("invalid request body %v", err.Error()))
	}

	user, err := r.au.Register(ctx.UserContext(), &entity.Register{
		Email:    body.Email,
		Password: body.Password,
		FullName: body.FullName,
		Phone:    body.Phone,
	})
	if err != nil {
		r.l.Error(err, "http - v1 - Register")

		return errorResponse(ctx, fiber.StatusInternalServerError, "failed to register user")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"id":      user.ID,
		"message": "user registered successfully",
	})
}
