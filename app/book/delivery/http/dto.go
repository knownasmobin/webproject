package http

import (
	"git.ecobin.ir/ecomicro/template/app/book/domain"
)

type CreateBookBody struct {
	Title       string
	Price       float32
	Author      string
	Description string
	Image       string
	Categories  []int
}
type UpdateBookBody struct {
	Id          int
	Title       string
	Price       float32
	Author      string
	Description string
	Image       string
	Categories  []int
}

func (c UpdateBookBody) toDomain() domain.Book {
	return domain.Book{
		Id:          c.Id,
		Title:       c.Title,
		Price:       c.Price,
		Author:      c.Author,
		Categories:  c.Categories,
		Description: c.Description,
		Image:       c.Image,
	}
}

func (c CreateBookBody) toDomain() domain.Book {
	return domain.Book{
		Title:       c.Title,
		Price:       c.Price,
		Author:      c.Author,
		Categories:  c.Categories,
		Description: c.Description,
		Image:       c.Image,
	}
}

type BookIdUri struct {
	Id int `uri:"id"`
}
type IdFromQuery struct {
	Id int `form:"CategoryId"`
}
