package model

import (
	"context"

	"github.com/google/uuid"
)

type User struct {
	Name string `json:"name"`
}

type UserSerializer interface {
	Decode(input []byte) (*User, error)
	Encode(input *User) ([]byte, error)
}

type UserService interface{}

type UserRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUsername(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, u *User) (*User, error)
	Update(ctx context.Context, u *User) (*User, error)
	Delete(ctx context.Context, u *User) error
}
