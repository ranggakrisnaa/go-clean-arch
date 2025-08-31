package persistent

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ranggakrisna/go-clean-arch/internal/entity"
	"github.com/ranggakrisna/go-clean-arch/pkg/postgres"
)

type AuthRepository struct {
	*postgres.Postgres
}

func NewAuthRepository(authPg *postgres.Postgres) *AuthRepository {
	return &AuthRepository{
		Postgres: authPg,
	}
}

func (r *AuthRepository) Register(ctx context.Context, a *entity.Register) (uuid.UUID, error) {
	sql, args, err := r.Builder.
		Insert("users").
		Columns("password", "email", "fullname", "phone_number").
		Values(a.Password, a.Email, a.FullName, a.Phone).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return uuid.Nil, fmt.Errorf("AuthRepo - Register - r.Builder: %w", err)
	}

	var id uuid.UUID
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("AuthRepo - Register - r.Pool.QueryRow: %w", err)
	}

	return id, nil
}
