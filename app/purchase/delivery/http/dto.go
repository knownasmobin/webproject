package http

import (
	"git.ecobin.ir/ecomicro/template/app/purchase/domain"
)

type CreatePurchaseBody struct {
	UserId int
	BookId int
	Price  float32
}
type UpdatePurchaseBody struct {
	Id     int
	UserId int
	BookId int
	Price  float32
}

func (c UpdatePurchaseBody) toDomain() domain.Purchase {
	return domain.Purchase{
		Id:     c.Id,
		UserId: c.UserId,
		BookId: c.BookId,
		Price:  c.Price,
	}
}

func (c CreatePurchaseBody) toDomain() domain.Purchase {
	return domain.Purchase{
		UserId: c.UserId,
		BookId: c.BookId,
		Price:  c.Price,
	}
}

type PurchaseIdUri struct {
	Id int `uri:"id"`
}
