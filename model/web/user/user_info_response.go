package user

import "time"

type UserInfoResponse struct {
	Id                     int            `json:"id,omitempty"`
	Nip                    string         `json:"nip"`
	Email                  string         `json:"email,omitempty"`
	IsActive               bool           `json:"is_active"`
	PasswordUpdatedAt      *time.Time     `json:"password_udpated_at"`
	PasswordChangeRequired bool           `json:"password_change_required"`
}
