package http

import "git.ecobin.ir/ecomicro/template/app/category/domain"

type CreateCategoryBody struct {
}
type UpdateCategoryBody struct {
	Id uint64 `json:"id"`
}

func (c UpdateCategoryBody) toDomain() domain.Category {
	return domain.Category{
		Id: c.Id,
	}
}

func (c CreateCategoryBody) toDomain() domain.Category {
	return domain.Category{}
}

type CategoryIdUri struct {
	Id uint64 `uri:"id"`
}
