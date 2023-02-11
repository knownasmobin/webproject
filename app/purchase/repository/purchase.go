package repository

import (
	"context"

	"git.ecobin.ir/ecomicro/template/app/purchase/domain"

	"gorm.io/gorm"
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
	purchaseDao := FromDomainPurchase(domainPurchase)
	result := ur.Conn.Debug().Create(&purchaseDao)
	if result.Error != nil {
		return nil, result.Error
	}
	purchase := purchaseDao.ToDomainPurchase()
	return &purchase, nil
}

func (ur *purchaseRepository) GetByCondition(ctx context.Context, domainPurchase domain.Purchase) ([]domain.Purchase, error) {
	var purchaseArray []Purchase
	err := ur.Conn.WithContext(ctx).Debug().Where(FromDomainPurchase(domainPurchase)).Find(&purchaseArray).Error
	if err != nil {
		return []domain.Purchase{}, err
	}
	domainPurchases := make([]domain.Purchase, len(purchaseArray))
	for idx, purchase := range purchaseArray {
		domainPurchases[idx] = purchase.ToDomainPurchase()
	}
	return domainPurchases, nil
}
