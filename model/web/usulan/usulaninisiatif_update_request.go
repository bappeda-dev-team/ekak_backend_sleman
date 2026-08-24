package usulan

type UsulanInisiatifUpdateRequest struct {
	Id        string `json:"id"`
	Usulan    string `json:"usulan"`
	Manfaat   string `json:"manfaat"`
	Uraian    string `json:"uraian"`
	Tahun     string `json:"tahun"`
	PegawaiId string `json:"pegawai_id"`
	KodeOpd   string `json:"kode_opd"`
	Status    string `json:"status"`
}
