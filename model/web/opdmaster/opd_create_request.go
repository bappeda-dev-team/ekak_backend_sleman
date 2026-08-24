package opdmaster

type OpdCreateRequest struct {
	KodeOpd       string `json:"kode_opd" validate:"required"`
	NamaOpd       string `json:"nama_opd" validate:"required"`
	Singkatan     string `json:"singkatan"`
	Alamat        string `json:"alamat"`
	Telepon       string `json:"telepon"`
	Fax           string `json:"fax"`
	Email         string `json:"email"`
	Website       string `json:"website"`
	NamaKepalaOpd string `json:"nama_kepala_opd" validate:"required"`
	NIPKepalaOpd  string `json:"nip_kepala_opd" validate:"required"`
	PangkatKepala string `json:"pangkat_kepala" validate:"required"`
	IdLembaga     string `json:"id_lembaga" validate:"required"`
}
