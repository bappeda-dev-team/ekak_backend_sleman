package repository

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/model/domain"
)

type UsulanInisiatifRepository interface {
	Create(ctx context.Context, tx *sql.Tx, usulan domain.UsulanInisiatif) (domain.UsulanInisiatif, error)
	Update(ctx context.Context, tx *sql.Tx, usulan domain.UsulanInisiatif) (domain.UsulanInisiatif, error)
	FindAll(ctx context.Context, tx *sql.Tx, pegawaiId *string, isActive *bool, rekinId *string) ([]domain.UsulanInisiatif, error)
	FindById(ctx context.Context, tx *sql.Tx, idUsulan string) (domain.UsulanInisiatif, error)
	Delete(ctx context.Context, tx *sql.Tx, idUsulan string) error
}
