package repository

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/domain"
	"log"
	"sort"
	"sync"
	"time"
)

type UserRepositoryImpl struct {
	loginAttempts sync.Map
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	repo := &UserRepositoryImpl{}

	// Background cleaner untuk menghapus login attempts yang sudah expired
	go repo.cleanupExpiredAttempts()

	return repo
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, users domain.Users) (domain.Users, error) {
	script := "INSERT INTO tb_users(nip, email, password, is_active) VALUES (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, script, users.Nip, users.Email, users.Password, users.IsActive)
	if err != nil {
		return users, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return users, err
	}
	users.Id = int(id)

	scriptRole := "INSERT INTO tb_user_role(user_id, role_id) VALUES (?, ?)"
	for _, role := range users.Role {
		_, err = tx.ExecContext(ctx, scriptRole, users.Id, role.Id)
		if err != nil {
			return users, err
		}
	}

	return users, nil
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, users domain.Users) (domain.Users, error) {
	script := "UPDATE tb_users SET nip = ?, email = ?, password = ?, is_active = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, users.Nip, users.Email, users.Password, users.IsActive, users.Id)
	if err != nil {
		return users, err
	}

	scriptDeleteRoles := "DELETE FROM tb_user_role WHERE user_id = ?"
	_, err = tx.ExecContext(ctx, scriptDeleteRoles, users.Id)
	if err != nil {
		return users, err
	}

	scriptRole := "INSERT INTO tb_user_role(user_id, role_id) VALUES (?, ?)"
	for _, role := range users.Role {
		_, err = tx.ExecContext(ctx, scriptRole, users.Id, role.Id)
		if err != nil {
			return users, err
		}
	}

	return repository.FindById(ctx, tx, users.Id)
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, kodeOpd string) ([]domain.Users, error) {
	script := `
        SELECT DISTINCT u.id, u.nip, u.email, u.is_active, ur.role_id, r.role 
        FROM tb_users u
        LEFT JOIN tb_user_role ur ON u.id = ur.user_id
        LEFT JOIN tb_role r ON ur.role_id = r.id
        INNER JOIN tb_pegawai p ON u.nip = p.nip
        WHERE 1=1
    `
	var params []any

	if kodeOpd != "" {
		script += " AND p.kode_opd = ?"
		params = append(params, kodeOpd)
	}

	script += " ORDER BY u.id, ur.role_id"

	rows, err := tx.QueryContext(ctx, script, params...)
	if err != nil {
		return []domain.Users{}, err
	}
	defer rows.Close()

	var users []domain.Users
	userMap := make(map[int]*domain.Users)

	for rows.Next() {
		var userId int
		var nip, email string
		var isActive bool
		var roleId sql.NullInt64
		var roleName sql.NullString

		err := rows.Scan(
			&userId,
			&nip,
			&email,
			&isActive,
			&roleId,
			&roleName,
		)
		if err != nil {
			return []domain.Users{}, err
		}

		user, exists := userMap[userId]
		if !exists {
			user = &domain.Users{
				Id:       userId,
				Nip:      nip,
				Email:    email,
				IsActive: isActive,
				Role:     []domain.Roles{},
			}
			userMap[userId] = user
		}

		if roleId.Valid && roleName.Valid {
			user.Role = append(user.Role, domain.Roles{
				Id:   int(roleId.Int64),
				Role: roleName.String,
			})
		}
	}

	for _, user := range userMap {
		users = append(users, *user)
	}

	return users, nil
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Users, error) {
	script := `
		SELECT u.id, u.nip, u.email, u.password, u.is_active, ur.role_id, r.role 
		FROM tb_users u
			LEFT JOIN tb_user_role ur ON u.id = ur.user_id
			LEFT JOIN tb_role r ON ur.role_id = r.id
		WHERE u.id = ?
		ORDER BY ur.role_id
	`
	rows, err := tx.QueryContext(ctx, script, id)
	if err != nil {
		return domain.Users{}, err
	}
	defer rows.Close()

	var user domain.Users
	first := true

	for rows.Next() {
		var roleId sql.NullInt64
		var roleName sql.NullString

		if first {
			err := rows.Scan(
				&user.Id,
				&user.Nip,
				&user.Email,
				&user.Password,
				&user.IsActive,
				&roleId,
				&roleName,
			)
			if err != nil {
				return domain.Users{}, err
			}
			first = false
		} else {
			var userId int
			var nip, email, password string
			var isActive bool
			err := rows.Scan(
				&userId,
				&nip,
				&email,
				&password,
				&isActive,
				&roleId,
				&roleName,
			)
			if err != nil {
				return domain.Users{}, err
			}
		}

		if roleId.Valid && roleName.Valid {
			user.Role = append(user.Role, domain.Roles{
				Id:   int(roleId.Int64),
				Role: roleName.String,
			})
		}
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindByNip(ctx context.Context, tx *sql.Tx, nip string) (domain.Users, error) {
	script := `
		SELECT u.id, u.nip, u.email, u.password, u.is_active, ur.role_id, r.role 
		FROM tb_users u
			LEFT JOIN tb_user_role ur ON u.id = ur.user_id
			LEFT JOIN tb_role r ON ur.role_id = r.id
		WHERE u.nip = ?
		ORDER BY ur.role_id
	`
	rows, err := tx.QueryContext(ctx, script, nip)
	if err != nil {
		return domain.Users{}, err
	}
	defer rows.Close()

	var user domain.Users
	first := true

	for rows.Next() {
		var roleId sql.NullInt64
		var roleName sql.NullString

		if first {
			err := rows.Scan(
				&user.Id,
				&user.Nip,
				&user.Email,
				&user.Password,
				&user.IsActive,
				&roleId,
				&roleName,
			)
			if err != nil {
				return domain.Users{}, err
			}
			first = false
		} else {
			var userId int
			var nip, email, password string
			var isActive bool
			err := rows.Scan(
				&userId,
				&nip,
				&email,
				&password,
				&isActive,
				&roleId,
				&roleName,
			)
			if err != nil {
				return domain.Users{}, err
			}
		}

		if roleId.Valid && roleName.Valid {
			user.Role = append(user.Role, domain.Roles{
				Id:   int(roleId.Int64),
				Role: roleName.String,
			})
		}
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	scriptRole := "DELETE FROM tb_user_role WHERE user_id = ?"
	_, err := tx.ExecContext(ctx, scriptRole, id)
	if err != nil {
		return err
	}

	scriptUser := "DELETE FROM tb_users WHERE id = ?"
	_, err = tx.ExecContext(ctx, scriptUser, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *UserRepositoryImpl) FindByEmailOrNip(ctx context.Context, tx *sql.Tx, username string) (domain.Users, error) {
	startTime := time.Now()
	log.Printf("Start finding user by NIP: %s", username)

	// Query untuk mendapatkan data user terlebih dahulu
	userScript := `
        SELECT 
            id, 
            nip, 
            email, 
            password, 
            is_active
        FROM tb_users
        WHERE nip = ?
    `

	var user domain.Users
	err := tx.QueryRowContext(ctx, userScript, username).Scan(
		&user.Id,
		&user.Nip,
		&user.Email,
		&user.Password,
		&user.IsActive,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No user found with NIP: %s", username)
			return domain.Users{}, err
		}
		log.Printf("Error querying user by NIP: %v", err)
		return domain.Users{}, err
	}

	// Query terpisah untuk mendapatkan role
	roleScript := `
        SELECT DISTINCT 
            r.id,
            r.role
        FROM tb_role r
        JOIN tb_user_role ur ON r.id = ur.role_id
        WHERE ur.user_id = ?
    `

	roleRows, err := tx.QueryContext(ctx, roleScript, user.Id)
	if err != nil {
		log.Printf("Error querying roles: %v", err)
		return domain.Users{}, err
	}
	defer roleRows.Close()

	user.Role = []domain.Roles{} // Inisialisasi slice kosong

	// Scan semua role yang dimiliki user
	for roleRows.Next() {
		var role domain.Roles
		err := roleRows.Scan(
			&role.Id,
			&role.Role,
		)
		if err != nil {
			log.Printf("Error scanning role: %v", err)
			return domain.Users{}, err
		}
		user.Role = append(user.Role, role)
	}

	if err = roleRows.Err(); err != nil {
		log.Printf("Error iterating roles: %v", err)
		return domain.Users{}, err
	}

	log.Printf("Successfully found user by NIP: %s with %d roles, execution time: %v",
		username, len(user.Role), time.Since(startTime))
	return user, nil
}
func (repository *UserRepositoryImpl) FindByKodeOpdAndRole(ctx context.Context, tx *sql.Tx, kodeOpd string, roleName string) ([]domain.Users, error) {
	script := `
        SELECT DISTINCT 
            u.id, 
            u.nip, 
            u.email, 
            u.is_active, 
            ur.role_id, 
            r.role, 
            p.id as pegawai_id,
            p.nama as nama_pegawai  
        FROM tb_users u
        LEFT JOIN tb_user_role ur ON u.id = ur.user_id
        LEFT JOIN tb_role r ON ur.role_id = r.id
        INNER JOIN tb_pegawai p ON u.nip = p.nip
        WHERE p.kode_opd = ? AND r.role = ?
        ORDER BY u.id, ur.role_id
    `

	rows, err := tx.QueryContext(ctx, script, kodeOpd, roleName)
	if err != nil {
		return []domain.Users{}, err
	}
	defer rows.Close()

	userMap := make(map[int]*domain.Users)

	for rows.Next() {
		var userId int
		var nip, email string
		var isActive bool
		var roleId sql.NullInt64
		var roleName sql.NullString
		var pegawaiId string
		var namaPegawai string

		err := rows.Scan(
			&userId,
			&nip,
			&email,
			&isActive,
			&roleId,
			&roleName,
			&pegawaiId,
			&namaPegawai,
		)
		if err != nil {
			return []domain.Users{}, err
		}

		user, exists := userMap[userId]
		if !exists {
			user = &domain.Users{
				Id:          userId,
				Nip:         nip,
				Email:       email,
				IsActive:    isActive,
				PegawaiId:   pegawaiId,
				NamaPegawai: namaPegawai,
				Role:        []domain.Roles{},
			}
			userMap[userId] = user
		}

		if roleId.Valid && roleName.Valid {
			user.Role = append(user.Role, domain.Roles{
				Id:   int(roleId.Int64),
				Role: roleName.String,
			})
		}
	}

	// Sort users berdasarkan ID untuk konsistensi
	var sortedUsers []domain.Users
	for _, user := range userMap {
		sortedUsers = append(sortedUsers, *user)
	}
	sort.Slice(sortedUsers, func(i, j int) bool {
		return sortedUsers[i].Id < sortedUsers[j].Id
	})

	return sortedUsers, nil
}

func (repository *UserRepositoryImpl) CekAdminOpd(ctx context.Context, tx *sql.Tx) ([]domain.Users, error) {
	script := `
		SELECT 
			u.id as user_id,
			u.nip,
			u.email,
			u.is_active,
			p.id as pegawai_id,
			p.nama as nama_pegawai,
			p.kode_opd,
			o.nama_opd,
			r.id as role_id,
			r.role
		FROM tb_users u
		INNER JOIN tb_user_role ur ON u.id = ur.user_id
		INNER JOIN tb_role r ON ur.role_id = r.id
		INNER JOIN tb_pegawai p ON u.nip = p.nip
		INNER JOIN tb_operasional_daerah o ON p.kode_opd = o.kode_opd
		WHERE r.role = 'admin_opd'
		ORDER BY o.kode_opd, u.id
	`

	rows, err := tx.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.Users
	for rows.Next() {
		var user domain.Users
		var role domain.Roles

		err := rows.Scan(
			&user.Id,
			&user.Nip,
			&user.Email,
			&user.IsActive,
			&user.PegawaiId,
			&user.NamaPegawai,
			&user.KodeOpd,
			&user.NamaOpd,
			&role.Id,
			&role.Role,
		)
		if err != nil {
			return nil, err
		}

		user.Role = []domain.Roles{role}
		users = append(users, user)
	}

	return users, nil
}

// In-memory storage untuk captcha
var (
	captchaStore = make(map[string]domain.Captcha)
	captchaMutex sync.RWMutex
)

// CreateCaptcha menyimpan captcha baru ke dalam memory
func (repository *UserRepositoryImpl) CreateCaptcha(ctx context.Context, captcha domain.Captcha) error {
	captchaMutex.Lock()
	defer captchaMutex.Unlock()

	captchaStore[captcha.ID] = captcha
	return nil
}

// ValidateCaptcha memvalidasi captcha berdasarkan ID dan value
func (repository *UserRepositoryImpl) ValidateCaptcha(ctx context.Context, captchaID string, captchaValue string) (bool, error) {
	captchaMutex.RLock()
	defer captchaMutex.RUnlock()

	captcha, exists := captchaStore[captchaID]
	if !exists {
		return false, nil
	}

	// Cek apakah captcha sudah expired
	if time.Now().After(captcha.ExpiresAt) {
		// Hapus captcha yang expired
		captchaMutex.RUnlock()
		captchaMutex.Lock()
		delete(captchaStore, captchaID)
		captchaMutex.Unlock()
		captchaMutex.RLock()
		return false, nil
	}

	// Validasi value (case-insensitive)
	return helper.ValidateCaptchaImage(captchaID, captchaValue), nil
}

// DeleteCaptcha menghapus captcha dari memory
func (repository *UserRepositoryImpl) DeleteCaptcha(ctx context.Context, captchaID string) error {
	captchaMutex.Lock()
	defer captchaMutex.Unlock()

	delete(captchaStore, captchaID)
	return nil
}

// CleanupExpiredCaptcha membersihkan captcha yang sudah expired (bisa dipanggil secara berkala)
func (repository *UserRepositoryImpl) CleanupExpiredCaptcha(ctx context.Context) {
	captchaMutex.Lock()
	defer captchaMutex.Unlock()

	now := time.Now()
	for id, captcha := range captchaStore {
		if now.After(captcha.ExpiresAt) {
			delete(captchaStore, id)
		}
	}
}

// RecordFailedLogin mencatat percobaan login yang gagal
func (repository *UserRepositoryImpl) RecordFailedLogin(ctx context.Context, nip string) error {
	now := time.Now()

	// Load existing attempt atau buat baru
	var attempt domain.LoginAttempt
	if val, ok := repository.loginAttempts.Load(nip); ok {
		attempt = val.(domain.LoginAttempt)
	} else {
		attempt = domain.LoginAttempt{
			Nip:         nip,
			FailedCount: 0,
		}
	}

	// Increment failed count
	attempt.FailedCount++
	attempt.LastAttempt = now

	// Jika sudah 3 kali gagal, lock selama 3 menit
	if attempt.FailedCount >= 3 {
		attempt.LockedUntil = now.Add(3 * time.Minute)
		log.Printf("[RATE LIMIT] NIP %s di-lock sampai %s", nip, attempt.LockedUntil.Format("15:04:05"))
	}

	// Simpan kembali
	repository.loginAttempts.Store(nip, attempt)

	return nil
}

// CheckLoginAttempts mengecek apakah user sedang di-lock
func (repository *UserRepositoryImpl) CheckLoginAttempts(ctx context.Context, nip string) (isLocked bool, remainingTime int, err error) {
	val, ok := repository.loginAttempts.Load(nip)
	if !ok {
		// Belum ada attempt
		return false, 0, nil
	}

	attempt := val.(domain.LoginAttempt)
	now := time.Now()

	// Cek apakah masih dalam periode lock
	if attempt.FailedCount >= 3 && now.Before(attempt.LockedUntil) {
		remainingSeconds := int(attempt.LockedUntil.Sub(now).Seconds())
		log.Printf("[RATE LIMIT] NIP %s masih di-lock, sisa %d detik", nip, remainingSeconds)
		return true, remainingSeconds, nil
	}

	// Jika sudah lewat waktu lock, reset otomatis
	if attempt.FailedCount >= 3 && now.After(attempt.LockedUntil) {
		repository.loginAttempts.Delete(nip)
		log.Printf("[RATE LIMIT] NIP %s lock sudah expired, direset", nip)
		return false, 0, nil
	}

	return false, 0, nil
}

// ResetLoginAttempts mereset percobaan login (dipanggil saat login berhasil)
func (repository *UserRepositoryImpl) ResetLoginAttempts(ctx context.Context, nip string) error {
	repository.loginAttempts.Delete(nip)
	log.Printf("[RATE LIMIT] NIP %s login berhasil, attempt direset", nip)
	return nil
}

// cleanupExpiredAttempts membersihkan login attempts yang sudah expired setiap 5 menit
func (repository *UserRepositoryImpl) cleanupExpiredAttempts() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()
		repository.loginAttempts.Range(func(key, value any) bool {
			attempt := value.(domain.LoginAttempt)

			// Hapus jika sudah lebih dari 10 menit tidak ada aktivitas
			if now.Sub(attempt.LastAttempt) > 10*time.Minute {
				repository.loginAttempts.Delete(key)
				log.Printf("[RATE LIMIT CLEANUP] Menghapus attempt untuk NIP %s", key)
			}

			return true
		})
	}
}

func (repository *UserRepositoryImpl) FindUserInfo(
	ctx context.Context,
	tx *sql.Tx,
	userId int,
) (domain.Users, error) {
	script := `
	SELECT
	   u.id,
	   u.nip,
	   u.email,
	   u.is_active,
	   u.password_updated_at
	FROM tb_users u
	WHERE u.id = ?
	`

	var user domain.Users

	err := tx.QueryRowContext(ctx, script, userId).Scan(
		&user.Id,
		&user.Nip,
		&user.Email,
		&user.IsActive,
		&user.PasswordUpdatedAt,
	)
	if err != nil {
		return domain.Users{}, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) UpdatePassword(ctx context.Context, tx *sql.Tx, users domain.Users) (domain.Users, error) {
	script := `
        UPDATE tb_users
        SET password = ?,
            updated_by = ?,
            updated_at = NOW(),
	    password_updated_at = NOW()
        WHERE id = ?
    `
	_, err := tx.ExecContext(ctx, script, users.Password, users.UpdatedBy, users.Id)
	if err != nil {
		return users, err
	}

	return repository.FindById(ctx, tx, users.Id)
}
