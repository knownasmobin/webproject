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

	dbCategory, err := uu.categoryRepo.Create(ctx, category)
	if err != nil {
		return nil, err
	}

	return dbCategory, nil
}

func (uu *categoryUsecase) Update(ctx context.Context, category domain.Category) (*domain.Category, error) {

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
func (uu *categoryUsecase) GetCategoryById(ctx context.Context, id int) (*domain.Category, error) {

	category, err := uu.categoryRepo.GetCategoryById(ctx, id)
	if err != nil {
		return nil, err
	}
	return category, nil
}
func (uu *categoryUsecase) GetAll(ctx context.Context) ([]domain.Category, error) {

	categories, err := uu.categoryRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
