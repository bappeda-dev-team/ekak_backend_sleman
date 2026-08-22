package pohonkinerja

type CrosscuttingOpdUpdateRequest struct {
	Id         int                      `json:"id" validate:"required"`
	ParentId   int                      `json:"parent_id" validate:"required"`
	NamaPohon  string                   `json:"nama_pohon" validate:"required"`
	JenisPohon string                   `json:"jenis_pohon" validate:"required"`
	LevelPohon int                      `json:"level_pohon" validate:"required"`
	KodeOpd    string                   `json:"kode_opd" validate:"required"`
	Keterangan string                   `json:"keterangan"`
	Tahun      string                   `json:"tahun" validate:"required"`
	Status     string                   `json:"status" validate:"required"`
	Indikator  []IndikatorUpdateRequest `json:"indikator"`
}
