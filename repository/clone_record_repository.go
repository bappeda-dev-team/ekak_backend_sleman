package repository

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/model/domain"
)

type CloneRecordRepository interface {
	Create(ctx context.Context, tx *sql.Tx, cloneRecord domain.CloneRecord) (domain.CloneRecord, error)
	GetCloneByKodeOpdTahunSumberTahunTujuan(ctx context.Context, tx *sql.Tx, kodeOpd string, tahunSumber string, tahunTujuan string) (domain.CloneRecord, error)
	UpdateStatus(ctx context.Context, tx *sql.Tx, id int, status string, errMsg string) error
}
