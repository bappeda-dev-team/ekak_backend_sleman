package user

type UserLoginResponse struct {
	Token           string `json:"token,omitempty"`
	IsLocked        bool   `json:"is_locked,omitempty"`
	RemainingTime   int    `json:"remaining_time,omitempty"`
	RemainingMinute int    `json:"remaining_minute,omitempty"`
	RemainingSecond int    `json:"remaining_second,omitempty"`
	Message         string `json:"message,omitempty"`
}
