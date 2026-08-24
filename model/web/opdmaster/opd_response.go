package opdmaster

import "ekak_kab_sleman/model/web/lembaga"

type OpdResponse struct {
	Id            string                  `json:"id"`
	KodeOpd       string                  `json:"kode_opd"`
	NamaOpd       string                  `json:"nama_opd"`
	Singkatan     string                  `json:"singkatan"`
	Alamat        string                  `json:"alamat"`
	Telepon       string                  `json:"telepon"`
	Fax           string                  `json:"fax"`
	Email         string                  `json:"email"`
	Website       string                  `json:"website"`
	NamaKepalaOpd string                  `json:"nama_kepala_opd"`
	NIPKepalaOpd  string                  `json:"nip_kepala_opd"`
	PangkatKepala string                  `json:"pangkat_kepala"`
	IdLembaga     lembaga.LembagaResponse `json:"id_lembaga"`
}

type OpdResponseForAll struct {
	KodeOpd string `json:"kode_opd,omitempty"`
	NamaOpd string `json:"nama_opd,omitempty"`
}
