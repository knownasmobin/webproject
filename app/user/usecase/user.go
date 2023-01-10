package usecase

import (
	"context"

	"git.ecobin.ir/ecomicro/template/domain"
	"git.ecobin.ir/ecomicro/tooty"
	"github.com/sony/sonyflake"
)

type userUsecase struct {
	userRepo domain.UserRepository
	sf       *sonyflake.Sonyflake
}

var _ domain.UserUsecase = &userUsecase{}

func NewUserUsecase(userRepo domain.UserRepository, sf *sonyflake.Sonyflake) *userUsecase {
	return &userUsecase{
		userRepo: userRepo,
		sf:       sf,
	}
}
func (uu *userUsecase) Create(
	ctx context.Context,
	user domain.User,
) (*domain.User, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[U] create new user", "usecase")
	defer tooty.CloseTheAPMSpan(span)

	id, err := uu.sf.NextID()
	if err != nil {
		return nil, err
	}
	user.Id = id
	dbUser, err := uu.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return dbUser, nil
}

func (uu *userUsecase) Update(ctx context.Context, user domain.User) (*domain.User, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[U] update user", "usecase")
	defer tooty.CloseTheAPMSpan(span)
	userArray, err := uu.userRepo.Update(ctx, domain.User{
		Id: user.Id,
	}, user)
	if err != nil {
		return nil, err
	}
	if len(userArray) == 0 {
		return nil, domain.ErrNotFound
	}
	return &userArray[0], nil
}
func (uu *userUsecase) GetUserById(ctx context.Context, id uint64) (*domain.User, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[U] get user by id", "usecase")
	defer tooty.CloseTheAPMSpan(span)
	user, err := uu.userRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
