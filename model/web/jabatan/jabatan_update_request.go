package jabatan

type JabatanUpdateRequest struct {
	Id           string `json:"id"`
	KodeJabatan  string `json:"kode_jabatan"`
	NamaJabatan  string `json:"nama_jabatan"`
	KelasJabatan string `json:"kelas_jabatan"`
	JenisJabatan string `json:"jenis_jabatan"`
	NilaiJabatan int    `json:"nilai_jabatan"`
	KodeOpd      string `json:"kode_opd"`
	IndexJabatan int    `json:"index_jabatan"`
	Tahun        string `json:"tahun"`
	Esselon      string `json:"esselon"`
}
