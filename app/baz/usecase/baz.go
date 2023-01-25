package usecase

import (
	"context"

	"git.ecobin.ir/ecomicro/template/app/baz/domain"
	"git.ecobin.ir/ecomicro/tooty"
)

type usecase struct {
	bazRepo domain.Repository
}

var _ domain.Usecase = &usecase{}

func NewBazUsecase(bazRepo domain.Repository) *usecase {
	return &usecase{
		bazRepo: bazRepo,
	}
}
func (uu *usecase) SetAdapters() {}

func (uu *usecase) Create(
	ctx context.Context,
	baz domain.Baz,
) (*domain.Baz, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[U] create new baz", "usecase")
	defer tooty.CloseTheAPMSpan(span)

	dbBaz, err := uu.bazRepo.Create(ctx, baz)
	if err != nil {
		return nil, err
	}
	return dbBaz, nil
}
