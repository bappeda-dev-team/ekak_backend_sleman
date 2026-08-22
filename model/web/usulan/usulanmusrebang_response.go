package usulan

type UsulanMusrebangResponse struct {
	Id        string `json:"id,omitempty"`
	Usulan    string `json:"usulan,omitempty"`
	Alamat    string `json:"alamat,omitempty"`
	Uraian    string `json:"uraian,omitempty"`
	Tahun     string `json:"tahun,omitempty"`
	RekinId   string `json:"rencana_kinerja_id,omitempty"`
	KodeOpd   string `json:"kode_opd,omitempty"`
	NamaOpd   string `json:"nama_opd,omitempty"`
	IsActive  bool   `json:"is_active,omitempty"`
	Status    string `json:"status,omitempty"`
	CreatedAt string `json:"dibuat_pada,omitempty" time_format:"2006-01-02 15:04:05"`
}
