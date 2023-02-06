package domain

import (
	"context"
	"time"
)

const DomainName = "user"

type User struct {
	Id          uint64
	CreatedDate time.Time
	UpdatedDate time.Time
	Roles       []string
	Allow       []string
	Deny        []string
}

type Usecase interface {
	GetUserById(ctx context.Context, id uint64) (*User, error)
	Create(ctx context.Context, u User) (*User, error)
	Update(ctx context.Context, u User) (*User, error)
}

type Repository interface {
	GetUserById(ctx context.Context, id uint64) (*User, error)
	Create(ctx context.Context, user User) (*User, error)
	Update(ctx context.Context, condition User, data User) ([]User, error)
}

type Adapter interface {
	SetAdapters(bazAdapter BazAdapter)
}
type BazAdapter interface {
	Create(ctx context.Context, user User) error
}

type FooAdapter interface {
	Bar(ctx context.Context, user User) error
}
