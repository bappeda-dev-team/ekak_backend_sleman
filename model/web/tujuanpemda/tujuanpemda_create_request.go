package tujuanpemda

type TujuanPemdaCreateRequest struct {
	IdVisi            int                      `json:"id_visi"`
	IdMisi            int                      `json:"id_misi"`
	TujuanPemda       string                   `json:"tujuan_pemda"`
	TematikId         int                      `json:"tema_id"`
	PeriodeId         int                      `json:"periode_id"`
	TahunAwalPeriode  string                   `json:"tahun_awal_periode"`
	TahunAkhirPeriode string                   `json:"tahun_akhir_periode"`
	JenisPeriode      string                   `json:"jenis_periode"`
	Indikator         []IndikatorCreateRequest `json:"indikator"`
}

type IndikatorCreateRequest struct {
	Indikator        string                `json:"indikator"`
	RumusPerhitungan string                `json:"rumus_perhitungan"`
	SumberData       string                `json:"sumber_data"`
	Target           []TargetCreateRequest `json:"target"`
}

type TargetCreateRequest struct {
	Target string `json:"target"`
	Satuan string `json:"satuan"`
	Tahun  string `json:"tahun"`
}
