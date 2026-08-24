package repository

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/domain"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type VisiPemdaRepositoryImpl struct {
}

func NewVisiPemdaRepositoryImpl() *VisiPemdaRepositoryImpl {
	return &VisiPemdaRepositoryImpl{}
}

func (repository *VisiPemdaRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, visiPemda domain.VisiPemda) (domain.VisiPemda, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	visiPemda.Id = r.Intn(9000) + 1000

	script := "INSERT INTO tb_visi_pemda (id, visi, tahun_awal_periode, tahun_akhir_periode, jenis_periode, keterangan) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, script,
		visiPemda.Id,
		visiPemda.Visi,
		visiPemda.TahunAwalPeriode,
		visiPemda.TahunAkhirPeriode,
		visiPemda.JenisPeriode,
		visiPemda.Keterangan)
	if err != nil {
		return visiPemda, err
	}

	return visiPemda, nil
}

func (repository *VisiPemdaRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, visiPemda domain.VisiPemda) (domain.VisiPemda, error) {
	script := `UPDATE tb_visi_pemda SET 
        visi = ?, 
        tahun_awal_periode = ?, 
        tahun_akhir_periode = ?, 
        jenis_periode = ?, 
        keterangan = ? 
    WHERE id = ?`

	_, err := tx.ExecContext(ctx, script,
		visiPemda.Visi,
		visiPemda.TahunAwalPeriode,
		visiPemda.TahunAkhirPeriode,
		visiPemda.JenisPeriode,
		visiPemda.Keterangan,
		visiPemda.Id)
	if err != nil {
		return visiPemda, err
	}

	return visiPemda, nil
}

func (repository *VisiPemdaRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, visiPemdaId int) error {
	script := "DELETE FROM tb_visi_pemda WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, visiPemdaId)
	return err
}

func (repository *VisiPemdaRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, visiPemdaId int) (domain.VisiPemda, error) {
	script := "SELECT id, visi, tahun_awal_periode, tahun_akhir_periode, jenis_periode, keterangan FROM tb_visi_pemda WHERE id = ?"
	rows, err := tx.QueryContext(ctx, script, visiPemdaId)
	if err != nil {
		return domain.VisiPemda{}, err
	}
	defer rows.Close()

	var visiPemda domain.VisiPemda
	if rows.Next() {
		err := rows.Scan(&visiPemda.Id, &visiPemda.Visi, &visiPemda.TahunAwalPeriode, &visiPemda.TahunAkhirPeriode, &visiPemda.JenisPeriode, &visiPemda.Keterangan)
		helper.PanicIfError(err)
	}
	return visiPemda, nil
}

func (repository *VisiPemdaRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, tahunAwal string, tahunAkhir string, jenisPeriode string) ([]domain.VisiPemda, error) {
	var conditions []string
	var params []interface{}

	script := "SELECT id, visi, tahun_awal_periode, tahun_akhir_periode, jenis_periode, keterangan FROM tb_visi_pemda"

	if tahunAwal != "" && tahunAkhir != "" {
		conditions = append(conditions, "CAST(? AS SIGNED) BETWEEN CAST(tahun_awal_periode AS SIGNED) AND CAST(tahun_akhir_periode AS SIGNED)")
		params = append(params, tahunAwal)
	}

	if jenisPeriode != "" {
		conditions = append(conditions, "jenis_periode = ?")
		params = append(params, jenisPeriode)
	}

	if len(conditions) > 0 {
		script += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := tx.QueryContext(ctx, script, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var visiPemdaList []domain.VisiPemda
	for rows.Next() {
		var visiPemda domain.VisiPemda
		err := rows.Scan(
			&visiPemda.Id,
			&visiPemda.Visi,
			&visiPemda.TahunAwalPeriode,
			&visiPemda.TahunAkhirPeriode,
			&visiPemda.JenisPeriode,
			&visiPemda.Keterangan,
		)
		if err != nil {
			return nil, err
		}
		visiPemdaList = append(visiPemdaList, visiPemda)
	}

	return visiPemdaList, nil
}

// Tambahkan fungsi baru untuk mendapatkan visi berdasarkan ID
func (repository *VisiPemdaRepositoryImpl) FindByIdWithDefault(ctx context.Context, tx *sql.Tx, visiPemdaId int) (domain.VisiPemda, error) {
	if visiPemdaId == 0 {
		return domain.VisiPemda{
			Id:                0,
			Visi:              "Belum ada visi",
			TahunAwalPeriode:  "",
			TahunAkhirPeriode: "",
			JenisPeriode:      "",
			Keterangan:        "",
		}, nil
	}

	script := "SELECT id, visi, tahun_awal_periode, tahun_akhir_periode, jenis_periode, keterangan FROM tb_visi_pemda WHERE id = ?"
	var visiPemda domain.VisiPemda
	err := tx.QueryRowContext(ctx, script, visiPemdaId).Scan(
		&visiPemda.Id,
		&visiPemda.Visi,
		&visiPemda.TahunAwalPeriode,
		&visiPemda.TahunAkhirPeriode,
		&visiPemda.JenisPeriode,
		&visiPemda.Keterangan,
	)

	if err == sql.ErrNoRows {
		return domain.VisiPemda{}, fmt.Errorf("visi pemda dengan id %d tidak ditemukan", visiPemdaId)
	}
	if err != nil {
		return domain.VisiPemda{}, err
	}

	return visiPemda, nil
}
