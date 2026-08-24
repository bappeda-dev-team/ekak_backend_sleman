package gambaranumum

type GambaranUmumUpdateRequest struct {
	Id           string `json:"id"`
	Urutan       int    `json:"urutan"`
	GambaranUmum string `json:"gambaran_umum"`
}
