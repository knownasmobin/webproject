package usecase

import (
	"context"

	"git.ecobin.ir/ecomicro/template/app/favorite/domain"
	"git.ecobin.ir/ecomicro/tooty"
)

type favoriteUsecase struct {
	favoriteRepo domain.Repository
}

var _ domain.Usecase = &favoriteUsecase{}
var _ domain.Adapter = &favoriteUsecase{}

func NewFavoriteUsecase(favoriteRepo domain.Repository) *favoriteUsecase {
	return &favoriteUsecase{
		favoriteRepo: favoriteRepo,
	}
}
func (uu *favoriteUsecase) SetAdapters() {

}
func (uu *favoriteUsecase) Create(
	ctx context.Context,
	favorite domain.Favorite,
) (*domain.Favorite, error) {

	span := tooty.OpenAnAPMSpan(ctx, "[U] create new favorite", "usecase")
	defer tooty.CloseTheAPMSpan(span)

	dbFavorite, err := uu.favoriteRepo.Create(ctx, favorite)
	if err != nil {
		return nil, err
	}

	return dbFavorite, nil
}

func (uu *favoriteUsecase) Update(ctx context.Context, favorite domain.Favorite) (*domain.Favorite, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[U] update favorite", "usecase")
	defer tooty.CloseTheAPMSpan(span)
	favoriteArray, err := uu.favoriteRepo.Update(ctx, domain.Favorite{
		Id: favorite.Id,
	}, favorite)
	if err != nil {
		return nil, err
	}
	if len(favoriteArray) == 0 {
		return nil, domain.ErrNotFound
	}
	return &favoriteArray[0], nil
}
func (uu *favoriteUsecase) GetFavoriteById(ctx context.Context, id uint64) (*domain.Favorite, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[U] get favorite by id", "usecase")
	defer tooty.CloseTheAPMSpan(span)
	favorite, err := uu.favoriteRepo.GetFavoriteById(ctx, id)
	if err != nil {
		return nil, err
	}
	return favorite, nil
}
