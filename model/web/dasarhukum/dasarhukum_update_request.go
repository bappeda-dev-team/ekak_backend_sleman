package dasarhukum

type DasarHukumUpdateRequest struct {
	Id               string `json:"id"`
	RekinId          string `json:"rencana_kinerja_id"`
	Urutan           int    `json:"urutan"`
	PeraturanTerkait string `json:"peraturan_terkait"`
	Uraian           string `json:"uraian"`
}
