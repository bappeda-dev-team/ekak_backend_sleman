package repository

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/model/domain"
)

type UsulanTerpilihRepository interface {
	Create(ctx context.Context, tx *sql.Tx, usulan domain.UsulanTerpilih) (domain.UsulanTerpilih, error)
	Delete(ctx context.Context, tx *sql.Tx, idUsulan string) error
	ExistsByJenisAndUsulanId(ctx context.Context, tx *sql.Tx, jenisUsulan string, usulanId string) (bool, error)
	ValidateJenisAndUsulanId(ctx context.Context, tx *sql.Tx, jenisUsulan string, usulanId string) (bool, error)
}
