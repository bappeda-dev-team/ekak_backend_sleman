package service

import (
	"context"
	visimisipemda "ekak_kab_sleman/model/web/visimisi"
)

type VisiPemdaService interface {
	Create(ctx context.Context, request visimisipemda.VisiPemdaCreateRequest) (visimisipemda.VisiPemdaResponse, error)
	Update(ctx context.Context, request visimisipemda.VisiPemdaUpdateRequest) (visimisipemda.VisiPemdaResponse, error)
	Delete(ctx context.Context, visiPemdaId int) error
	FindById(ctx context.Context, visiPemdaId int) (visimisipemda.VisiPemdaResponse, error)
	FindAll(ctx context.Context, tahunAwal string, tahunAkhir string, jenisPeriode string) ([]visimisipemda.VisiPemdaResponse, error)
}
