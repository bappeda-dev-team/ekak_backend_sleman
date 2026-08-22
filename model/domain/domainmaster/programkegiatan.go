package domainmaster

import "ekak_kab_sleman/model/domain"

type ProgramKegiatan struct {
	Id          string
	KodeProgram string
	NamaProgram string
	KodeOPD     string
	IsActive    bool
	Tahun       string
	Indikator   []domain.Indikator
}
