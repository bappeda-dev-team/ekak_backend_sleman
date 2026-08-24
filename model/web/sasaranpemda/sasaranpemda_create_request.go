package sasaranpemda

type SasaranPemdaCreateRequest struct {
	SubtemaId     int                      `json:"subtema_id"`
	TujuanPemdaId int                      `json:"tujuan_pemda_id"`
	SasaranPemda  string                   `json:"sasaran_pemda"`
	PeriodeId     int                      `json:"periode_id"`
	TahunAwal     string                   `json:"tahun_awal"`
	TahunAkhir    string                   `json:"tahun_akhir"`
	JenisPeriode  string                   `json:"jenis_periode"`
	Indikator     []IndikatorCreateRequest `json:"indikator"`
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
