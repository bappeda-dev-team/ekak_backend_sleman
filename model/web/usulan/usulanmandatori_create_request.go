package usulan

type UsulanMandatoriCreateRequest struct {
	Usulan           string `json:"usulan"`
	Uraian           string `json:"uraian"`
	Tahun            string `json:"tahun"`
	RekinId          string `json:"rencana_kinerja_id"`
	PegawaiId        string `json:"pegawai_id"`
	KodeOpd          string `json:"kode_opd"`
	PeraturanTerkait string `json:"peraturan_terkait"`
	Status           string `json:"status"`
}
