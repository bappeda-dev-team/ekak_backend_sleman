package dasarhukum

import "ekak_kab_sleman/model/web"

type DasarHukumResponse struct {
	Id               string             `json:"id"`
	RekinId          string             `json:"rencana_kinerja_id"`
	KodeOpd          string             `json:"kode_opd"`
	Urutan           int                `json:"urutan"`
	PeraturanTerkait string             `json:"peraturan_terkait"`
	Uraian           string             `json:"uraian"`
	Action           []web.ActionButton `json:"action,omitempty"`
}
