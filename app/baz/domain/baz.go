package domain

import (
	"context"
	"time"
)

const DomainName = "baz"

type Baz struct {
	UserId      uint64
	CreatedDate time.Time
	UpdatedDate time.Time
}

type Usecase interface {
	Create(ctx context.Context, baz Baz) (*Baz, error)
}
type Repository interface {
	Create(ctx context.Context, baz Baz) (*Baz, error)
	Update(ctx context.Context, condition, baz Baz) ([]Baz, error)
	Get(ctx context.Context, userId uint64) (*Baz, error)
}

type Adapter interface {
	SetAdapters()
}
