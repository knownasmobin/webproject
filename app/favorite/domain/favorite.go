package domain

import (
	"context"
	"time"
)

const DomainName = "favorite"

type Favorite struct {
	Id          int
	BookId      int
	UserId      int
	CreatedDate time.Time
	UpdatedDate time.Time
}

type Usecase interface {
	GetByCondition(ctx context.Context, u Favorite) ([]Favorite, error)
	GetMostFavorites(ctx context.Context) ([]Favorite, error)
	Create(ctx context.Context, u Favorite) (*Favorite, error)
	Delete(ctx context.Context, id int) (*Favorite, error)
}

type Repository interface {
	GetByCondition(ctx context.Context, u Favorite) ([]Favorite, error)
	GetMostFavorites(ctx context.Context) ([]Favorite, error)
	Create(ctx context.Context, favorite Favorite) (*Favorite, error)
	Delete(ctx context.Context, id int) (*Favorite, error)
}

type Adapter interface {
	SetAdapters()
}
