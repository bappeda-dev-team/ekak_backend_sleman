package repository

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/model/domain"
	"ekak_kab_sleman/model/domain/domainmaster"
)

type KegiatanRepository interface {
	Create(ctx context.Context, tx *sql.Tx, kegiatan domainmaster.Kegiatan) (domainmaster.Kegiatan, error)
	Update(ctx context.Context, tx *sql.Tx, kegiatan domainmaster.Kegiatan) (domainmaster.Kegiatan, error)
	Delete(ctx context.Context, tx *sql.Tx, id string) error
	FindById(ctx context.Context, tx *sql.Tx, id string) (domainmaster.Kegiatan, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domainmaster.Kegiatan, error)
	FindIndikatorByKegiatanId(ctx context.Context, tx *sql.Tx, kegiatanId string) ([]domain.Indikator, error)
	FindTargetByIndikatorId(ctx context.Context, tx *sql.Tx, indikatorId string) ([]domain.Target, error)
}
