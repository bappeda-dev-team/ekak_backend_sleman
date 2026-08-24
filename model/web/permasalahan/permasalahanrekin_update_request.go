package permasalahan

type PermasalahanRekinUpdateRequest struct {
	Id                int    `json:"id"`
	Permasalahan      string `json:"permasalahan"`
	PenyebabInternal  string `json:"penyebab_internal"`
	PenyebabEksternal string `json:"penyebab_eksternal"`
}
