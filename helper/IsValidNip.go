package helper

import "regexp"

func IsValidNIP(nip string) bool {
	// Sesuaikan pattern regex dengan format NIP yang valid
	// Contoh: NIP harus 18 digit angka
	pattern := `^\d{16}$`
	match, _ := regexp.MatchString(pattern, nip)
	return match
}
