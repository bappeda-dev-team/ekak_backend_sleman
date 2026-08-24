package service

import (
	"context"
	"ekak_kab_sleman/model/web/urusanrespon"
)

type UrusanService interface {
	Create(ctx context.Context, request urusanrespon.UrusanCreateRequest) (urusanrespon.UrusanResponse, error)
	Update(ctx context.Context, request urusanrespon.UrusanUpdateRequest) (urusanrespon.UrusanResponse, error)
	FindById(ctx context.Context, id string) (urusanrespon.UrusanResponse, error)
	FindAll(ctx context.Context) ([]urusanrespon.UrusanResponse, error)
	Delete(ctx context.Context, id string) error
	FindByKodeOpd(ctx context.Context, kodeOpd string) ([]urusanrespon.UrusanResponse, error)
	FindUrusanAndBidangByKodeOpd(ctx context.Context, kodeOpd string) ([]urusanrespon.UrusanResponse, error)
}
