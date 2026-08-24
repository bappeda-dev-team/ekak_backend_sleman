package repository

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/model/domain"
	"fmt"
)

type UsulanPokokPikiranRepositoryImpl struct {
}

func NewUsulanPokokPikiranRepositoryImpl() *UsulanPokokPikiranRepositoryImpl {
	return &UsulanPokokPikiranRepositoryImpl{}
}

func (repository *UsulanPokokPikiranRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, usulan domain.UsulanPokokPikiran) (domain.UsulanPokokPikiran, error) {
	script := "INSERT INTO tb_usulan_pokok_pikiran (id, usulan, alamat, uraian, tahun, rekin_id, kode_opd, status) VALUES (?,?,?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, script, usulan.Id, usulan.Usulan, usulan.Alamat, usulan.Uraian, usulan.Tahun, usulan.RekinId, usulan.KodeOpd, usulan.Status)
	if err != nil {
		return domain.UsulanPokokPikiran{}, fmt.Errorf("error saat menyimpan usulan pokok pikiran: %v", err)
	}
	return usulan, nil
}

func (repository *UsulanPokokPikiranRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, usulan domain.UsulanPokokPikiran) (domain.UsulanPokokPikiran, error) {
	script := "UPDATE tb_usulan_pokok_pikiran SET usulan = ?, alamat = ?, uraian = ?, tahun = ?, kode_opd = ?, status = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, usulan.Usulan, usulan.Alamat, usulan.Uraian, usulan.Tahun, usulan.KodeOpd, usulan.Status, usulan.Id)
	if err != nil {
		return domain.UsulanPokokPikiran{}, fmt.Errorf("error saat mengupdate usulan pokok pikiran: %v", err)
	}
	return usulan, nil
}

func (repository *UsulanPokokPikiranRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, idUsulan string) (domain.UsulanPokokPikiran, error) {
	script := "SELECT id, usulan, alamat, uraian, tahun, kode_opd, is_active, status, created_at FROM tb_usulan_pokok_pikiran WHERE id = ?"
	row := tx.QueryRowContext(ctx, script, idUsulan)

	var usulan domain.UsulanPokokPikiran
	err := row.Scan(&usulan.Id, &usulan.Usulan, &usulan.Alamat, &usulan.Uraian, &usulan.Tahun, &usulan.KodeOpd, &usulan.IsActive, &usulan.Status, &usulan.CreatedAt)
	if err != nil {
		return domain.UsulanPokokPikiran{}, fmt.Errorf("error saat mencari usulan pokok pikiran: %v", err)
	}

	return usulan, nil
}

func (repository *UsulanPokokPikiranRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, kodeOpd *string, isActive *bool, rekinId *string, status *string) ([]domain.UsulanPokokPikiran, error) {
	script := "SELECT id, usulan, alamat, uraian, tahun, rekin_id, kode_opd, is_active, status, created_at FROM tb_usulan_pokok_pikiran WHERE 1=1"
	var params []interface{}

	if kodeOpd != nil {
		script += " AND kode_opd = ?"
		params = append(params, *kodeOpd)
	}

	if isActive != nil {
		script += " AND is_active = ?"
		params = append(params, *isActive)
	}

	if rekinId != nil {
		script += " AND rekin_id = ?"
		params = append(params, *rekinId)
	}

	if status != nil {
		script += " AND status = ?"
		params = append(params, *status)
	}

	script += " ORDER BY created_at ASC"

	rows, err := tx.QueryContext(ctx, script, params...)
	if err != nil {
		return nil, fmt.Errorf("error saat mencari usulan pokok pikiran: %v", err)
	}

	defer rows.Close()

	var usulans []domain.UsulanPokokPikiran
	for rows.Next() {
		var usulan domain.UsulanPokokPikiran
		err := rows.Scan(&usulan.Id, &usulan.Usulan, &usulan.Alamat, &usulan.Uraian, &usulan.Tahun, &usulan.RekinId, &usulan.KodeOpd, &usulan.IsActive, &usulan.Status, &usulan.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error saat membaca usulan pokok pikiran: %v", err)
		}
		usulans = append(usulans, usulan)
	}

	return usulans, nil
}

func (repository *UsulanPokokPikiranRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, idUsulan string) error {
	script := "DELETE FROM tb_usulan_pokok_pikiran WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, idUsulan)
	if err != nil {
		return fmt.Errorf("error saat menghapus usulan pokok pikiran: %v", err)
	}
	return nil
}

func (repository *UsulanPokokPikiranRepositoryImpl) CreateRekin(ctx context.Context, tx *sql.Tx, idUsulan string, rekinId string) error {
	script := "UPDATE tb_usulan_pokok_pikiran SET rekin_id = ?, status = 'usulan_diambil' WHERE id = ?"
	result, err := tx.ExecContext(ctx, script, rekinId, idUsulan)
	if err != nil {
		return fmt.Errorf("error saat mengupdate rekin usulan musrebang: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error saat memeriksa rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usulan pokok pikiran dengan id %s tidak ditemukan", idUsulan)
	}

	return nil
}

func (repository *UsulanPokokPikiranRepositoryImpl) DeleteUsulanTerpilih(ctx context.Context, tx *sql.Tx, idUsulan string) error {
	script := "UPDATE tb_usulan_pokok_pikiran SET rekin_id = '', status = 'belum_diambil' WHERE id = ?"
	result, err := tx.ExecContext(ctx, script, idUsulan)
	if err != nil {
		return fmt.Errorf("error saat menghapus usulan terpilih: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error saat memeriksa rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usulan pokok pikiran dengan id %s tidak ditemukan", idUsulan)
	}

	return nil
}
