package domain

import (
	"context"
	"time"
)

const DomainName = "purchase"

type Purchase struct {
	Id          int
	UserId      int
	BookId      int
	Price       float32
	CreatedDate time.Time
	UpdatedDate time.Time
}

type Usecase interface {
	GetByCondition(ctx context.Context, u Purchase) ([]Purchase, error)
	Create(ctx context.Context, u Purchase) (*Purchase, error)
}

type Repository interface {
	GetByCondition(ctx context.Context, u Purchase) ([]Purchase, error)
	Create(ctx context.Context, u Purchase) (*Purchase, error)
}

type Adapter interface {
	SetAdapters()
}
