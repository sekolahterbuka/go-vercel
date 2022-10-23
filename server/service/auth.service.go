package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sekolahkita/go-api/server/model"
)

type authService struct {
	// AuthRepository model.AuthRepository
}

type AuthConfig struct {
	// AuthRepository model.AuthRepository
}

func NewAuthService(c *AuthConfig) *authService {
	return &authService{}
}

func (s authService) Register(ctx context.Context, arg model.RegisterParams) (model.Auth, error) {
	// user, err := s.AuthRepository.Register(ctx, arg)

	// if err != nil {
	// 	return model.Auth{}, err
	// }
	user := model.Auth{}
	return user, nil
}

func (s authService) Login(ctx context.Context, id uuid.UUID) (model.Auth, error) {
	// user, err := s.AuthRepository.Login(ctx, id)

	// if err != nil {
	// 	return model.Auth{}, err
	// }
	user := model.Auth{}

	return user, nil

}
