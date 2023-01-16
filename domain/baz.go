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

type BazUsecase interface {
	Create(ctx context.Context, baz Baz) (*Baz, error)
}
type BazRepository interface {
	Create(ctx context.Context, baz Baz) (*Baz, error)
}
