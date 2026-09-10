package internal

import (
	"context"
)

type DataMasterClient interface {
	FindMappingOpd(ctx context.Context, kodeOpd string) (*MappingOpd, error)
	FindPegawaiByKodeOpd(
		ctx context.Context,
		kodeOpdMaster string,
	) ([]Pegawai, error)
}
