package rencanaaksi

type RencanaAksiUpdateRequest struct {
	Id              string `json:"id"`
	Urutan          int    `json:"urutan"`
	NamaRencanaAksi string `json:"nama_rencana_aksi"`
}
