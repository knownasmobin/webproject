package repository

import (
	"time"

	"git.ecobin.ir/ecomicro/template/app/category/domain"
	"gorm.io/gorm"
)

type Category struct {
	Id        uint64 `gorm:"primaryKey;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromDomainCategory(category domain.Category) Category {
	return Category{
		Id:        category.Id,
		CreatedAt: category.CreatedDate,
		UpdatedAt: category.UpdatedDate,
	}
}

func (u *Category) ToDomainCategory() domain.Category {
	return domain.Category{
		Id:          u.Id,
		CreatedDate: u.CreatedAt,
		UpdatedDate: u.UpdatedAt,
	}
}
