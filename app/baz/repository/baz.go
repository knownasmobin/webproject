package repository

import (
	"context"

	"git.ecobin.ir/ecomicro/template/domain"

	"git.ecobin.ir/ecomicro/tooty"
	"gorm.io/gorm"
)

type bazRepository struct {
	Conn *gorm.DB
}

var _ domain.BazRepository = &bazRepository{}

func NewBazRepository(dbConnection *gorm.DB) *bazRepository {
	err := dbConnection.AutoMigrate(&Baz{})
	if err != nil {
		panic(err)
	}
	return &bazRepository{dbConnection}
}
func (ur *bazRepository) Create(ctx context.Context, domainBaz domain.Baz) (*domain.Baz, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[R] create baz", "repository")
	defer tooty.CloseTheAPMSpan(span)

	bazDao := FromDomainBaz(domainBaz)
	result := ur.Conn.Debug().Create(&bazDao)
	if result.Error != nil {
		return nil, result.Error
	}
	baz := bazDao.ToDomainBaz()
	return &baz, nil
}
