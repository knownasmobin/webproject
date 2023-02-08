package domain

import (
	"context"
	"time"
)

const DomainName = "user"

type User struct {
	Id          int
	Name        string
	Username    string
	Password    string
	IsAdmin     bool
	CreatedDate time.Time
	UpdatedDate time.Time
}

type Usecase interface {
	GetUserById(ctx context.Context, id int) (*User, error)
	LoginUserCredential(ctx context.Context, user User) (*User, string, error)
	GetByCondition(ctx context.Context, user User) ([]User, error)
	Create(ctx context.Context, u User) (*User, error)
	Update(ctx context.Context, u User) (*User, error)
	ValidateToken(ctx context.Context, tokenStr string) (*User, error)
}

type Repository interface {
	GetUserByPassword(ctx context.Context, username, password string) (*User, error)
	GetUserById(ctx context.Context, id int) (*User, error)
	GetByCondition(ctx context.Context, user User) ([]User, error)
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
