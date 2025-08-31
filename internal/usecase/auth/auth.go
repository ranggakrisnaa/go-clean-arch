package auth

import (
	"context"
	"fmt"

	"github.com/ranggakrisna/go-clean-arch/internal/entity"
	"github.com/ranggakrisna/go-clean-arch/internal/repo/persistent"
	"github.com/ranggakrisna/go-clean-arch/pkg/bcrypt"
)

type UseCase struct {
	repo *persistent.AuthRepository
}

func NewUseCase(repo *persistent.AuthRepository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (uc *UseCase) Register(ctx context.Context, payload *entity.Register) (entity.User, error) {
	hashed, err := bcrypt.HashPassword(payload.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("UseCase - Register - bcrypt.HashPassword: %w", err)
	}

	payload.Password = hashed

	id, err := uc.repo.Register(ctx, payload)
	if err != nil {
		return entity.User{}, fmt.Errorf("UseCase - Register - uc.repo.Register: %w", err)
	}

	return entity.User{ID: id}, nil
}
