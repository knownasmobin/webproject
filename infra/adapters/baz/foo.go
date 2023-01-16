package baz

import (
	"context"

	userAdapters "git.ecobin.ir/ecomicro/template/app/user/adapter"
	"git.ecobin.ir/ecomicro/template/domain"
	"git.ecobin.ir/ecomicro/tooty"
)

type bazAdapter struct {
	usecase domain.BazUsecase
}

var _ userAdapters.BazAdapter = &bazAdapter{}

func NewBazUsecaseAdapter(usecase domain.BazUsecase) *bazAdapter {
	return &bazAdapter{usecase}
}
func (e *bazAdapter) Create(ctx context.Context, user domain.User) error {
	span := tooty.OpenAnAPMSpan(ctx, "[A] create", "adapter")
	defer tooty.CloseTheAPMSpan(span)

	_, err := e.usecase.Create(ctx, domain.Baz{UserId: user.Id})

	return err
}
