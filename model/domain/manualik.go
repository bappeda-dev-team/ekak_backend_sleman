package domain

type ManualIK struct {
	Id                  int
	IndikatorId         string
	Perspektif          string
	TujuanRekin         string
	Definisi            string
	KeyActivities       string
	Formula             string
	JenisIndikator      string
	Kinerja             bool
	Penduduk            bool
	Spatial             bool
	UnitPenanggungJawab string
	UnitPenyediaData    string
	SumberData          string
	JangkaWaktuAwal     string
	JangkaWaktuAkhir    string
	PeriodePelaporan    string
	DataIndikator       Indikator `json:"data_indikator"`
	DataTarget          []Target  `json:"data_target"`
}
