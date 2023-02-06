package http

import "git.ecobin.ir/ecomicro/template/app/favorite/domain"

type CreateFavoriteBody struct {
}
type UpdateFavoriteBody struct {
	Id uint64 `json:"id"`
}

func (c UpdateFavoriteBody) toDomain() domain.Favorite {
	return domain.Favorite{
		Id: c.Id,
	}
}

func (c CreateFavoriteBody) toDomain() domain.Favorite {
	return domain.Favorite{}
}

type FavoriteIdUri struct {
	Id uint64 `uri:"id"`
}
