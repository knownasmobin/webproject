package repository

import (
	"time"

	"git.ecobin.ir/ecomicro/template/app/book/domain"
	"gorm.io/gorm"
)

type Book struct {
	Id         int `gorm:"primaryKey;unique"`
	Title      string
	Price      float32
	Author     string
	CategoryId int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func FromDomainBook(book domain.Book) Book {
	return Book{
		Id:         book.Id,
		Title:      book.Title,
		Price:      book.Price,
		Author:     book.Author,
		CategoryId: book.CategoryId,
		CreatedAt:  book.CreatedDate,
		UpdatedAt:  book.UpdatedDate,
	}
}

func (u *Book) ToDomainBook() domain.Book {
	return domain.Book{
		Id:          u.Id,
		Title:       u.Title,
		Price:       u.Price,
		Author:      u.Author,
		CategoryId:  u.CategoryId,
		CreatedDate: u.CreatedAt,
		UpdatedDate: u.UpdatedAt,
	}
}
