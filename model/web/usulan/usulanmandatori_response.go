package usulan

import "ekak_kab_sleman/model/web"

type UsulanMandatoriResponse struct {
	Id               string             `json:"id"`
	Usulan           string             `json:"usulan"`
	PeraturanTerkait string             `json:"peraturan_terkait"`
	Uraian           string             `json:"uraian"`
	Tahun            string             `json:"tahun"`
	RekinId          string             `json:"rencana_kinerja_id,omitempty"`
	PegawaiId        string             `json:"pegawai_id"`
	NamaPegawai      string             `json:"nama_pegawai"`
	KodeOpd          string             `json:"kode_opd"`
	NamaOpd          string             `json:"nama_opd"`
	IsActive         bool               `json:"is_active,omitempty"`
	Status           string             `json:"status"`
	CreatedAt        string             `json:"dibuat_pada" time_format:"2006-01-02 15:04:05"`
	Action           []web.ActionButton `json:"action,omitempty"`
}
