package domain

import "time"

type UsulanPokokPikiran struct {
	Id        string
	Usulan    string
	Alamat    string
	Uraian    string
	Tahun     string
	RekinId   string
	KodeOpd   string
	NamaOpd   string
	IsActive  bool
	Status    string
	CreatedAt time.Time
}
