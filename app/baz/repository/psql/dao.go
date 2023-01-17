package repository

import (
	"time"

	"git.ecobin.ir/ecomicro/template/app/baz/domain"
	"gorm.io/gorm"
)

type Baz struct {
	UserId    uint64 `gorm:"primaryKey;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromDomainBaz(baz domain.Baz) Baz {
	return Baz{
		UserId:    baz.UserId,
		CreatedAt: baz.CreatedDate,
		UpdatedAt: baz.UpdatedDate,
	}
}

func (b *Baz) ToDomainBaz() domain.Baz {
	return domain.Baz{
		UserId:      b.UserId,
		CreatedDate: b.CreatedAt,
		UpdatedDate: b.UpdatedAt,
	}
}
