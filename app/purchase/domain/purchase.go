package domain

import (
	"context"
	"time"
)

const DomainName = "purchase"

type Purchase struct {
	Id          uint64
	CreatedDate time.Time
	UpdatedDate time.Time
}

type Usecase interface {
	GetPurchaseById(ctx context.Context, id uint64) (*Purchase, error)
	Create(ctx context.Context, u Purchase) (*Purchase, error)
	Update(ctx context.Context, u Purchase) (*Purchase, error)
}

type Repository interface {
	GetPurchaseById(ctx context.Context, id uint64) (*Purchase, error)
	Create(ctx context.Context, purchase Purchase) (*Purchase, error)
	Update(ctx context.Context, condition Purchase, data Purchase) ([]Purchase, error)
}

type Adapter interface {
	SetAdapters()
}
