package repository

import (
	"context"
	"log"

	"git.ecobin.ir/ecomicro/template/app/user/domain"

	"git.ecobin.ir/ecomicro/tooty"
	"golang.org/x/crypto/bcrypt"
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
	userDao := FromDomainUser(domainUser)
	if userDao.Password != "" {
		// hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDao.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		log.Println("password", userDao.Password, string(hashedPassword))
		stringHashedPassword := string(hashedPassword)
		userDao.Password = stringHashedPassword
	}
	result := ur.Conn.Debug().Create(&userDao)
	if result.Error != nil {
		return nil, result.Error
	}
	user := userDao.ToDomainUser()
	return &user, nil
}
func (ur *userRepository) GetUserById(ctx context.Context, id int) (*domain.User, error) {
	var userDao User
	err := ur.Conn.WithContext(ctx).Debug().Where(User{Id: id}).First(&userDao).Error
	if err != nil {
		return nil, err
	}
	user := userDao.ToDomainUser()
	return &user, nil
}
func (ur *userRepository) GetUserByPassword(ctx context.Context, username, password string) (*domain.User, error) {

	var userDao User
	err := ur.Conn.WithContext(ctx).Debug().Where(User{Username: username}).First(&userDao).Error
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userDao.Password), []byte(password)); err != nil {
		return nil, err
	}
	user := userDao.ToDomainUser()
	return &user, nil
}
func (ur *userRepository) Update(ctx context.Context, condition domain.User, domainUser domain.User) ([]domain.User, error) {
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

func (ur *userRepository) GetByCondition(ctx context.Context, condition domain.User) ([]domain.User, error) {
	var userArray []User
	err := ur.Conn.WithContext(ctx).Debug().Where(FromDomainUser(condition)).Find(&userArray).Error
	if err != nil {
		return []domain.User{}, err
	}
	domainUsers := make([]domain.User, len(userArray))
	for idx, user := range userArray {
		domainUsers[idx] = user.ToDomainUser()
	}
	return domainUsers, nil
}
