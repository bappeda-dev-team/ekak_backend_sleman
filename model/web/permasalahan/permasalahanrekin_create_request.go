package permasalahan

type PermasalahanRekinCreateRequest struct {
	Id                int    `json:"id"`
	RekinId           string `json:"rekin_id"`
	Permasalahan      string `json:"permasalahan"`
	PenyebabInternal  string `json:"penyebab_internal"`
	PenyebabEksternal string `json:"penyebab_eksternal"`
	JenisPermasalahan string `json:"jenis_permasalahan"`
}
