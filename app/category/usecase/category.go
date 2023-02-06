package usecase

import (
	"context"

	"git.ecobin.ir/ecomicro/template/app/category/domain"
	"git.ecobin.ir/ecomicro/tooty"
)

type categoryUsecase struct {
	categoryRepo domain.Repository
}

var _ domain.Usecase = &categoryUsecase{}
var _ domain.Adapter = &categoryUsecase{}

func NewCategoryUsecase(categoryRepo domain.Repository) *categoryUsecase {
	return &categoryUsecase{
		categoryRepo: categoryRepo,
	}
}
func (uu *categoryUsecase) SetAdapters() {

}
func (uu *categoryUsecase) Create(
	ctx context.Context,
	category domain.Category,
) (*domain.Category, error) {

	span := tooty.OpenAnAPMSpan(ctx, "[U] create new category", "usecase")
	defer tooty.CloseTheAPMSpan(span)

	dbCategory, err := uu.categoryRepo.Create(ctx, category)
	if err != nil {
		return nil, err
	}

	return dbCategory, nil
}

func (uu *categoryUsecase) Update(ctx context.Context, category domain.Category) (*domain.Category, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[U] update category", "usecase")
	defer tooty.CloseTheAPMSpan(span)
	categoryArray, err := uu.categoryRepo.Update(ctx, domain.Category{
		Id: category.Id,
	}, category)
	if err != nil {
		return nil, err
	}
	if len(categoryArray) == 0 {
		return nil, domain.ErrNotFound
	}
	return &categoryArray[0], nil
}
func (uu *categoryUsecase) GetCategoryById(ctx context.Context, id uint64) (*domain.Category, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[U] get category by id", "usecase")
	defer tooty.CloseTheAPMSpan(span)
	category, err := uu.categoryRepo.GetCategoryById(ctx, id)
	if err != nil {
		return nil, err
	}
	return category, nil
}
