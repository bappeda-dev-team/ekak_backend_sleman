package service

import (
	"context"
	"ekak_kab_sleman/model/web/permasalahan"
)

type PermasalahanRekinService interface {
	Create(ctx context.Context, request permasalahan.PermasalahanRekinCreateRequest) (permasalahan.PermasalahanRekinResponse, error)
	Update(ctx context.Context, request permasalahan.PermasalahanRekinUpdateRequest) (permasalahan.PermasalahanRekinResponse, error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context, rekinId *string) ([]permasalahan.PermasalahanRekinResponse, error)
	FindById(ctx context.Context, id int) (permasalahan.PermasalahanRekinResponse, error)
}
