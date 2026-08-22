package domain

type Periode struct {
	Id           int
	TahunAwal    string
	TahunAkhir   string
	JenisPeriode string
}

type TahunPeriode struct {
	Id        int
	IdPeriode int
	Tahun     string
}
