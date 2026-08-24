package repository

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/model/domain"
	"fmt"
)

type RoleRepositoryImpl struct {
}

func NewRoleRepositoryImpl() *RoleRepositoryImpl {
	return &RoleRepositoryImpl{}
}

func (repository *RoleRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, role domain.Roles) (domain.Roles, error) {
	script := "INSERT INTO tb_role (role) VALUES (?)"
	result, err := tx.ExecContext(ctx, script, role.Role)
	if err != nil {
		return domain.Roles{}, fmt.Errorf("error saat menyimpan role: %v", err)
	}

	// Mengambil ID yang baru dibuat
	id, err := result.LastInsertId()
	if err != nil {
		return domain.Roles{}, fmt.Errorf("error saat mengambil id: %v", err)
	}

	// Set ID ke struct role
	role.Id = int(id)
	return role, nil
}

func (repository *RoleRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, role domain.Roles) (domain.Roles, error) {
	script := "UPDATE tb_role SET role = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, role.Role, role.Id)
	if err != nil {
		return domain.Roles{}, fmt.Errorf("error saat mengupdate role: %v", err)
	}
	return role, nil
}

func (repository *RoleRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	script := "DELETE FROM tb_role WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, id)
	if err != nil {
		return fmt.Errorf("error saat menghapus role: %v", err)
	}
	return nil
}

func (repository *RoleRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Roles, error) {
	script := "SELECT id, role FROM tb_role"
	rows, err := tx.QueryContext(ctx, script)
	if err != nil {
		return []domain.Roles{}, fmt.Errorf("error saat mengambil semua role: %v", err)
	}
	defer rows.Close()

	var roles []domain.Roles
	for rows.Next() {
		var role domain.Roles
		err := rows.Scan(&role.Id, &role.Role)
		if err != nil {
			return []domain.Roles{}, fmt.Errorf("error saat mengambil semua role: %v", err)
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (repository *RoleRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Roles, error) {
	script := "SELECT id, role FROM tb_role WHERE id = ?"
	var role domain.Roles
	err := tx.QueryRowContext(ctx, script, id).Scan(&role.Id, &role.Role)
	if err != nil {
		return domain.Roles{}, fmt.Errorf("error saat mengambil role berdasarkan id: %v", err)
	}
	return role, nil
}
