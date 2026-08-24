package domain

import "time"

type UsulanInisiatif struct {
	Id          string
	Usulan      string
	Manfaat     string
	Uraian      string
	Tahun       string
	RekinId     string
	PegawaiId   string
	NamaPegawai string
	KodeOpd     string
	NamaOpd     string
	IsActive    bool
	Status      string
	CreatedAt   time.Time
}
