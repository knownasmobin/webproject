package domain

import (
	"context"
	"time"
)

const DomainName = "category"

type Category struct {
	Id          int
	EnName      string
	FaName      string
	CreatedDate time.Time
	UpdatedDate time.Time
}

type Usecase interface {
	GetCategoryById(ctx context.Context, id int) (*Category, error)
	GetAll(ctx context.Context) ([]Category, error)
	Create(ctx context.Context, u Category) (*Category, error)
	Update(ctx context.Context, u Category) (*Category, error)
}

type Repository interface {
	GetCategoryById(ctx context.Context, id int) (*Category, error)
	GetAll(ctx context.Context) ([]Category, error)
	Create(ctx context.Context, category Category) (*Category, error)
	Update(ctx context.Context, condition Category, data Category) ([]Category, error)
}

type Adapter interface {
	SetAdapters()
}
