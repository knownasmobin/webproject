package http

import "git.ecobin.ir/ecomicro/template/app/purchase/domain"

type CreatePurchaseBody struct {
}
type UpdatePurchaseBody struct {
	Id uint64 `json:"id"`
}

func (c UpdatePurchaseBody) toDomain() domain.Purchase {
	return domain.Purchase{
		Id: c.Id,
	}
}

func (c CreatePurchaseBody) toDomain() domain.Purchase {
	return domain.Purchase{}
}

type PurchaseIdUri struct {
	Id uint64 `uri:"id"`
}
