package repository

import (
	"context"

	"git.ecobin.ir/ecomicro/template/app/purchase/domain"

	"git.ecobin.ir/ecomicro/tooty"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type purchaseRepository struct {
	Conn *gorm.DB
}

var _ domain.Repository = &purchaseRepository{}

func NewPurchaseRepository(dbConnection *gorm.DB) *purchaseRepository {
	err := dbConnection.AutoMigrate(&Purchase{})
	if err != nil {
		panic(err)
	}
	return &purchaseRepository{dbConnection}
}

func (ur *purchaseRepository) Create(ctx context.Context, domainPurchase domain.Purchase) (*domain.Purchase, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[R] create purchase", "repository")
	defer tooty.CloseTheAPMSpan(span)

	purchaseDao := FromDomainPurchase(domainPurchase)
	result := ur.Conn.Debug().Create(&purchaseDao)
	if result.Error != nil {
		return nil, result.Error
	}
	purchase := purchaseDao.ToDomainPurchase()
	return &purchase, nil
}

func (ur *purchaseRepository) GetPurchaseById(ctx context.Context, id uint64) (*domain.Purchase, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[R] get purchase by id", "repository")
	defer tooty.CloseTheAPMSpan(span)
	var purchaseDao Purchase
	err := ur.Conn.WithContext(ctx).Debug().Where(Purchase{Id: id}).Find(&purchaseDao).Error
	if err != nil {
		return nil, err
	}
	purchase := purchaseDao.ToDomainPurchase()
	return &purchase, nil
}

func (ur *purchaseRepository) Update(ctx context.Context, condition domain.Purchase, domainPurchase domain.Purchase) ([]domain.Purchase, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[R] update purchase", "repository")
	defer tooty.CloseTheAPMSpan(span)
	var purchaseArray []Purchase
	err := ur.Conn.WithContext(ctx).Debug().Model(&purchaseArray).Clauses(clause.Returning{}).Where(FromDomainPurchase(condition)).Updates(FromDomainPurchase(domainPurchase)).Error
	if err != nil {
		return []domain.Purchase{}, err
	}
	domainPurchases := make([]domain.Purchase, len(purchaseArray))
	for idx, purchase := range purchaseArray {
		domainPurchases[idx] = purchase.ToDomainPurchase()
	}
	return domainPurchases, nil
}
