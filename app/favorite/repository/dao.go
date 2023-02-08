package repository

import (
	"time"

	"git.ecobin.ir/ecomicro/template/app/favorite/domain"
	"gorm.io/gorm"
)

type Favorite struct {
	Id        int `gorm:"primaryKey;unique"`
	BookId    int
	UserId    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromDomainFavorite(favorite domain.Favorite) Favorite {
	return Favorite{
		Id:        favorite.Id,
		UserId:    favorite.UserId,
		BookId:    favorite.BookId,
		CreatedAt: favorite.CreatedDate,
		UpdatedAt: favorite.UpdatedDate,
	}
}

func (u *Favorite) ToDomainFavorite() domain.Favorite {
	return domain.Favorite{
		Id:          u.Id,
		UserId:      u.UserId,
		BookId:      u.BookId,
		CreatedDate: u.CreatedAt,
		UpdatedDate: u.UpdatedAt,
	}
}
