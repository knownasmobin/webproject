package usecase

import (
	"context"

	"git.ecobin.ir/ecomicro/template/app/favorite/domain"
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

	dbFavorite, err := uu.favoriteRepo.Create(ctx, favorite)
	if err != nil {
		return nil, err
	}

	return dbFavorite, nil
}

func (uu *favoriteUsecase) Delete(ctx context.Context, id int) (*domain.Favorite, error) {
	favoriteArray, err := uu.favoriteRepo.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	return favoriteArray, nil
}
func (uu *favoriteUsecase) GetMostFavorites(ctx context.Context) ([]domain.Favorite, error) {
	favorite, err := uu.favoriteRepo.GetMostFavorites(ctx)
	if err != nil {
		return nil, err
	}
	return favorite, nil
}

func (uu *favoriteUsecase) GetByCondition(ctx context.Context, favorite domain.Favorite) ([]domain.Favorite, error) {
	favorites, err := uu.favoriteRepo.GetByCondition(ctx, favorite)
	if err != nil {
		return nil, err
	}
	return favorites, nil
}
