package repository

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/model/domain"
)

type RoleRepository interface {
	Create(ctx context.Context, tx *sql.Tx, role domain.Roles) (domain.Roles, error)
	Update(ctx context.Context, tx *sql.Tx, role domain.Roles) (domain.Roles, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Roles, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Roles, error)
}
