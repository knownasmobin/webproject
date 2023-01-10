package domain

import (
	"context"
)

type Permission struct {
	UniqueName string
}

type PermissionUsecase interface {
	Add(ctx context.Context, permission Permission) (*Permission, error)
	GetPermissions(ctx context.Context) ([]Permission, error)
	Remove(ctx context.Context, uniqueName string) (*Permission, error)
}

type PermissionRepository interface {
	Create(ctx context.Context, permission Permission) (*Permission, error)
	Update(ctx context.Context, condition Permission, data Permission) (*Permission, error)
	GetPermissions(ctx context.Context) ([]Permission, error)
}
