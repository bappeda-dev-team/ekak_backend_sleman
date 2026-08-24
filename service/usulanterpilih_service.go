package service

import (
	"context"
	"ekak_kab_sleman/model/web/usulan"
)

type UsulanTerpilihService interface {
	Create(ctx context.Context, request usulan.UsulanTerpilihCreateRequest) (usulan.UsulanTerpilihResponse, error)
	Delete(ctx context.Context, idUsulan string) error
}
