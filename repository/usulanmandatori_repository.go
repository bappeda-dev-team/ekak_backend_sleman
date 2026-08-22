package repository

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/model/domain"
)

type UsulanMandatoriRepository interface {
	Create(ctx context.Context, tx *sql.Tx, usulan domain.UsulanMandatori) (domain.UsulanMandatori, error)
	FindAll(ctx context.Context, tx *sql.Tx, kodeOpd *string, pegawaiId *string, isActive *bool, rekinId *string) ([]domain.UsulanMandatori, error)
	FindById(ctx context.Context, tx *sql.Tx, idUsulan string) (domain.UsulanMandatori, error)
	Update(ctx context.Context, tx *sql.Tx, usulan domain.UsulanMandatori) (domain.UsulanMandatori, error)
	Delete(ctx context.Context, tx *sql.Tx, idUsulan string) error
}
