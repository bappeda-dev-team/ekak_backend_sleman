package service

import (
	"context"
	"ekak_kab_sleman/model/web/tujuanpemda"
)

type TujuanPemdaService interface {
	Create(ctx context.Context, request tujuanpemda.TujuanPemdaCreateRequest) (tujuanpemda.TujuanPemdaResponse, error)
	Update(ctx context.Context, request tujuanpemda.TujuanPemdaUpdateRequest) (tujuanpemda.TujuanPemdaResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, tujuanPemdaId int) (tujuanpemda.TujuanPemdaResponse, error)
	FindAll(ctx context.Context, tahun string, jenisPeriode string) ([]tujuanpemda.TujuanPemdaResponse, error)
	UpdatePeriode(ctx context.Context, request tujuanpemda.TujuanPemdaUpdateRequest) (tujuanpemda.TujuanPemdaResponse, error)
	FindAllWithPokin(ctx context.Context, tahunAwal string, tahunAkhir string, jenisPeriode string) ([]tujuanpemda.TujuanPemdaWithPokinResponse, error)
	FindPokinWithPeriode(ctx context.Context, pokinId int, jenisPeriode string) (tujuanpemda.PokinWithPeriodeResponse, error)
}
