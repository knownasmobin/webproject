package usecase

import (
	"context"

	"git.ecobin.ir/ecomicro/template/app/user/domain"
	"git.ecobin.ir/ecomicro/tooty"
	"git.ecobin.ir/ecomicro/x"
)

type usecase struct {
	userRepo   domain.Repository
	sf         x.Sonyflake
	fooAdapter domain.FooAdapter
	bazAdapter domain.BazAdapter
}

var _ domain.Usecase = &usecase{}
var _ domain.Adapter = &usecase{}

func NewUserUsecase(userRepo domain.Repository, sf x.Sonyflake) *usecase {
	return &usecase{
		userRepo: userRepo,
		sf:       sf,
	}
}
func (uu *usecase) SetAdapters(fooAdapter domain.FooAdapter, bazAdapter domain.BazAdapter) {
	uu.fooAdapter = fooAdapter
	uu.bazAdapter = bazAdapter
}
func (uu *usecase) Create(
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
	err = uu.fooAdapter.Bar(ctx, *dbUser)
	if err != nil {
		x.LogError(err, ctx)
	}
	err = uu.bazAdapter.Create(ctx, *dbUser)
	if err != nil {
		x.LogError(err, ctx)
	}
	return dbUser, nil
}

func (uu *usecase) Update(ctx context.Context, user domain.User) (*domain.User, error) {
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
func (uu *usecase) GetUserById(ctx context.Context, id uint64) (*domain.User, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[U] get user by id", "usecase")
	defer tooty.CloseTheAPMSpan(span)
	user, err := uu.userRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
