package repository

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/model/domain"
)

type VisiPemdaRepository interface {
	Create(ctx context.Context, tx *sql.Tx, visiPemda domain.VisiPemda) (domain.VisiPemda, error)
	Update(ctx context.Context, tx *sql.Tx, visiPemda domain.VisiPemda) (domain.VisiPemda, error)
	Delete(ctx context.Context, tx *sql.Tx, visiPemdaId int) error
	FindById(ctx context.Context, tx *sql.Tx, visiPemdaId int) (domain.VisiPemda, error)
	FindAll(ctx context.Context, tx *sql.Tx, tahunAwal string, tahunAkhir string, jenisPeriode string) ([]domain.VisiPemda, error)
	FindByIdWithDefault(ctx context.Context, tx *sql.Tx, visiPemdaId int) (domain.VisiPemda, error)
}
