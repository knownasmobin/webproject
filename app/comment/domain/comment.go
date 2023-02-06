package domain

import (
	"context"
	"time"
)

const DomainName = "comment"

type Comment struct {
	Id          uint64
	CreatedDate time.Time
	UpdatedDate time.Time
}

type Usecase interface {
	GetCommentById(ctx context.Context, id uint64) (*Comment, error)
	Create(ctx context.Context, u Comment) (*Comment, error)
	Update(ctx context.Context, u Comment) (*Comment, error)
}

type Repository interface {
	GetCommentById(ctx context.Context, id uint64) (*Comment, error)
	Create(ctx context.Context, comment Comment) (*Comment, error)
	Update(ctx context.Context, condition Comment, data Comment) ([]Comment, error)
}

type Adapter interface {
	SetAdapters()
}
