package domain

import "time"

type UsulanMandatori struct {
	Id               string
	Usulan           string
	PeraturanTerkait string
	Uraian           string
	Tahun            string
	RekinId          string
	PegawaiId        string
	NamaPegawai      string
	KodeOpd          string
	IsActive         bool
	Status           string
	CreatedAt        time.Time
}
