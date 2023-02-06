package domain

import (
	"context"
	"time"
)

const DomainName = "favorite"

type Favorite struct {
	Id          uint64
	CreatedDate time.Time
	UpdatedDate time.Time
}

type Usecase interface {
	GetFavoriteById(ctx context.Context, id uint64) (*Favorite, error)
	Create(ctx context.Context, u Favorite) (*Favorite, error)
	Update(ctx context.Context, u Favorite) (*Favorite, error)
}

type Repository interface {
	GetFavoriteById(ctx context.Context, id uint64) (*Favorite, error)
	Create(ctx context.Context, favorite Favorite) (*Favorite, error)
	Update(ctx context.Context, condition Favorite, data Favorite) ([]Favorite, error)
}

type Adapter interface {
	SetAdapters()
}
