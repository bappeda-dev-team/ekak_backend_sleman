package usulan

type UsulanMusrebangCreateRequest struct {
	Usulan  string `json:"usulan"`
	Alamat  string `json:"alamat"`
	Uraian  string `json:"uraian"`
	Tahun   string `json:"tahun"`
	RekinId string `json:"rencana_kinerja_id"`
	KodeOpd string `json:"kode_opd"`
	Status  string `json:"status"`
}

type UsulanMusrebangCreateRekinRequest struct {
	IdUsulan string `json:"id_usulan" validate:"required"`
	RekinId  string `json:"rekin_id" validate:"required"`
}
