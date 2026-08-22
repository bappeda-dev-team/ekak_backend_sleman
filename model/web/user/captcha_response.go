package user

type CaptchaResponse struct {
	CaptchaID    string `json:"captcha_id"`
	CaptchaImage string `json:"captcha_image"`
	CaptchaValue string `json:"captcha_value,omitempty"`
}
