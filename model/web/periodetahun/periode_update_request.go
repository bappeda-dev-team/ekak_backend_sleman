package periodetahun

type PeriodeUpdateRequest struct {
	Id           int    `json:"id"`
	TahunAwal    string `json:"tahun_awal"`
	TahunAkhir   string `json:"tahun_akhir"`
	JenisPeriode string `json:"jenis_periode"`
}
