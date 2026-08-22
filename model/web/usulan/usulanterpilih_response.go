package usulan

type UsulanTerpilihResponse struct {
	Id          string `json:"id"`
	JenisUsulan string `json:"jenis_usulan"`
	UsulanId    string `json:"usulan_id"`
	RekinId     string `json:"rekin_id"`
	Tahun       string `json:"tahun"`
	KodeOpd     string `json:"kode_opd"`
	Keterangan  string `json:"keterangan"`
}
