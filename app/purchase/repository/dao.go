package repository

import (
	"time"

	"git.ecobin.ir/ecomicro/template/app/purchase/domain"
	"gorm.io/gorm"
)

type Purchase struct {
	Id        int `gorm:"primaryKey;unique"`
	UserId    int
	BookId    int
	Price     float32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromDomainPurchase(purchase domain.Purchase) Purchase {
	return Purchase{
		Id:        purchase.Id,
		Price:     purchase.Price,
		UserId:    purchase.UserId,
		BookId:    purchase.BookId,
		CreatedAt: purchase.CreatedDate,
		UpdatedAt: purchase.UpdatedDate,
	}
}

func (u *Purchase) ToDomainPurchase() domain.Purchase {
	return domain.Purchase{
		Id:          u.Id,
		Price:       u.Price,
		UserId:      u.UserId,
		BookId:      u.BookId,
		CreatedDate: u.CreatedAt,
		UpdatedDate: u.UpdatedAt,
	}
}
