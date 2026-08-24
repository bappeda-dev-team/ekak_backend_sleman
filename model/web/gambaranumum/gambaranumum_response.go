package gambaranumum

import "ekak_kab_sleman/model/web"

type GambaranUmumResponse struct {
	Id           string             `json:"id"`
	RekinId      string             `json:"rencana_kinerja_id"`
	KodeOpd      string             `json:"kode_opd"`
	Urutan       int                `json:"urutan"`
	GambaranUmum string             `json:"gambaran_umum"`
	Action       []web.ActionButton `json:"action,omitempty"`
}
