package gambaranumum

type GambaranUmumCreateRequest struct {
	RekinId      string `json:"rekin_id"`
	KodeOpd      string `json:"kode_opd"`
	Urutan       int    `json:"urutan"`
	GambaranUmum string `json:"gambaran_umum"`
}
