package repository

import (
	"time"

	"git.ecobin.ir/ecomicro/template/app/purchase/domain"
	"gorm.io/gorm"
)

type Purchase struct {
	Id        uint64 `gorm:"primaryKey;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromDomainPurchase(purchase domain.Purchase) Purchase {
	return Purchase{
		Id:        purchase.Id,
		CreatedAt: purchase.CreatedDate,
		UpdatedAt: purchase.UpdatedDate,
	}
}

func (u *Purchase) ToDomainPurchase() domain.Purchase {
	return domain.Purchase{
		Id:          u.Id,
		CreatedDate: u.CreatedAt,
		UpdatedDate: u.UpdatedAt,
	}
}
