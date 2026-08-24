package dataseeder

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/model/domain"
	"ekak_kab_sleman/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserSeederImpl struct {
	UserRepository repository.UserRepository
	RoleRepository repository.RoleRepository
}

func NewUserSeederImpl(userRepository repository.UserRepository, roleRepository repository.RoleRepository) *UserSeederImpl {
	return &UserSeederImpl{
		UserRepository: userRepository,
		RoleRepository: roleRepository,
	}
}

func (seeder *UserSeederImpl) Seed(ctx context.Context, tx *sql.Tx) error {
	users, err := seeder.UserRepository.FindAll(ctx, tx, "")
	if err != nil {
		return err
	}

	if len(users) > 0 {
		log.Println("Users sudah ada, skip seeding users")
		return nil
	}

	roles, err := seeder.RoleRepository.FindAll(ctx, tx)
	if err != nil {
		return err
	}

	roleMap := make(map[string]domain.Roles)
	for _, role := range roles {
		roleMap[role.Role] = role
	}

	defaultUsers := []struct {
		user     domain.Users
		roleKeys []string
	}{
		{
			user: domain.Users{
				Nip:      "admin1",
				Email:    "admin@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"super_admin"},
		},
		{
			user: domain.Users{
				Nip:      "admin2",
				Email:    "admin2@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"super_admin"},
		},
		{
			user: domain.Users{
				Nip:      "admin_dindik",
				Email:    "dindik@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_dinkes",
				Email:    "dinkes@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_kebonsari",
				Email:    "puskesmas_kebonsari@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_gantrung",
				Email:    "puskesmas_gantrung@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_geger",
				Email:    "puskesmas_geger@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_kaibon",
				Email:    "puskesmas_kaibon@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_mlilir",
				Email:    "puskesmas_mlilir@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_bangunsari",
				Email:    "puskesmas_bangunsari@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_dagangan",
				Email:    "puskesmas_dagangan@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_jetis",
				Email:    "puskesmas_jetis@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_wungu",
				Email:    "puskesmas_wungu@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_mojopurno",
				Email:    "puskesmas_mojopurno@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_kare",
				Email:    "puskesmas_kare@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_gemarang",
				Email:    "puskesmas_gemarang@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_saradan",
				Email:    "puskesmas_saradan@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_sumbersari",
				Email:    "puskesmas_sumbersari@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_pilangkenceng",
				Email:    "puskesmas_pilangkenceng@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_krebet",
				Email:    "puskesmas_krebet@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_mejayan",
				Email:    "puskesmas_mejayan@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_klecorejo",
				Email:    "puskesmas_klecorejo@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_wonoasri",
				Email:    "puskesmas_wonoasri@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_balerejo",
				Email:    "puskesmas_balerejo@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_simo",
				Email:    "puskesmas_simo@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_madiun",
				Email:    "puskesmas_madiun@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_dimong",
				Email:    "puskesmas_dimong@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_sawahan",
				Email:    "puskesmas_sawahan@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_klegenserut",
				Email:    "puskesmas_klegenserut@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_puskesmas_jiwan",
				Email:    "puskesmas_jiwan@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_rsud_caruban",
				Email:    "rsud_caruban@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_rsud_dolopo",
				Email:    "rsud_dolopo@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_pupr",
				Email:    "pupr@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_perkim",
				Email:    "dinas_perkim@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_satpolpp_damkar",
				Email:    "satpolpp_damkar@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_bpbd",
				Email:    "bpbd@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_dinsos",
				Email:    "dinsos@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_disnaker",
				Email:    "disnaker@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_dkpp",
				Email:    "dkpp@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_dlh",
				Email:    "dlh@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_disdukcapil",
				Email:    "disdukcapil@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_dpmd",
				Email:    "dpmd@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_dinas_ppkb_ppa",
				Email:    "dinas_ppkb_ppa@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_dishub",
				Email:    "dishub@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kominfo",
				Email:    "kominfo@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_disperdagkop",
				Email:    "disperdagkop@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_dpmptsp",
				Email:    "dpmptsp@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_disparpora",
				Email:    "disparpora@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_perpus",
				Email:    "perpus@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_disperta",
				Email:    "disperta@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_sekda",
				Email:    "sekda@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_bag_adpem",
				Email:    "bag_adpem@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_bag_hukum",
				Email:    "bag_hukum@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_bag_pbj",
				Email:    "bag_pbj@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_bag_kesra",
				Email:    "bag_kesra@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_bag_perekonomian",
				Email:    "bag_perekonomian@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_bag_umum",
				Email:    "bag_umum@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_bag_organisasi",
				Email:    "bag_organisasi@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_bag_protokol",
				Email:    "bag_protokol@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_bag_adbang",
				Email:    "bag_adbang@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_setwan",
				Email:    "setwan@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_bapperida",
				Email:    "bapperida@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_bpkad",
				Email:    "bpkad@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_bapenda",
				Email:    "bapenda@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_bkpsdm",
				Email:    "bkpsdm@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_inspektorat",
				Email:    "inspektorat@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kecamatan_balerejo",
				Email:    "kecamatan_balerejo@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kecamatan_dagangan",
				Email:    "kecamatan_dagangan@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kecamatan_dolopo",
				Email:    "kecamatan_dolopo@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kelurahan_bangunsari_dolopo",
				Email:    "kelurahan_bangunsari_dolopo@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kelurahan_mlilir",
				Email:    "kelurahan_mlilir@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kecamatan_geger",
				Email:    "kecamatan_geger@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kecamatan_gemarang",
				Email:    "kecamatan_gemarang@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kecamatan_jiwan",
				Email:    "kecamatan_jiwan@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kecamatan_kebonsari",
				Email:    "kecamatan_kebonsari@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kecamatan_kare",
				Email:    "kecamatan_kare@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kecamatan_madiun",
				Email:    "kecamatan_madiun@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kelurahan_nglames",
				Email:    "kelurahan_nglames@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kecamatan_mejayan",
				Email:    "kecamatan_mejayan@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kelurahan_bangunsari_mejayan",
				Email:    "kelurahan_bangunsari_mejayan@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_keurahan_krajan",
				Email:    "keurahan_krajan@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kelurahan_pandean",
				Email:    "kelurahan_pandean@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kecamatan_pilangkenceng",
				Email:    "kecamatan_pilangkenceng@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kecamatan_sawahan",
				Email:    "kecamatan_sawahan@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kecamatan_saradan",
				Email:    "kecamatan_saradan@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kecamatan_wungu",
				Email:    "kecamatan_wungu@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kelurahan_wungu",
				Email:    "kelurahan_wungu@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kelurahan_munggut",
				Email:    "kelurahan_munggut@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_kecamatan_wonoasri",
				Email:    "kecamatan_wonoasri@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
		{
			user: domain.Users{
				Nip:      "admin_bakesbangpol",
				Email:    "bakesbangpol@madiunkabtest.com",
				Password: "KabKabMadiun2024",
				IsActive: true,
			},
			roleKeys: []string{"admin_opd"},
		},
	}

	for _, userData := range defaultUsers {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		userData.user.Password = string(hashedPassword)

		for _, roleKey := range userData.roleKeys {
			if role, exists := roleMap[roleKey]; exists {
				userData.user.Role = append(userData.user.Role, role)
			}
		}

		_, err = seeder.UserRepository.Create(ctx, tx, userData.user)
		if err != nil {
			return err
		}
	}
	log.Println("Users berhasil di-seed")
	return nil
}
