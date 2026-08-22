package repository

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/model/domain/domainmaster"
)

type LembagaRepository interface {
	Create(ctx context.Context, tx *sql.Tx, lembaga domainmaster.Lembaga) domainmaster.Lembaga
	Update(ctx context.Context, tx *sql.Tx, lembaga domainmaster.Lembaga) domainmaster.Lembaga
	Delete(ctx context.Context, tx *sql.Tx, id string) error
	FindById(ctx context.Context, tx *sql.Tx, id string) (domainmaster.Lembaga, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domainmaster.Lembaga, error)
}
