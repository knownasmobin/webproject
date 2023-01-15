package adapter

import (
	"context"

	"git.ecobin.ir/ecomicro/template/domain"
)

type FooAdapter interface {
	Bar(ctx context.Context, user domain.User) error
}
