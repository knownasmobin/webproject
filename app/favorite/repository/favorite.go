package repository

import (
	"context"

	"git.ecobin.ir/ecomicro/template/app/favorite/domain"

	"git.ecobin.ir/ecomicro/tooty"
	"gorm.io/gorm"
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

	favoriteDao := FromDomainFavorite(domainFavorite)
	result := ur.Conn.Debug().Create(&favoriteDao)
	if result.Error != nil {
		return nil, result.Error
	}
	favorite := favoriteDao.ToDomainFavorite()
	return &favorite, nil
}

func (ur *favoriteRepository) GetByCondition(ctx context.Context, condition domain.Favorite) ([]domain.Favorite, error) {
	var favoriteArray []Favorite
	err := ur.Conn.WithContext(ctx).Debug().Where(FromDomainFavorite(condition)).Find(&favoriteArray).Error
	if err != nil {
		return []domain.Favorite{}, err
	}
	domainFavorites := make([]domain.Favorite, len(favoriteArray))
	for idx, favorite := range favoriteArray {
		domainFavorites[idx] = favorite.ToDomainFavorite()
	}
	return domainFavorites, nil
}
func (ur *favoriteRepository) GetMostFavorites(ctx context.Context) ([]domain.Favorite, error) {
	var favoriteArray []Favorite
	err := ur.Conn.WithContext(ctx).Debug().Model(&Favorite{}).
		Select("book_id, count(*)").Group("book_id").Order("count desc").Scan(&favoriteArray).Error
	if err != nil {
		return []domain.Favorite{}, err
	}
	domainFavorites := make([]domain.Favorite, len(favoriteArray))
	for idx, favorite := range favoriteArray {
		domainFavorites[idx] = favorite.ToDomainFavorite()
	}
	return domainFavorites, nil
}
func (ur *favoriteRepository) Delete(ctx context.Context, id int) (*domain.Favorite, error) {
	var favoriteArray []Favorite
	result := ur.Conn.Debug().Delete(&favoriteArray, Favorite{Id: id})
	if result.Error != nil {
		return nil, result.Error
	}
	var favorite *domain.Favorite = nil
	if len(favoriteArray) > 0 {
		temp := favoriteArray[0].ToDomainFavorite()
		favorite = &temp
	}
	return favorite, nil
}
