package helper

import (
	"bytes"
	"encoding/base64"
	"time"

	"github.com/dchest/captcha"
)

// Inisialisasi captcha store (gunakan memory store)
var captchaStore = captcha.NewMemoryStore(captcha.CollectNum, 5*time.Minute)

// InitCaptcha menginisialisasi captcha store
func InitCaptcha() {
	captcha.SetCustomStore(captchaStore)
}

// GenerateCaptchaImage menghasilkan captcha image dan mengembalikan ID, base64 image, dan value
func GenerateCaptchaImage() (string, string, string) {
	// Generate captcha ID (ini akan menjadi ID unik untuk captcha)
	captchaID := captcha.New()

	// Generate captcha value (untuk validasi)
	captchaValue := captcha.RandomDigits(6) // 6 digit angka

	// Set value ke store
	captchaStore.Set(captchaID, captchaValue)

	// Generate image sebagai base64
	var buf bytes.Buffer
	err := captcha.WriteImage(&buf, captchaID, 240, 80) // width: 240, height: 80
	if err != nil {
		return "", "", ""
	}

	// Convert ke base64
	imgBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())

	// Convert value ke string untuk response (opsional, untuk testing)
	valueStr := ""
	for _, d := range captchaValue {
		valueStr += string('0' + d)
	}

	return captchaID, imgBase64, valueStr
}

// ValidateCaptchaImage memvalidasi captcha berdasarkan ID dan value
func ValidateCaptchaImage(captchaID string, captchaValue string) bool {
	return captcha.VerifyString(captchaID, captchaValue)
}

// GetCaptchaExpirationTime mengembalikan waktu expiration (5 menit dari sekarang)
func GetCaptchaExpirationTime() time.Time {
	return time.Now().Add(5 * time.Minute)
}
