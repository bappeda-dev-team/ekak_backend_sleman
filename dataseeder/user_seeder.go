package dataseeder

import (
	"context"
	"database/sql"
)

type UserSeeder interface {
	Seed(ctx context.Context, tx *sql.Tx) error
}
