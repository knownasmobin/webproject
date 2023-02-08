package http

import "git.ecobin.ir/ecomicro/template/app/favorite/domain"

type CreateFavoriteBody struct {
	BookId int
	UserId int
}
type UpdateFavoriteBody struct {
	Id     int
	BookId int
	UserId int
}

func (c UpdateFavoriteBody) toDomain() domain.Favorite {
	return domain.Favorite{
		Id:     c.Id,
		BookId: c.BookId,
		UserId: c.UserId,
	}
}

func (c CreateFavoriteBody) toDomain() domain.Favorite {
	return domain.Favorite{
		BookId: c.BookId,
		UserId: c.UserId,
	}
}

type FavoriteIdUri struct {
	Id int `uri:"id"`
}
