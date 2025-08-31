// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/ranggakrisna/go-clean-arch/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_usecase_test.go -package=usecase_test

type (
	// Translation -.
	Translation interface {
		Translate(context.Context, entity.Translation) (entity.Translation, error)
		History(context.Context) (entity.TranslationHistory, error)
	}

	Auth interface {
		// Login(context.Context, entity.Login) (entity.User, error)
		Register(context.Context, *entity.Register) (entity.User, error)
	}
)
