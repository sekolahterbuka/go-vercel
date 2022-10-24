package service

import (
	"context"

	"github.com/google/uuid"
	pg "github.com/sekolahkita/go-api/server/repository/postgres/sqlc"
)

type authService struct {
	Querier pg.Querier
}

type AuthConfig struct {
	Querier pg.Querier
}

func NewAuthService(c *AuthConfig) *authService {
	return &authService{
		Querier: c.Querier,
	}
}

func (s authService) Login(ctx context.Context, id uuid.UUID) (pg.Auth, error) {

	data, err := s.Querier.Login(ctx, id)
	if err != nil {
		return pg.Auth{}, err
	}

	return data, nil
}

func (s authService) Register(ctx context.Context, arg pg.RegisterParams) (pg.Auth, error) {
	data, err := s.Querier.Register(ctx, pg.RegisterParams{
		Username: arg.Username,
		Email:    arg.Email,
		Password: arg.Password,
	})

	if err != nil {
		return pg.Auth{}, err
	}

	return data, nil
}

func (s authService) Logout(ctx context.Context) error {
	return nil
}
