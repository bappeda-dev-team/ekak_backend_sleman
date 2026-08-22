package helper

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomNumber(length int) int {
	// Inisialisasi seed random dengan waktu saat ini
	rand.Seed(time.Now().UnixNano())

	// Menentukan batas minimum dan maksimum
	min := int(pow10(length - 1))
	max := int(pow10(length) - 1)

	// Generate angka random
	randomNum := rand.Intn(max-min+1) + min

	return randomNum
}

// Fungsi helper untuk menghitung 10^n
func pow10(n int) int64 {
	result := int64(1)
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}

// Fungsi alternatif menggunakan string
func GenerateRandomNumberString(length int) string {
	// Inisialisasi seed random
	rand.Seed(time.Now().UnixNano())

	// Generate setiap digit
	digits := make([]byte, length)
	// Digit pertama tidak boleh 0
	digits[0] = byte(rand.Intn(9) + 1 + '0')

	// Generate digit selanjutnya
	for i := 1; i < length; i++ {
		digits[i] = byte(rand.Intn(10) + '0')
	}

	// Konversi ke string
	result, _ := strconv.Atoi(string(digits))
	return strconv.Itoa(result)
}
