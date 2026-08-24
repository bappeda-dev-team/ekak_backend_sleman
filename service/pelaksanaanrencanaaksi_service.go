package service

import (
	"context"
	"ekak_kab_sleman/model/web/rencanaaksi"
)

type PelaksanaanRencanaAksiService interface {
	Create(ctx context.Context, request rencanaaksi.PelaksanaanRencanaAksiCreateRequest) (rencanaaksi.PelaksanaanRencanaAksiResponse, error)
	Update(ctx context.Context, request rencanaaksi.PelaksanaanRencanaAksiUpdateRequest) (rencanaaksi.PelaksanaanRencanaAksiResponse, error)
	FindById(ctx context.Context, id string) (rencanaaksi.PelaksanaanRencanaAksiResponse, error)
	FindByRencanaAksiId(ctx context.Context, rencanaAksiId string) ([]rencanaaksi.PelaksanaanRencanaAksiResponse, error)
	Delete(ctx context.Context, id string) error
}
