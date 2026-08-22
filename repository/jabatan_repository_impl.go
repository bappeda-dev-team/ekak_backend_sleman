package repository

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/model/domain/domainmaster"
)

type JabatanRepositoryImpl struct {
}

func NewJabatanRepositoryImpl() *JabatanRepositoryImpl {
	return &JabatanRepositoryImpl{}
}

func (repository *JabatanRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, jabatan domainmaster.Jabatan) domainmaster.Jabatan {
	script := "INSERT INTO tb_jabatan (id, kode_jabatan, nama_jabatan, kelas_jabatan, jenis_jabatan, nilai_jabatan, kode_opd, index_jabatan, tahun, esselon) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, script, jabatan.Id, jabatan.KodeJabatan, jabatan.NamaJabatan, jabatan.KelasJabatan, jabatan.JenisJabatan, jabatan.NilaiJabatan, jabatan.KodeOpd, jabatan.IndexJabatan, jabatan.Tahun, jabatan.Esselon)
	if err != nil {
		return jabatan
	}
	return jabatan
}

func (repository *JabatanRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, jabatan domainmaster.Jabatan) domainmaster.Jabatan {
	script := "UPDATE tb_jabatan SET kode_jabatan = ?, nama_jabatan = ?, kelas_jabatan = ?, jenis_jabatan = ?, nilai_jabatan = ?, kode_opd = ?, index_jabatan = ?, tahun = ?, esselon = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, jabatan.KodeJabatan, jabatan.NamaJabatan, jabatan.KelasJabatan, jabatan.JenisJabatan, jabatan.NilaiJabatan, jabatan.KodeOpd, jabatan.IndexJabatan, jabatan.Tahun, jabatan.Esselon, jabatan.Id)
	if err != nil {
		return jabatan
	}
	return jabatan
}

func (repository *JabatanRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) error {
	script := "DELETE FROM tb_jabatan WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, id)
	if err != nil {
		return err
	}
	return nil
}

func (repository *JabatanRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id string) (domainmaster.Jabatan, error) {
	script := "SELECT id, kode_jabatan, nama_jabatan, kelas_jabatan, jenis_jabatan, nilai_jabatan, kode_opd, index_jabatan, tahun, esselon FROM tb_jabatan WHERE id = ?"
	var jabatan domainmaster.Jabatan
	err := tx.QueryRowContext(ctx, script, id).Scan(&jabatan.Id, &jabatan.KodeJabatan, &jabatan.NamaJabatan, &jabatan.KelasJabatan, &jabatan.JenisJabatan, &jabatan.NilaiJabatan, &jabatan.KodeOpd, &jabatan.IndexJabatan, &jabatan.Tahun, &jabatan.Esselon)
	if err != nil {
		return jabatan, err
	}
	return jabatan, nil
}

func (repository *JabatanRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, kodeOpd string, tahun string) ([]domainmaster.Jabatan, error) {
	script := "SELECT id, kode_jabatan, nama_jabatan, kelas_jabatan, jenis_jabatan, nilai_jabatan, kode_opd, index_jabatan, tahun, esselon FROM tb_jabatan WHERE 1=1"
	params := []interface{}{}

	if kodeOpd != "" {
		script += " AND kode_opd = ?"
		params = append(params, kodeOpd)
	}
	if tahun != "" {
		script += " AND tahun = ?"
		params = append(params, tahun)
	}
	rows, err := tx.QueryContext(ctx, script, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jabatans []domainmaster.Jabatan
	for rows.Next() {
		jabatan := domainmaster.Jabatan{}
		rows.Scan(&jabatan.Id, &jabatan.KodeJabatan, &jabatan.NamaJabatan, &jabatan.KelasJabatan, &jabatan.JenisJabatan, &jabatan.NilaiJabatan, &jabatan.KodeOpd, &jabatan.IndexJabatan, &jabatan.Tahun, &jabatan.Esselon)
		jabatans = append(jabatans, jabatan)
	}
	return jabatans, nil
}
