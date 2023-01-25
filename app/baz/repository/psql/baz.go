package repository

import (
	"context"

	"git.ecobin.ir/ecomicro/template/app/baz/domain"

	"git.ecobin.ir/ecomicro/tooty"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	Conn *gorm.DB
}

var _ domain.Repository = &repository{}

func NewRepository(dbConnection *gorm.DB) *repository {
	err := dbConnection.AutoMigrate(&Baz{})
	if err != nil {
		panic(err)
	}
	return &repository{dbConnection}
}
func (ur *repository) Create(ctx context.Context, domainBaz domain.Baz) (*domain.Baz, error) {
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
func (ur *repository) Get(ctx context.Context, userId uint64) (*domain.Baz, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[R] get baz", "repository")
	defer tooty.CloseTheAPMSpan(span)

	var bazDao Baz
	err := ur.Conn.WithContext(ctx).Debug().Where(Baz{UserId: userId}).First(&bazDao).Error
	if err != nil {
		return nil, err
	}
	baz := bazDao.ToDomainBaz()
	return &baz, nil
}
func (ur *repository) Update(ctx context.Context, condition, domainBaz domain.Baz) ([]domain.Baz, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[R] update baz", "repository")
	defer tooty.CloseTheAPMSpan(span)

	var bazArray []Baz
	err := ur.Conn.WithContext(ctx).Debug().
		Model(&bazArray).Clauses(clause.Returning{}).
		Where(FromDomainBaz(condition)).Updates(FromDomainBaz(domainBaz)).Error
	if err != nil {
		return []domain.Baz{}, err
	}
	domainBazs := make([]domain.Baz, len(bazArray))
	for idx, baz := range bazArray {
		domainBazs[idx] = baz.ToDomainBaz()
	}
	return domainBazs, nil
}
