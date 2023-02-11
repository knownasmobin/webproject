package usecase

import (
	"context"

	"git.ecobin.ir/ecomicro/template/app/purchase/domain"
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

	dbPurchase, err := uu.purchaseRepo.Create(ctx, purchase)
	if err != nil {
		return nil, err
	}

	return dbPurchase, nil
}

func (uu *purchaseUsecase) GetByCondition(ctx context.Context, purchase domain.Purchase) ([]domain.Purchase, error) {
	purchaseArray, err := uu.purchaseRepo.GetByCondition(ctx, purchase)
	if err != nil {
		return nil, err
	}
	return purchaseArray, nil
}
