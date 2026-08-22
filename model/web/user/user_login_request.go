package user

type UserLoginRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	CaptchaID    string `json:"captcha_key"`
	CaptchaValue string `json:"captcha_answer"`
}
