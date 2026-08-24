package service

import (
	"context"
	"ekak_kab_sleman/model/web/programkegiatan"
)

type ProgramService interface {
	Create(ctx context.Context, request programkegiatan.ProgramKegiatanCreateRequest) (programkegiatan.ProgramKegiatanResponse, error)
	Update(ctx context.Context, request programkegiatan.ProgramKegiatanUpdateRequest) (programkegiatan.ProgramKegiatanResponse, error)
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (programkegiatan.ProgramKegiatanResponse, error)
	FindAll(ctx context.Context) ([]programkegiatan.ProgramKegiatanResponse, error)
}
