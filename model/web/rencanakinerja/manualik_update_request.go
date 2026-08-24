package rencanakinerja

type ManualIKUpdateRequest struct {
	Id                  int                     `json:"id,omitempty"`
	IdIndikator         string                  `json:"id_indikator"`
	Perspektif          string                  `json:"perspektif" validate:"required"`
	TujuanRekin         string                  `json:"tujuan_rekin"`
	Definisi            string                  `json:"definisi"`
	KeyActivities       string                  `json:"key_activities"`
	Formula             string                  `json:"formula"`
	JenisIndikator      string                  `json:"jenis_indikator"`
	OutputData          OutputDataUpdateRequest `json:"output_data"`
	UnitPenanggungJawab string                  `json:"unit_penanggung_jawab"`
	UnitPenyediaData    string                  `json:"unit_penyedia_data"`
	SumberData          string                  `json:"sumber_data"`
	JangkaWaktuAwal     string                  `json:"jangka_waktu_awal"`
	JangkaWaktuAkhir    string                  `json:"jangka_waktu_akhir"`
	PeriodePelaporan    string                  `json:"periode_pelaporan"`
}

type OutputDataUpdateRequest struct {
	Kinerja  bool `json:"kinerja"`
	Penduduk bool `json:"penduduk"`
	Spatial  bool `json:"spatial"`
}
