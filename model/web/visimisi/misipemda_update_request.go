package visimisipemda

type MisiPemdaUpdateRequest struct {
	Id                int    `json:"id"`
	IdVisi            int    `json:"id_visi"`
	Misi              string `json:"misi"`
	Urutan            int    `json:"urutan"`
	TahunAwalPeriode  string `json:"tahun_awal_periode"`
	TahunAkhirPeriode string `json:"tahun_akhir_periode"`
	JenisPeriode      string `json:"jenis_periode"`
	Keterangan        string `json:"keterangan"`
}
