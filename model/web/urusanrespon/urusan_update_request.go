package urusanrespon

type UrusanUpdateRequest struct {
	Id         string `json:"id"`
	KodeUrusan string `json:"kode_urusan"`
	NamaUrusan string `json:"nama_urusan"`
}
