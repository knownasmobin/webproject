package repository

import (
	"time"

	"git.ecobin.ir/ecomicro/template/app/favorite/domain"
	"gorm.io/gorm"
)

type Favorite struct {
	Id        uint64 `gorm:"primaryKey;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromDomainFavorite(favorite domain.Favorite) Favorite {
	return Favorite{
		Id:        favorite.Id,
		CreatedAt: favorite.CreatedDate,
		UpdatedAt: favorite.UpdatedDate,
	}
}

func (u *Favorite) ToDomainFavorite() domain.Favorite {
	return domain.Favorite{
		Id:          u.Id,
		CreatedDate: u.CreatedAt,
		UpdatedDate: u.UpdatedAt,
	}
}
