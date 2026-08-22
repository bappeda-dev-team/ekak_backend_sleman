package domain

import "time"

type Captcha struct {
	ID        string
	Value     string
	ExpiresAt time.Time
}
