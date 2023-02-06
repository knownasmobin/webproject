package http

import "git.ecobin.ir/ecomicro/template/app/book/domain"

type CreateBookBody struct {
	Title      string
	Price      float32
	Author     string
	CategoryId int
}
type UpdateBookBody struct {
	Id         int
	Title      string
	Price      float32
	Author     string
	CategoryId int
}

func (c UpdateBookBody) toDomain() domain.Book {
	return domain.Book{
		Id:         c.Id,
		Title:      c.Title,
		Price:      c.Price,
		Author:     c.Author,
		CategoryId: c.CategoryId,
	}
}

func (c CreateBookBody) toDomain() domain.Book {
	return domain.Book{
		Title:      c.Title,
		Price:      c.Price,
		Author:     c.Author,
		CategoryId: c.CategoryId,
	}
}

type BookIdUri struct {
	Id int `uri:"id"`
}
type IdFromQuery struct {
	Id int `form:"categoryId"`
}
