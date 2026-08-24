package web

type ActionButton struct {
	NameAction  string `json:"name"`
	Method      string `json:"method"`
	Url         string `json:"url"`
	JenisUsulan string `json:"jenis_usulan,omitempty"`
}
