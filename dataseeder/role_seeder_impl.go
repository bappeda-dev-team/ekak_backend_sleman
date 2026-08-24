package dataseeder

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/model/domain"
	"ekak_kab_sleman/repository"
	"log"
)

type RoleSeederImpl struct {
	RoleRepository repository.RoleRepository
}

func NewRoleSeederImpl(roleRepository repository.RoleRepository) *RoleSeederImpl {
	return &RoleSeederImpl{
		RoleRepository: roleRepository,
	}
}

func (seeder *RoleSeederImpl) Seed(ctx context.Context, tx *sql.Tx) error {
	roles, err := seeder.RoleRepository.FindAll(ctx, tx)
	if err != nil {
		return err
	}

	if len(roles) > 0 {
		log.Println("Roles sudah ada, skip seeding roles")
		return nil
	}

	defaultRoles := []domain.Roles{
		{
			Role: "super_admin",
		},
		{
			Role: "admin_opd",
		},
		{
			Role: "admin_kecamatan",
		},
		{
			Role: "eselon_1",
		},
		{
			Role: "eselon_2",
		},
		{
			Role: "eselon_3",
		},
		{
			Role: "eselon_4",
		},
		{
			Role: "staff",
		},
		{
			Role: "reviewer",
		},
	}

	for _, role := range defaultRoles {
		_, err := seeder.RoleRepository.Create(ctx, tx, role)
		if err != nil {
			return err
		}
	}
	log.Println("Roles berhasil di-seed")
	return nil
}
