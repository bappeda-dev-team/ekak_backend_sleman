package repository

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/model/domain"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, users domain.Users) (domain.Users, error)
	Update(ctx context.Context, tx *sql.Tx, users domain.Users) (domain.Users, error)
	FindAll(ctx context.Context, tx *sql.Tx, kodeOpd string) ([]domain.Users, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Users, error)
	FindByNip(ctx context.Context, tx *sql.Tx, nip string) (domain.Users, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindByEmailOrNip(ctx context.Context, tx *sql.Tx, username string) (domain.Users, error)
	FindByKodeOpdAndRole(ctx context.Context, tx *sql.Tx, kodeOpd string, roleName string) ([]domain.Users, error)
	CekAdminOpd(ctx context.Context, tx *sql.Tx) ([]domain.Users, error)
	CreateCaptcha(ctx context.Context, captcha domain.Captcha) error
	ValidateCaptcha(ctx context.Context, captchaID string, captchaValue string) (bool, error)
	DeleteCaptcha(ctx context.Context, captchaID string) error
	FindUserInfo(ctx context.Context, tx *sql.Tx, userId int) (domain.Users, error)
	UpdatePassword(ctx context.Context, tx *sql.Tx, users domain.Users) (domain.Users, error)

	// Rate limiting methods
	RecordFailedLogin(ctx context.Context, nip string) error
	CheckLoginAttempts(ctx context.Context, nip string) (isLocked bool, remainingTime int, err error)
	ResetLoginAttempts(ctx context.Context, nip string) error
}
