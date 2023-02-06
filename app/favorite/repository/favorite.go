package repository

import (
	"context"

	"git.ecobin.ir/ecomicro/template/app/favorite/domain"

	"git.ecobin.ir/ecomicro/tooty"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type favoriteRepository struct {
	Conn *gorm.DB
}

var _ domain.Repository = &favoriteRepository{}

func NewFavoriteRepository(dbConnection *gorm.DB) *favoriteRepository {
	err := dbConnection.AutoMigrate(&Favorite{})
	if err != nil {
		panic(err)
	}
	return &favoriteRepository{dbConnection}
}

func (ur *favoriteRepository) Create(ctx context.Context, domainFavorite domain.Favorite) (*domain.Favorite, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[R] create favorite", "repository")
	defer tooty.CloseTheAPMSpan(span)

	favoriteDao := FromDomainFavorite(domainFavorite)
	result := ur.Conn.Debug().Create(&favoriteDao)
	if result.Error != nil {
		return nil, result.Error
	}
	favorite := favoriteDao.ToDomainFavorite()
	return &favorite, nil
}

func (ur *favoriteRepository) GetFavoriteById(ctx context.Context, id uint64) (*domain.Favorite, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[R] get favorite by id", "repository")
	defer tooty.CloseTheAPMSpan(span)
	var favoriteDao Favorite
	err := ur.Conn.WithContext(ctx).Debug().Where(Favorite{Id: id}).Find(&favoriteDao).Error
	if err != nil {
		return nil, err
	}
	favorite := favoriteDao.ToDomainFavorite()
	return &favorite, nil
}

func (ur *favoriteRepository) Update(ctx context.Context, condition domain.Favorite, domainFavorite domain.Favorite) ([]domain.Favorite, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[R] update favorite", "repository")
	defer tooty.CloseTheAPMSpan(span)
	var favoriteArray []Favorite
	err := ur.Conn.WithContext(ctx).Debug().Model(&favoriteArray).Clauses(clause.Returning{}).Where(FromDomainFavorite(condition)).Updates(FromDomainFavorite(domainFavorite)).Error
	if err != nil {
		return []domain.Favorite{}, err
	}
	domainFavorites := make([]domain.Favorite, len(favoriteArray))
	for idx, favorite := range favoriteArray {
		domainFavorites[idx] = favorite.ToDomainFavorite()
	}
	return domainFavorites, nil
}
