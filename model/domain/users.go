package domain

import "time"

type Users struct {
	Id                int
	Nip               string
	Email             string
	Password          string
	IsActive          bool
	Role              []Roles
	PegawaiId         string
	NamaPegawai       string
	KodeOpd           string
	NamaOpd           string
	PasswordUpdatedAt *time.Time
	UpdatedBy         int
}

type LoginAttempt struct {
	Nip         string
	FailedCount int
	LockedUntil time.Time
	LastAttempt time.Time
}
