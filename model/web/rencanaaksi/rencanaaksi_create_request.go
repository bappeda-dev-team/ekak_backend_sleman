package rencanaaksi

type RencanaAksiCreateRequest struct {
	RencanaKinerjaId string `json:"rekin_id"`
	KodeOpd          string `json:"kode_opd"`
	Urutan           int    `json:"urutan"`
	NamaRencanaAksi  string `json:"nama_rencana_aksi"`
}
