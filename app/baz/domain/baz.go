package domain

import (
	"context"
	"time"
)

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
}

type Adapter interface {
	SetAdapters()
}
