package pegawai

type PegawaiUpdateRequest struct {
	Id          string `json:"id"`
	NamaPegawai string `json:"nama_pegawai"`
	Nip         string `json:"nip"`
	KodeOpd     string `json:"kode_opd"`
}
