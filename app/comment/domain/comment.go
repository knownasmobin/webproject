package domain

import (
	"context"
	"time"
)

const DomainName = "comment"

type Comment struct {
	Id          int
	UserId      int
	BookId      int
	Message     string
	CreatedDate time.Time
	UpdatedDate time.Time
}

type Usecase interface {
	GetCommentById(ctx context.Context, id int) (*Comment, error)
	GetByCondition(ctx context.Context, u Comment) ([]Comment, error)
	Delete(ctx context.Context, commentId int) (*Comment, error)
	Create(ctx context.Context, u Comment) (*Comment, error)
	Update(ctx context.Context, u Comment) (*Comment, error)
}

type Repository interface {
	GetCommentById(ctx context.Context, id int) (*Comment, error)
	GetByCondition(ctx context.Context, u Comment) ([]Comment, error)
	Create(ctx context.Context, comment Comment) (*Comment, error)
	Update(ctx context.Context, condition Comment, data Comment) ([]Comment, error)
	Delete(ctx context.Context, commentId int) (*Comment, error)
}

type Adapter interface {
	SetAdapters()
}
