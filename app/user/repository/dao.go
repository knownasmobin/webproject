package repository

import (
	"time"

	"git.ecobin.ir/ecomicro/template/app/user/domain"
	"gorm.io/gorm"
)

type User struct {
	Id        int `gorm:"primaryKey;unique"`
	Name      string
	Username  string
	Password  string
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromDomainUser(user domain.User) User {
	return User{
		Id:        user.Id,
		Name:      user.Name,
		Username:  user.Username,
		Password:  user.Password,
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedDate,
		UpdatedAt: user.UpdatedDate,
	}
}

func (u *User) ToDomainUser() domain.User {
	return domain.User{
		Id:          u.Id,
		Name:        u.Name,
		Username:    u.Username,
		Password:    u.Password,
		IsAdmin:     u.IsAdmin,
		CreatedDate: u.CreatedAt,
		UpdatedDate: u.UpdatedAt,
	}
}
