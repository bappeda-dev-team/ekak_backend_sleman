package usulan

type UsulanMusrebangUpdateRequest struct {
	Id      string `json:"id"`
	Usulan  string `json:"usulan"`
	Alamat  string `json:"alamat"`
	Uraian  string `json:"uraian"`
	Tahun   string `json:"tahun"`
	KodeOpd string `json:"kode_opd"`
	Status  string `json:"status"`
}
