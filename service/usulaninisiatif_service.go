package service

import (
	"context"
	"ekak_kab_sleman/model/web/usulan"
)

type UsulanInisiatifService interface {
	Create(ctx context.Context, request usulan.UsulanInisiatifCreateRequest) (usulan.UsulanInisiatifResponse, error)
	Update(ctx context.Context, request usulan.UsulanInisiatifUpdateRequest) (usulan.UsulanInisiatifResponse, error)
	FindById(ctx context.Context, idUsulan string) (usulan.UsulanInisiatifResponse, error)
	FindAll(ctx context.Context, pegawaiId *string, isActive *bool, rekinId *string) ([]usulan.UsulanInisiatifResponse, error)
	Delete(ctx context.Context, idUsulan string) error
}
