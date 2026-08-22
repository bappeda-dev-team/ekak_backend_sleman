package repository

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/model/domain"
)

type UsulanPokokPikiranRepository interface {
	Create(ctx context.Context, tx *sql.Tx, usulan domain.UsulanPokokPikiran) (domain.UsulanPokokPikiran, error)
	Update(ctx context.Context, tx *sql.Tx, usulan domain.UsulanPokokPikiran) (domain.UsulanPokokPikiran, error)
	FindAll(ctx context.Context, tx *sql.Tx, kodeOpd *string, isActive *bool, rekinId *string, status *string) ([]domain.UsulanPokokPikiran, error)
	FindById(ctx context.Context, tx *sql.Tx, idUsulan string) (domain.UsulanPokokPikiran, error)
	Delete(ctx context.Context, tx *sql.Tx, idUsulan string) error
	CreateRekin(ctx context.Context, tx *sql.Tx, idUsulan string, rekinId string) error
	DeleteUsulanTerpilih(ctx context.Context, tx *sql.Tx, idUsulan string) error
}
