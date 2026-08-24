package dataseeder

import (
	"context"
	"database/sql"
)

type PegawaiSeeder interface {
	Seed(ctx context.Context, tx *sql.Tx) error
}
