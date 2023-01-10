package domain

import (
	"context"
	"time"
)

type User struct {
	Id          uint64
	CreatedDate time.Time
	UpdatedDate time.Time
	Roles       *[]string
	Allow       *[]string
	Deny        *[]string
}

type UserUsecase interface {
	GetUserById(ctx context.Context, id uint64) (*User, error)
	Create(ctx context.Context, u User) (*User, error)
	Update(ctx context.Context, u User) (*User, error)
}
type UserRepository interface {
	GetUserById(ctx context.Context, id uint64) (*User, error)
	Create(ctx context.Context, user User) (*User, error)
	Update(ctx context.Context, condition User, data User) ([]User, error)
}
