package service

import (
	"context"
	"ekak_kab_sleman/model/web/user"
)

type RoleService interface {
	Create(ctx context.Context, request user.RoleCreateRequest) (user.RoleResponse, error)
	Update(ctx context.Context, request user.RoleUpdateRequest) (user.RoleResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (user.RoleResponse, error)
	FindAll(ctx context.Context) ([]user.RoleResponse, error)
}
