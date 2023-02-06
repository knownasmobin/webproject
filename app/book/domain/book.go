package domain

import (
	"context"
	"time"
)

const DomainName = "book"

type Book struct {
	Id          int
	Title       string
	Price       float32
	Author      string
	CategoryId  int
	CreatedDate time.Time
	UpdatedDate time.Time
}

type Usecase interface {
	GetBookById(ctx context.Context, id int) (*Book, error)
	GetAll(ctx context.Context, categoryId *int) ([]Book, error)
	Create(ctx context.Context, u Book) (*Book, error)
	Update(ctx context.Context, u Book) (*Book, error)
}

type Repository interface {
	GetBookById(ctx context.Context, id int) (*Book, error)
	GetByCategory(ctx context.Context, categoryId *int) ([]Book, error)
	Create(ctx context.Context, book Book) (*Book, error)
	Update(ctx context.Context, condition Book, data Book) ([]Book, error)
}

type Adapter interface {
	SetAdapters()
}
