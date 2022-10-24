package Interface

import (
	"context"

	"github.com/google/uuid"
	pg "github.com/sekolahkita/go-api/server/repository/postgres/sqlc"
)

type AuthService interface {
	Login(ctx context.Context, id uuid.UUID) (pg.Auth, error)
	Register(ctx context.Context, arg pg.RegisterParams) (pg.Auth, error)
	Logout(ctx context.Context) error
}
