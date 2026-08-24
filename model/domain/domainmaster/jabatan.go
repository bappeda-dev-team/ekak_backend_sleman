package domainmaster

type Jabatan struct {
	Id           string
	KodeJabatan  string
	NamaJabatan  string
	KelasJabatan string
	JenisJabatan string
	NilaiJabatan int //default 0
	KodeOpd      string
	NamaOpd      string
	IndexJabatan int //default 0
	Tahun        string
	Esselon      string
}
