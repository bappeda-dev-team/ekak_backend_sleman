package usulan

type UsulanPokokPikiranCreateRequest struct {
	Usulan    string `json:"usulan"`
	Alamat    string `json:"alamat"`
	Uraian    string `json:"uraian"`
	Tahun     string `json:"tahun"`
	RekinId   string `json:"rencana_kinerja_id"`
	PegawaiId string `json:"pegawai_id"`
	KodeOpd   string `json:"kode_opd"`
	Status    string `json:"status"`
}

type UsulanPokokPikiranCreateRekinRequest struct {
	IdUsulan string `json:"id_usulan" validate:"required"`
	RekinId  string `json:"rekin_id" validate:"required"`
}
