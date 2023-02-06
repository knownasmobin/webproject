package usecase

import (
	"context"

	"git.ecobin.ir/ecomicro/template/app/purchase/domain"
	"git.ecobin.ir/ecomicro/tooty"
)

type purchaseUsecase struct {
	purchaseRepo domain.Repository
}

var _ domain.Usecase = &purchaseUsecase{}
var _ domain.Adapter = &purchaseUsecase{}

func NewPurchaseUsecase(purchaseRepo domain.Repository) *purchaseUsecase {
	return &purchaseUsecase{
		purchaseRepo: purchaseRepo,
	}
}
func (uu *purchaseUsecase) SetAdapters() {

}
func (uu *purchaseUsecase) Create(
	ctx context.Context,
	purchase domain.Purchase,
) (*domain.Purchase, error) {

	span := tooty.OpenAnAPMSpan(ctx, "[U] create new purchase", "usecase")
	defer tooty.CloseTheAPMSpan(span)

	dbPurchase, err := uu.purchaseRepo.Create(ctx, purchase)
	if err != nil {
		return nil, err
	}

	return dbPurchase, nil
}

func (uu *purchaseUsecase) Update(ctx context.Context, purchase domain.Purchase) (*domain.Purchase, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[U] update purchase", "usecase")
	defer tooty.CloseTheAPMSpan(span)
	purchaseArray, err := uu.purchaseRepo.Update(ctx, domain.Purchase{
		Id: purchase.Id,
	}, purchase)
	if err != nil {
		return nil, err
	}
	if len(purchaseArray) == 0 {
		return nil, domain.ErrNotFound
	}
	return &purchaseArray[0], nil
}
func (uu *purchaseUsecase) GetPurchaseById(ctx context.Context, id uint64) (*domain.Purchase, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[U] get purchase by id", "usecase")
	defer tooty.CloseTheAPMSpan(span)
	purchase, err := uu.purchaseRepo.GetPurchaseById(ctx, id)
	if err != nil {
		return nil, err
	}
	return purchase, nil
}
