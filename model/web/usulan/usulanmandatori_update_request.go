package usulan

type UsulanMandatoriUpdateRequest struct {
	Id               string `json:"id"`
	Usulan           string `json:"usulan"`
	PeraturanTerkait string `json:"peraturan_terkait"`
	Uraian           string `json:"uraian"`
	Tahun            string `json:"tahun"`
	PegawaiId        string `json:"pegawai_id"`
	KodeOpd          string `json:"kode_opd"`
	Status           string `json:"status"`
}
