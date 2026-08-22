package rencanaaksi

import "github.com/google/uuid"

type PelaksanaanRencanaAksiCreateRequest struct {
	Id            uuid.UUID `json:"id"`
	RencanaAksiId string    `json:"rencana_aksi_id"`
	Bobot         int       `json:"bobot"`
	Bulan         int       `json:"bulan"`
}
