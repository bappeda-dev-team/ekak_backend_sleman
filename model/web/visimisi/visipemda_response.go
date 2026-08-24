package visimisipemda

type VisiPemdaResponse struct {
	Id                int    `json:"id"`
	Visi              string `json:"visi"`
	TahunAwalPeriode  string `json:"tahun_awal_periode"`
	TahunAkhirPeriode string `json:"tahun_akhir_periode"`
	JenisPeriode      string `json:"jenis_periode"`
	Keterangan        string `json:"keterangan"`
}
