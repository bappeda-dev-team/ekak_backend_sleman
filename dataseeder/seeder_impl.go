package dataseeder

import (
	"context"
	"database/sql"

	"ekak_kab_sleman/helper"
	"log"
)

type SeederImpl struct {
	DB            *sql.DB
	RoleSeeder    RoleSeeder
	UserSeeder    UserSeeder
	PegawaiSeeder PegawaiSeeder
}

func NewSeederImpl(db *sql.DB, roleSeeder RoleSeeder, userSeeder UserSeeder, pegawaiSeeder PegawaiSeeder) *SeederImpl {
	return &SeederImpl{
		DB:            db,
		RoleSeeder:    roleSeeder,
		UserSeeder:    userSeeder,
		PegawaiSeeder: pegawaiSeeder,
	}
}

func (seeder *SeederImpl) SeedAll() {
	tx, err := seeder.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer helper.CommitOrRollback(tx)

	ctx := context.Background()

	// Seed roles dulu
	err = seeder.RoleSeeder.Seed(ctx, tx)
	if err != nil {
		log.Fatal(err)
	}

	// Kemudian seed users
	err = seeder.UserSeeder.Seed(ctx, tx)
	if err != nil {
		log.Fatal(err)
	}

	// Seed pegawai
	err = seeder.PegawaiSeeder.Seed(ctx, tx)
	if err != nil {
		log.Fatal(err)
	}
}
