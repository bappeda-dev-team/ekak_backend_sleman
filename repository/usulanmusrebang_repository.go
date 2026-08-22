package repository

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/model/domain"
)

type UsulanMusrebangRepository interface {
	Create(ctx context.Context, tx *sql.Tx, usulan domain.UsulanMusrebang) (domain.UsulanMusrebang, error)
	Update(ctx context.Context, tx *sql.Tx, usulan domain.UsulanMusrebang) (domain.UsulanMusrebang, error)
	FindById(ctx context.Context, tx *sql.Tx, idUsulan string) (domain.UsulanMusrebang, error)
	FindAll(ctx context.Context, tx *sql.Tx, kodeOpd *string, is_active *bool, rekinId *string, status *string) ([]domain.UsulanMusrebang, error)
	Delete(ctx context.Context, tx *sql.Tx, idUsulan string) error
	CreateRekin(ctx context.Context, tx *sql.Tx, idUsulan string, rekinId string) error
	DeleteUsulanTerpilih(ctx context.Context, tx *sql.Tx, idUsulan string) error
}
