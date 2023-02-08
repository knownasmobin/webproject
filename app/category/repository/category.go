package repository

import (
	"context"

	"git.ecobin.ir/ecomicro/template/app/category/domain"

	"git.ecobin.ir/ecomicro/tooty"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type categoryRepository struct {
	Conn *gorm.DB
}

var _ domain.Repository = &categoryRepository{}

func NewCategoryRepository(dbConnection *gorm.DB) *categoryRepository {
	err := dbConnection.AutoMigrate(&Category{})
	if err != nil {
		panic(err)
	}
	return &categoryRepository{dbConnection}
}

func (ur *categoryRepository) Create(ctx context.Context, domainCategory domain.Category) (*domain.Category, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[R] create category", "repository")
	defer tooty.CloseTheAPMSpan(span)

	categoryDao := FromDomainCategory(domainCategory)
	result := ur.Conn.Debug().Create(&categoryDao)
	if result.Error != nil {
		return nil, result.Error
	}
	category := categoryDao.ToDomainCategory()
	return &category, nil
}

func (ur *categoryRepository) GetCategoryById(ctx context.Context, id int) (*domain.Category, error) {

	var categoryDao Category
	err := ur.Conn.WithContext(ctx).Debug().Where(Category{Id: id}).Find(&categoryDao).Error
	if err != nil {
		return nil, err
	}
	category := categoryDao.ToDomainCategory()
	return &category, nil
}
func (ur *categoryRepository) GetAll(ctx context.Context) ([]domain.Category, error) {

	var categoryArray []Category
	err := ur.Conn.WithContext(ctx).Debug().Find(&categoryArray).Error
	if err != nil {
		return []domain.Category{}, err
	}
	domainCategorys := make([]domain.Category, len(categoryArray))
	for idx, category := range categoryArray {
		domainCategorys[idx] = category.ToDomainCategory()
	}
	return domainCategorys, nil
}
func (ur *categoryRepository) Update(ctx context.Context, condition domain.Category, domainCategory domain.Category) ([]domain.Category, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[R] update category", "repository")
	defer tooty.CloseTheAPMSpan(span)
	var categoryArray []Category
	err := ur.Conn.WithContext(ctx).Debug().Model(&categoryArray).Clauses(clause.Returning{}).Where(FromDomainCategory(condition)).Updates(FromDomainCategory(domainCategory)).Error
	if err != nil {
		return []domain.Category{}, err
	}
	domainCategorys := make([]domain.Category, len(categoryArray))
	for idx, category := range categoryArray {
		domainCategorys[idx] = category.ToDomainCategory()
	}
	return domainCategorys, nil
}
