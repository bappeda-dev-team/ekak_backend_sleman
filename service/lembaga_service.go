package service

import (
	"context"
	"ekak_kab_sleman/model/web/lembaga"
)

type LembagaService interface {
	Create(ctx context.Context, request lembaga.LembagaCreateRequest) (lembaga.LembagaResponse, error)
	Update(ctx context.Context, request lembaga.LembagaUpdateRequest) (lembaga.LembagaResponse, error)
	FindById(ctx context.Context, id string) (lembaga.LembagaResponse, error)
	FindAll(ctx context.Context) ([]lembaga.LembagaResponse, error)
	Delete(ctx context.Context, id string) error
}
