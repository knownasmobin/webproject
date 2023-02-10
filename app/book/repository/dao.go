package repository

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"git.ecobin.ir/ecomicro/template/app/book/domain"
	"gorm.io/gorm"
)

type Book struct {
	Id          int `gorm:"primaryKey;unique"`
	Title       string
	Price       float32
	Author      string
	Description string
	Image       string
	Categories  JSON `gorm:"type:jsonb"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type JSON json.RawMessage

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

func FromDomainBook(book domain.Book) Book {
	var categories JSON
	categories, _ = json.Marshal(book.Categories)
	return Book{
		Id:          book.Id,
		Title:       book.Title,
		Price:       book.Price,
		Author:      book.Author,
		Categories:  categories,
		Image:       book.Image,
		Description: book.Description,
		CreatedAt:   book.CreatedDate,
		UpdatedAt:   book.UpdatedDate,
	}
}

func (u *Book) ToDomainBook() domain.Book {
	var categories []int
	_ = json.Unmarshal(u.Categories, &categories)
	return domain.Book{
		Id:          u.Id,
		Title:       u.Title,
		Price:       u.Price,
		Author:      u.Author,
		Categories:  categories,
		Image:       u.Image,
		Description: u.Description,
		CreatedDate: u.CreatedAt,
		UpdatedDate: u.UpdatedAt,
	}
}
