package repository

import (
	"context"

	"git.ecobin.ir/ecomicro/template/app/user/domain"

	"git.ecobin.ir/ecomicro/tooty"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	Conn *gorm.DB
}

var _ domain.Repository = &userRepository{}

func NewRepository(dbConnection *gorm.DB) *userRepository {
	err := dbConnection.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
	return &userRepository{dbConnection}
}
func (ur *userRepository) Create(ctx context.Context, domainUser domain.User) (*domain.User, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[R] create user", "repository")
	defer tooty.CloseTheAPMSpan(span)

	userDao := FromDomainUser(domainUser)
	result := ur.Conn.Debug().Create(&userDao)
	if result.Error != nil {
		return nil, result.Error
	}
	user := userDao.ToDomainUser()
	return &user, nil
}
func (ur *userRepository) GetUserById(ctx context.Context, id uint64) (*domain.User, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[R] get user by id", "repository")
	defer tooty.CloseTheAPMSpan(span)
	var userDao User
	err := ur.Conn.WithContext(ctx).Debug().Where(User{Id: id}).First(&userDao).Error
	if err != nil {
		return nil, err
	}
	user := userDao.ToDomainUser()
	return &user, nil
}
func (ur *userRepository) Update(ctx context.Context, condition domain.User, domainUser domain.User) ([]domain.User, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[R] update user", "repository")
	defer tooty.CloseTheAPMSpan(span)
	var userArray []User
	err := ur.Conn.WithContext(ctx).Debug().Model(&userArray).Clauses(clause.Returning{}).Where(FromDomainUser(condition)).Updates(FromDomainUser(domainUser)).Error
	if err != nil {
		return []domain.User{}, err
	}
	domainUsers := make([]domain.User, len(userArray))
	for idx, user := range userArray {
		domainUsers[idx] = user.ToDomainUser()
	}
	return domainUsers, nil
}
