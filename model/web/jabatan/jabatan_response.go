package jabatan

import "ekak_kab_sleman/model/web/opdmaster"

type JabatanResponse struct {
	Id           string                      `json:"id"`
	KodeJabatan  string                      `json:"kode_jabatan"`
	NamaJabatan  string                      `json:"nama_jabatan"`
	KodeOpd      opdmaster.OpdResponseForAll `json:"operasional_daerah,omitempty"`
	KelasJabatan string                      `json:"kelas_jabatan,omitempty"`
	JenisJabatan string                      `json:"jenis_jabatan,omitempty"`
	NilaiJabatan int                         `json:"nilai_jabatan,omitempty"`
	IndexJabatan int                         `json:"index_jabatan,omitempty"`
	Tahun        string                      `json:"tahun,omitempty"`
	Esselon      string                      `json:"esselon,omitempty"`
}
