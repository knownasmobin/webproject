package adapter

import (
	"context"

	"git.ecobin.ir/ecomicro/template/domain"
)

type BazAdapter interface {
	Create(ctx context.Context, user domain.User) error
}
