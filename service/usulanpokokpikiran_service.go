package service

import (
	"context"
	"ekak_kab_sleman/model/web/usulan"
)

type UsulanPokokPikiranService interface {
	Create(ctx context.Context, request usulan.UsulanPokokPikiranCreateRequest) (usulan.UsulanPokokPikiranResponse, error)
	Update(ctx context.Context, request usulan.UsulanPokokPikiranUpdateRequest) (usulan.UsulanPokokPikiranResponse, error)
	FindById(ctx context.Context, idUsulan string) (usulan.UsulanPokokPikiranResponse, error)
	FindAll(ctx context.Context, kodeOpd *string, isActive *bool, rekinId *string, status *string) ([]usulan.UsulanPokokPikiranResponse, error)
	Delete(ctx context.Context, idUsulan string) error
	CreateRekin(ctx context.Context, request usulan.UsulanPokokPikiranCreateRekinRequest) ([]usulan.UsulanPokokPikiranResponse, error)
	DeleteUsulanTerpilih(ctx context.Context, idUsulan string) error
}
