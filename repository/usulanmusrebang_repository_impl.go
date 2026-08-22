package repository

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/model/domain"
	"fmt"
)

type UsulanMusrebangRepositoryImpl struct {
}

func NewUsulanMusrebangRepositoryImpl() *UsulanMusrebangRepositoryImpl {
	return &UsulanMusrebangRepositoryImpl{}
}

func (repository *UsulanMusrebangRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, usulan domain.UsulanMusrebang) (domain.UsulanMusrebang, error) {
	script := "INSERT INTO tb_usulan_musrebang (id, usulan, alamat, uraian, tahun, rekin_id, kode_opd, status) VALUES (?,?,?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, script, usulan.Id, usulan.Usulan, usulan.Alamat, usulan.Uraian, usulan.Tahun, usulan.RekinId, usulan.KodeOpd, usulan.Status)
	if err != nil {
		return domain.UsulanMusrebang{}, fmt.Errorf("error saat menyimpan usulan musrebang: %v", err)
	}
	return usulan, nil
}

func (repository *UsulanMusrebangRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, usulan domain.UsulanMusrebang) (domain.UsulanMusrebang, error) {
	script := "UPDATE tb_usulan_musrebang SET usulan = ?, alamat = ?, uraian = ?, tahun = ?, kode_opd = ?, status = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, usulan.Usulan, usulan.Alamat, usulan.Uraian, usulan.Tahun, usulan.KodeOpd, usulan.Status, usulan.Id)
	if err != nil {
		return domain.UsulanMusrebang{}, fmt.Errorf("error saat mengupdate usulan musrebang: %v", err)
	}
	return usulan, nil
}

func (repository *UsulanMusrebangRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, idUsulan string) (domain.UsulanMusrebang, error) {
	script := "SELECT id, usulan, alamat, uraian, tahun, rekin_id, kode_opd, is_active, status, created_at FROM tb_usulan_musrebang WHERE id = ?"
	row := tx.QueryRowContext(ctx, script, idUsulan)

	var usulan domain.UsulanMusrebang
	err := row.Scan(&usulan.Id, &usulan.Usulan, &usulan.Alamat, &usulan.Uraian, &usulan.Tahun, &usulan.RekinId, &usulan.KodeOpd, &usulan.IsActive, &usulan.Status, &usulan.CreatedAt)
	if err != nil {
		return domain.UsulanMusrebang{}, fmt.Errorf("error saat mencari usulan musrebang: %v", err)
	}
	return usulan, nil
}

func (repository *UsulanMusrebangRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, kodeOpd *string, is_active *bool, rekinId *string, status *string) ([]domain.UsulanMusrebang, error) {
	script := "SELECT id, usulan, alamat, uraian, tahun, rekin_id, kode_opd, is_active, status, created_at FROM tb_usulan_musrebang WHERE 1=1"
	var args []interface{}

	if kodeOpd != nil {
		script += " AND kode_opd = ?"
		args = append(args, *kodeOpd)
	}

	if is_active != nil {
		script += " AND is_active = ?"
		args = append(args, *is_active)
	}

	if rekinId != nil {
		script += " AND rekin_id = ?"
		args = append(args, *rekinId)
	}

	if status != nil {
		script += " AND status = ?"
		args = append(args, *status)
	}

	script += " order by created_at asc"

	rows, err := tx.QueryContext(ctx, script, args...)
	if err != nil {
		return []domain.UsulanMusrebang{}, fmt.Errorf("error saat mencari usulan musrebang: %v", err)
	}
	defer rows.Close()

	var usulanMusrebang []domain.UsulanMusrebang
	for rows.Next() {
		var usulan domain.UsulanMusrebang
		err := rows.Scan(&usulan.Id, &usulan.Usulan, &usulan.Alamat, &usulan.Uraian, &usulan.Tahun, &usulan.RekinId, &usulan.KodeOpd, &usulan.IsActive, &usulan.Status, &usulan.CreatedAt)
		if err != nil {
			return []domain.UsulanMusrebang{}, fmt.Errorf("error saat memindai usulan musrebang: %v", err)
		}
		usulanMusrebang = append(usulanMusrebang, usulan)
	}
	return usulanMusrebang, nil
}

func (repository *UsulanMusrebangRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, idUsulan string) error {
	script := "DELETE FROM tb_usulan_musrebang WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, idUsulan)
	if err != nil {
		return fmt.Errorf("error saat menghapus usulan musrebang: %v", err)
	}
	return nil
}

func (repository *UsulanMusrebangRepositoryImpl) CreateRekin(ctx context.Context, tx *sql.Tx, idUsulan string, rekinId string) error {
	script := "UPDATE tb_usulan_musrebang SET rekin_id = ?, status = 'usulan_diambil' WHERE id = ?"
	result, err := tx.ExecContext(ctx, script, rekinId, idUsulan)
	if err != nil {
		return fmt.Errorf("error saat mengupdate rekin usulan musrebang: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error saat memeriksa rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usulan musrebang dengan id %s tidak ditemukan", idUsulan)
	}

	return nil
}

func (repository *UsulanMusrebangRepositoryImpl) DeleteUsulanTerpilih(ctx context.Context, tx *sql.Tx, idUsulan string) error {
	script := "UPDATE tb_usulan_musrebang SET rekin_id = '', status = 'belum_diambil' WHERE id = ?"
	result, err := tx.ExecContext(ctx, script, idUsulan)
	if err != nil {
		return fmt.Errorf("error saat menghapus usulan terpilih: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error saat memeriksa rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usulan musrebang dengan id %s tidak ditemukan", idUsulan)
	}

	return nil
}
