package http

import "git.ecobin.ir/ecomicro/template/app/category/domain"

type CreateCategoryBody struct {
	EnName string
	FaName string
}
type UpdateCategoryBody struct {
	Id     int `json:"id"`
	EnName string
	FaName string
}

func (c UpdateCategoryBody) toDomain() domain.Category {
	return domain.Category{
		Id:     c.Id,
		FaName: c.FaName,
		EnName: c.EnName,
	}
}

func (c CreateCategoryBody) toDomain() domain.Category {
	return domain.Category{
		FaName: c.FaName,
		EnName: c.EnName,
	}
}

type CategoryIdUri struct {
	Id int `uri:"id"`
}
