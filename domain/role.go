package domain

import (
	"context"
	"time"
)

type Role struct {
	Id          uint64
	UniqueName  string
	DisplayName string
	Permissions *[]string
	CreatedDate time.Time
	UpdatedDate time.Time
	CreatedBy   uint64
	UpdatedBy   uint64
}

type RoleUsecase interface {
	GetRoleById(ctx context.Context, id uint64) (*Role, error)
	GetRoleByUniqueName(ctx context.Context, uniqueName string) (*Role, error)
	GetRoles(ctx context.Context) ([]Role, error)
	Create(ctx context.Context, role Role) (*Role, error)
	Update(ctx context.Context, condition Role, data Role) (*Role, error)
}

type RoleRepository interface {
	GetRoleById(ctx context.Context, id uint64) (*Role, error)
	GetRoleByUniqueName(ctx context.Context, uniqueName string) (*Role, error)
	GetRoles(ctx context.Context) ([]Role, error)
	Create(ctx context.Context, role Role) (*Role, error)
	Update(ctx context.Context, condition Role, data Role) (*Role, error)
}
