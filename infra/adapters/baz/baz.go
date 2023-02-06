package baz

import (
	"context"

	bazDomain "git.ecobin.ir/ecomicro/template/app/baz/domain"
	userDomain "git.ecobin.ir/ecomicro/template/app/user/domain"
	"git.ecobin.ir/ecomicro/tooty"
)

type bazAdapter struct {
	usecase bazDomain.Usecase
}

var _ userDomain.BazAdapter = &bazAdapter{}

func NewBazUsecaseAdapter(usecase bazDomain.Usecase) *bazAdapter {
	return &bazAdapter{usecase}
}
func (e *bazAdapter) Create(ctx context.Context, user userDomain.User) error {
	span := tooty.OpenAnAPMSpan(ctx, "[A] create", "adapter")
	defer tooty.CloseTheAPMSpan(span)

	_, err := e.usecase.Create(ctx, bazDomain.Baz{UserId: user.Id})

	return err
}
