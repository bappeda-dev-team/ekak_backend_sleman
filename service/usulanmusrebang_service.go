package service

import (
	"context"
	"ekak_kab_sleman/model/web/usulan"
)

type UsulanMusrebangService interface {
	Create(ctx context.Context, request usulan.UsulanMusrebangCreateRequest) (usulan.UsulanMusrebangResponse, error)
	Update(ctx context.Context, request usulan.UsulanMusrebangUpdateRequest) (usulan.UsulanMusrebangResponse, error)
	FindById(ctx context.Context, idUsulan string) (usulan.UsulanMusrebangResponse, error)
	FindAll(ctx context.Context, kodeOpd *string, is_active *bool, rekinId *string, status *string) ([]usulan.UsulanMusrebangResponse, error)
	Delete(ctx context.Context, idUsulan string) error
	CreateRekin(ctx context.Context, request usulan.UsulanMusrebangCreateRekinRequest) ([]usulan.UsulanMusrebangResponse, error)
	DeleteUsulanTerpilih(ctx context.Context, idUsulan string) error
}
