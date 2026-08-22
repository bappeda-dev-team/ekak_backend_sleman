package repository

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/model/domain"
)

type PelaksanaanRencanaAksiRepository interface {
	Create(ctx context.Context, tx *sql.Tx, pelaksanaanRencanaAksi domain.PelaksanaanRencanaAksi) (domain.PelaksanaanRencanaAksi, error)
	Update(ctx context.Context, tx *sql.Tx, pelaksanaanRencanaAksi domain.PelaksanaanRencanaAksi) (domain.PelaksanaanRencanaAksi, error)
	Delete(ctx context.Context, tx *sql.Tx, id string) error
	FindById(ctx context.Context, tx *sql.Tx, id string) (domain.PelaksanaanRencanaAksi, error)
	FindByRencanaAksiId(ctx context.Context, tx *sql.Tx, rencanaAksiId string) ([]domain.PelaksanaanRencanaAksi, error)
	ExistsByRencanaAksiIdAndBulan(ctx context.Context, tx *sql.Tx, rencanaAksiId string, bulan int) (bool, error)
}
