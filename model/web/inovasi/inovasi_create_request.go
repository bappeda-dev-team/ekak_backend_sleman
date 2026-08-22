package inovasi

type InovasiCreateRequest struct {
	RekinId               string `json:"rencana_kinerja_id"`
	KodeOpd               string `json:"kode_opd"`
	JudulInovasi          string `json:"judul_inovasi"`
	JenisInovasi          string `json:"jenis_inovasi"`
	GambaranNilaiKebaruan string `json:"gambaran_nilai_kebaruan"`
}
