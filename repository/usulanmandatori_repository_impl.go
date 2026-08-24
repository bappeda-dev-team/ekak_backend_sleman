package repository

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/domain"
	"fmt"
)

type UsulanMandatoriRepositoryImpl struct {
}

func NewUsulanMandatoriRepositoryImpl() *UsulanMandatoriRepositoryImpl {
	return &UsulanMandatoriRepositoryImpl{}
}

func (repository *UsulanMandatoriRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, usulan domain.UsulanMandatori) (domain.UsulanMandatori, error) {
	script := "INSERT INTO tb_usulan_mandatori (id, usulan, peraturan_terkait, uraian, tahun, rekin_id, pegawai_id, kode_opd, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, script, usulan.Id, usulan.Usulan, usulan.PeraturanTerkait, usulan.Uraian, usulan.Tahun, usulan.RekinId, usulan.PegawaiId, usulan.KodeOpd, usulan.Status)
	if err != nil {
		return domain.UsulanMandatori{}, fmt.Errorf("gagal membuat usulan mandatori: %w", err)
	}
	return usulan, nil
}

func (repository *UsulanMandatoriRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, kodeOpd *string, pegawaiId *string, isActive *bool, rekinId *string) ([]domain.UsulanMandatori, error) {
	script := "SELECT id, usulan, peraturan_terkait, uraian, tahun, rekin_id, pegawai_id, kode_opd, is_active, status, created_at FROM tb_usulan_mandatori WHERE 1=1"
	var args []interface{}

	if pegawaiId != nil {
		script += " AND pegawai_id = ?"
		args = append(args, *pegawaiId)
	}

	if isActive != nil {
		script += " AND is_active = ?"
		args = append(args, *isActive)
	}

	if rekinId != nil {
		script += " AND rekin_id = ?"
		args = append(args, *rekinId)
	}

	if kodeOpd != nil {
		script += " AND kode_opd = ?"
		args = append(args, *kodeOpd)
	}

	script += " order by created_at asc"

	rows, err := tx.QueryContext(ctx, script, args...)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil semua usulan mandatori: %w", err)
	}
	defer rows.Close()

	var usulanMandatori []domain.UsulanMandatori
	for rows.Next() {
		var usulan domain.UsulanMandatori
		err := rows.Scan(&usulan.Id, &usulan.Usulan, &usulan.PeraturanTerkait, &usulan.Uraian, &usulan.Tahun, &usulan.RekinId, &usulan.PegawaiId, &usulan.KodeOpd, &usulan.IsActive, &usulan.Status, &usulan.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai baris usulan mandatori: %w", err)
		}
		usulanMandatori = append(usulanMandatori, usulan)
	}
	return usulanMandatori, nil
}

func (repository *UsulanMandatoriRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, idUsulan string) (domain.UsulanMandatori, error) {
	script := "SELECT id, usulan, peraturan_terkait, uraian, tahun, rekin_id, pegawai_id, kode_opd, is_active, status, created_at FROM tb_usulan_mandatori WHERE id = ?"
	row := tx.QueryRowContext(ctx, script, idUsulan)

	var usulan domain.UsulanMandatori
	err := row.Scan(&usulan.Id, &usulan.Usulan, &usulan.PeraturanTerkait, &usulan.Uraian, &usulan.Tahun, &usulan.RekinId, &usulan.PegawaiId, &usulan.KodeOpd, &usulan.IsActive, &usulan.Status, &usulan.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.UsulanMandatori{}, fmt.Errorf("usulan mandatori dengan id %s tidak ditemukan", idUsulan)
		}
		return domain.UsulanMandatori{}, fmt.Errorf("gagal mencari usulan mandatori berdasarkan id: %w", err)
	}
	return usulan, nil
}

func (repository *UsulanMandatoriRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, usulan domain.UsulanMandatori) (domain.UsulanMandatori, error) {
	script := "UPDATE tb_usulan_mandatori SET usulan = ?, peraturan_terkait = ?, uraian = ?, tahun = ?, pegawai_id = ?, kode_opd = ?, status = ? WHERE id = ?"
	result, err := tx.ExecContext(ctx, script, usulan.Usulan, usulan.PeraturanTerkait, usulan.Uraian, usulan.Tahun, usulan.PegawaiId, usulan.KodeOpd, usulan.Status, usulan.Id)
	if err != nil {
		return domain.UsulanMandatori{}, fmt.Errorf("gagal memperbarui usulan mandatori: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return domain.UsulanMandatori{}, fmt.Errorf("gagal mendapatkan jumlah baris yang terpengaruh: %w", err)
	}
	if rowsAffected == 0 {
		return domain.UsulanMandatori{}, fmt.Errorf("usulan mandatori dengan id %s tidak ditemukan", usulan.Id)
	}
	return usulan, nil
}

func (repository *UsulanMandatoriRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, idUsulan string) error {
	script := "DELETE FROM tb_usulan_mandatori WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, idUsulan)
	helper.PanicIfError(err)
	return nil
}
