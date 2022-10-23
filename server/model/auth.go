package model

import (
	"context"

	"github.com/google/uuid"
)

type Auth struct {
	UID      uuid.UUID `json:"uid"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

type RegisterParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type AuthSerializer interface {
	Decode(input []byte) (*Auth, error)
	Encode(input *Auth) ([]byte, error)
}
type RegisterSerializer interface {
	Decode(input []byte) (*RegisterParams, error)
	Encode(input *RegisterParams) ([]byte, error)
}

type AuthService interface {
	Login(ctx context.Context, id uuid.UUID) (Auth, error)
	Register(ctx context.Context, arg RegisterParams) (Auth, error)
}

type AuthRepository interface {
	Login(ctx context.Context, id uuid.UUID) (Auth, error)
	Register(ctx context.Context, arg RegisterParams) (Auth, error)
}
