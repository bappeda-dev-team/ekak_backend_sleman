package inovasi

type InovasiUpdateRequest struct {
	Id                    string `json:"id"`
	RekinId               string `json:"rencana_kinerja_id"`
	JudulInovasi          string `json:"judul_inovasi"`
	JenisInovasi          string `json:"jenis_inovasi"`
	GambaranNilaiKebaruan string `json:"gambaran_nilai_kebaruan"`
}
