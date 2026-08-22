package dataseeder

import (
	"context"
	"database/sql"
)

type RoleSeeder interface {
	Seed(ctx context.Context, tx *sql.Tx) error
}
