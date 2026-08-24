package dasarhukum

type DasarHukumCreateRequest struct {
	RekinId          string `json:"rencana_kinerja_id"`
	KodeOpd          string `json:"kode_opd"`
	Urutan           int    `json:"urutan"`
	PeraturanTerkait string `json:"peraturan_terkait"`
	Uraian           string `json:"uraian"`
}
