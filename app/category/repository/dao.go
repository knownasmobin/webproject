package repository

import (
	"time"

	"git.ecobin.ir/ecomicro/template/app/category/domain"
	"gorm.io/gorm"
)

type Category struct {
	Id        int `gorm:"primaryKey;unique"`
	EnName    string
	FaName    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromDomainCategory(category domain.Category) Category {
	return Category{
		Id:        category.Id,
		FaName:    category.FaName,
		EnName:    category.EnName,
		CreatedAt: category.CreatedDate,
		UpdatedAt: category.UpdatedDate,
	}
}

func (u *Category) ToDomainCategory() domain.Category {
	return domain.Category{
		Id:          u.Id,
		FaName:      u.FaName,
		EnName:      u.EnName,
		CreatedDate: u.CreatedAt,
		UpdatedDate: u.UpdatedAt,
	}
}
