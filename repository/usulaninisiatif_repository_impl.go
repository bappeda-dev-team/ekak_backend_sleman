package repository

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/model/domain"
	"fmt"
)

type UsulanInisiatifRepositoryImpl struct {
}

func NewUsulanInisiatifRepositoryImpl() *UsulanInisiatifRepositoryImpl {
	return &UsulanInisiatifRepositoryImpl{}
}

func (repository *UsulanInisiatifRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, usulan domain.UsulanInisiatif) (domain.UsulanInisiatif, error) {
	script := "INSERT INTO tb_usulan_inisiatif (id, usulan, manfaat, uraian, tahun, rekin_id, pegawai_id, kode_opd, status) VALUES (?,?,?,?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, script, usulan.Id, usulan.Usulan, usulan.Manfaat, usulan.Uraian, usulan.Tahun, usulan.RekinId, usulan.PegawaiId, usulan.KodeOpd, usulan.Status)
	if err != nil {
		return domain.UsulanInisiatif{}, fmt.Errorf("error saat membuat usulan inovasi: %v", err)
	}
	return usulan, nil
}

func (repository *UsulanInisiatifRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, usulan domain.UsulanInisiatif) (domain.UsulanInisiatif, error) {
	script := "UPDATE tb_usulan_inisiatif SET usulan = ?, manfaat = ?, uraian = ?, tahun = ?, pegawai_id = ?, kode_opd = ?, status = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, usulan.Usulan, usulan.Manfaat, usulan.Uraian, usulan.Tahun, usulan.PegawaiId, usulan.KodeOpd, usulan.Status, usulan.Id)
	if err != nil {
		return domain.UsulanInisiatif{}, fmt.Errorf("error saat mengupdate usulan inovasi: %v", err)
	}
	return usulan, nil
}

func (repository *UsulanInisiatifRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, pegawaiId *string, isActive *bool, rekinId *string) ([]domain.UsulanInisiatif, error) {
	script := "SELECT id, usulan, manfaat, uraian, tahun, rekin_id, pegawai_id, kode_opd, is_active, status, created_at FROM tb_usulan_inisiatif WHERE 1=1"
	var params []interface{}

	if pegawaiId != nil {
		script += " AND pegawai_id = ?"
		params = append(params, *pegawaiId)
	}

	if isActive != nil {
		script += " AND is_active = ?"
		params = append(params, *isActive)
	}

	if rekinId != nil {
		script += " AND rekin_id = ?"
		params = append(params, *rekinId)
	}

	script += " order by created_at asc"

	rows, err := tx.QueryContext(ctx, script, params...)
	if err != nil {
		return nil, fmt.Errorf("error saat mencari usulan inovasi: %v", err)
	}

	defer rows.Close()

	var usulanInovasi []domain.UsulanInisiatif
	for rows.Next() {
		var usulan domain.UsulanInisiatif
		err := rows.Scan(&usulan.Id, &usulan.Usulan, &usulan.Manfaat, &usulan.Uraian, &usulan.Tahun, &usulan.RekinId, &usulan.PegawaiId, &usulan.KodeOpd, &usulan.IsActive, &usulan.Status, &usulan.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error saat mencari usulan inovasi: %v", err)
		}
		usulanInovasi = append(usulanInovasi, usulan)
	}

	return usulanInovasi, nil
}

func (repository *UsulanInisiatifRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, idUsulan string) (domain.UsulanInisiatif, error) {
	script := "SELECT id, usulan, manfaat, uraian, tahun, rekin_id, pegawai_id, kode_opd, is_active, status, created_at FROM tb_usulan_inisiatif WHERE id = ?"
	row := tx.QueryRowContext(ctx, script, idUsulan)

	var usulan domain.UsulanInisiatif
	err := row.Scan(&usulan.Id, &usulan.Usulan, &usulan.Manfaat, &usulan.Uraian, &usulan.Tahun, &usulan.RekinId, &usulan.PegawaiId, &usulan.KodeOpd, &usulan.IsActive, &usulan.Status, &usulan.CreatedAt)
	if err != nil {
		return domain.UsulanInisiatif{}, fmt.Errorf("error saat mencari usulan inovasi: %v", err)
	}
	return usulan, nil
}

func (repository *UsulanInisiatifRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, idUsulan string) error {
	script := "DELETE FROM tb_usulan_inisiatif WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, idUsulan)
	if err != nil {
		return fmt.Errorf("error saat menghapus usulan inovasi: %v", err)
	}
	return nil
}
