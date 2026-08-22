package repository

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/model/domain/domainmaster"
)

type JabatanPegawaiRepository interface {
	TambahJabatanPegawai(ctx context.Context, tx *sql.Tx, jabatanPegawai domainmaster.JabatanPegawai) error
}
