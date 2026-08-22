package dataseeder

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/model/domain/domainmaster"
	"ekak_kab_sleman/repository"
	"log"

	"github.com/google/uuid"
)

type PegawaiSeederImpl struct {
	DB                *sql.DB
	PegawaiRepository repository.PegawaiRepository
}

func NewPegawaiSeederImpl(db *sql.DB, pegawaiRepository repository.PegawaiRepository) *PegawaiSeederImpl {
	return &PegawaiSeederImpl{
		DB:                db,
		PegawaiRepository: pegawaiRepository,
	}
}

func (pegawai *PegawaiSeederImpl) Seed(ctx context.Context, tx *sql.Tx) error {
	// Cek dan buat pegawai pertama
	_, err := pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin1")
	if err == sql.ErrNoRows {
		superAdmin1 := domainmaster.Pegawai{
			Id:          "ADMIN-" + uuid.New().String()[:4],
			NamaPegawai: "super admin satu",
			Nip:         "admin1",
			KodeOpd:     "",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, superAdmin1)
		if err != nil {
			return err
		}
		log.Println("Pegawai Super Admin pertama berhasil di-seed")
	}

	// Cek dan buat pegawai kedua
	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin2")
	if err == sql.ErrNoRows {
		superAdmin2 := domainmaster.Pegawai{
			Id:          "ADMIN-" + uuid.New().String()[:4],
			NamaPegawai: "super admin dua",
			Nip:         "admin2",
			KodeOpd:     "",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, superAdmin2)
		if err != nil {
			return err
		}
		log.Println("Pegawai Super Admin kedua berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_dindik")
	if err == sql.ErrNoRows {
		adminDindik := domainmaster.Pegawai{
			Id:          "ADMIN-" + uuid.New().String()[:4],
			NamaPegawai: "admin dindik",
			Nip:         "admin_dindik",
			KodeOpd:     "1.01.2.22.0.00.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminDindik)
		if err != nil {
			return err
		}
		log.Println("Pegawai pertama berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_dinkes")
	if err == sql.ErrNoRows {
		adminDinkes := domainmaster.Pegawai{
			Id:          "ADMIN-" + uuid.New().String()[:4],
			NamaPegawai: "admin dinkes",
			Nip:         "admin_dinkes",
			KodeOpd:     "1.02.0.00.0.00.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminDinkes)
		if err != nil {
			return err
		}
		log.Println("Pegawai kedua berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_kebonsari")
	if err == sql.ErrNoRows {
		adminPuskesmasKebonsari := domainmaster.Pegawai{
			Id:          "ADMIN-KEBONSARI-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas kebonsari",
			Nip:         "admin_puskesmas_kebonsari",
			KodeOpd:     "1.02.0.00.0.00.01.0001",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasKebonsari)
		if err != nil {
			return err
		}
		log.Println("Pegawai ketiga berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_gantrung")
	if err == sql.ErrNoRows {
		adminPuskesmasGantrung := domainmaster.Pegawai{
			Id:          "ADMIN-GANTRUNG-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas gantrung",
			Nip:         "admin_puskesmas_gantrung",
			KodeOpd:     "1.02.0.00.0.00.01.0002",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasGantrung)
		if err != nil {
			return err
		}
		log.Println("Pegawai keempat berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_geger")
	if err == sql.ErrNoRows {
		adminPuskesmasGeger := domainmaster.Pegawai{
			Id:          "ADMIN-GEGER-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas geger",
			Nip:         "admin_puskesmas_geger",
			KodeOpd:     "1.02.0.00.0.00.01.0003",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasGeger)
		if err != nil {
			return err
		}
		log.Println("Pegawai kelima berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_kaibon")
	if err == sql.ErrNoRows {
		adminPuskesmasKaibon := domainmaster.Pegawai{
			Id:          "ADMIN-KAIBON-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas kaibon",
			Nip:         "admin_puskesmas_kaibon",
			KodeOpd:     "1.02.0.00.0.00.01.0004",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasKaibon)
		if err != nil {
			return err
		}
		log.Println("Pegawai keenam berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_mlilir")
	if err == sql.ErrNoRows {
		adminPuskesmasMlilir := domainmaster.Pegawai{
			Id:          "ADMIN-MLILIR-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas mlilir",
			Nip:         "admin_puskesmas_mlilir",
			KodeOpd:     "1.02.0.00.0.00.01.0005",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasMlilir)
		if err != nil {
			return err
		}
		log.Println("Pegawai ketujuh berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_bangunsari")
	if err == sql.ErrNoRows {
		adminPuskesmasBangunsari := domainmaster.Pegawai{
			Id:          "ADMIN-BANGUNSARI-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas bangunsari",
			Nip:         "admin_puskesmas_bangunsari",
			KodeOpd:     "1.02.0.00.0.00.01.0006",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasBangunsari)
		if err != nil {
			return err
		}
		log.Println("Pegawai kedelapan berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_dagangan")
	if err == sql.ErrNoRows {
		adminPuskesmasDagangan := domainmaster.Pegawai{
			Id:          "ADMIN-DAGANGAN-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas dagangan",
			Nip:         "admin_puskesmas_dagangan",
			KodeOpd:     "1.02.0.00.0.00.01.0007",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasDagangan)
		if err != nil {
			return err
		}
		log.Println("Pegawai kesembilan berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_jetis")
	if err == sql.ErrNoRows {
		adminPuskesmasJetis := domainmaster.Pegawai{
			Id:          "ADMIN-JETIS-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas jetis",
			Nip:         "admin_puskesmas_jetis",
			KodeOpd:     "1.02.0.00.0.00.01.0008",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasJetis)
		if err != nil {
			return err
		}
		log.Println("Pegawai kesepuluh berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_wungu")
	if err == sql.ErrNoRows {
		adminPuskesmasWungu := domainmaster.Pegawai{
			Id:          "ADMIN-WUNGU-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas wungu",
			Nip:         "admin_puskesmas_wungu",
			KodeOpd:     "1.02.0.00.0.00.01.0009",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasWungu)
		if err != nil {
			return err
		}
		log.Println("Pegawai kesebelas berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_mojopurno")
	if err == sql.ErrNoRows {
		adminPuskesmasMojopurno := domainmaster.Pegawai{
			Id:          "ADMIN-MOJOPURNO-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas mojopurno",
			Nip:         "admin_puskesmas_mojopurno",
			KodeOpd:     "1.02.0.00.0.00.01.0010",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasMojopurno)
		if err != nil {
			return err
		}
		log.Println("Pegawai kedua belas berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_kare")
	if err == sql.ErrNoRows {
		adminPuskesmasKare := domainmaster.Pegawai{
			Id:          "ADMIN-KARE-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas kare",
			Nip:         "admin_puskesmas_kare",
			KodeOpd:     "1.02.0.00.0.00.01.0011",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasKare)
		if err != nil {
			return err
		}
		log.Println("Pegawai ketiga belas berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_gemarang")
	if err == sql.ErrNoRows {
		adminPuskesmasGemarang := domainmaster.Pegawai{
			Id:          "ADMIN-GEMARANG-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas gemarang",
			Nip:         "admin_puskesmas_gemarang",
			KodeOpd:     "1.02.0.00.0.00.01.0012",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasGemarang)
		if err != nil {
			return err
		}
		log.Println("Pegawai keempat belas berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_saradan")
	if err == sql.ErrNoRows {
		adminPuskesmasSaradan := domainmaster.Pegawai{
			Id:          "ADMIN-SARADAN-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas saradan",
			Nip:         "admin_puskesmas_saradan",
			KodeOpd:     "1.02.0.00.0.00.01.0013",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasSaradan)
		if err != nil {
			return err
		}
		log.Println("Pegawai kelima belas berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_sumbersari")
	if err == sql.ErrNoRows {
		adminPuskesmasSumbersari := domainmaster.Pegawai{
			Id:          "ADMIN-SUMBERSARI-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas sumbersari",
			Nip:         "admin_puskesmas_sumbersari",
			KodeOpd:     "1.02.0.00.0.00.01.0014",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasSumbersari)
		if err != nil {
			return err
		}
		log.Println("Pegawai enam belas berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_pilangkenceng")
	if err == sql.ErrNoRows {
		adminPuskesmasPilangkenceng := domainmaster.Pegawai{
			Id:          "ADMIN-PILANGKENCENG-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas pilangkenceng",
			Nip:         "admin_puskesmas_pilangkenceng",
			KodeOpd:     "1.02.0.00.0.00.01.0015",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasPilangkenceng)
		if err != nil {
			return err
		}
		log.Println("Pegawai tujuh belas berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_krebet")
	if err == sql.ErrNoRows {
		adminPuskesmasKrebet := domainmaster.Pegawai{
			Id:          "ADMIN-KREBET-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas krebet",
			Nip:         "admin_puskesmas_krebet",
			KodeOpd:     "1.02.0.00.0.00.01.0016",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasKrebet)
		if err != nil {
			return err
		}
		log.Println("Pegawai delapan belas berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_mejayan")
	if err == sql.ErrNoRows {
		adminPuskesmasMejayan := domainmaster.Pegawai{
			Id:          "ADMIN-MEJAYAN-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas mejayan",
			Nip:         "admin_puskesmas_mejayan",
			KodeOpd:     "1.02.0.00.0.00.01.0017",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasMejayan)
		if err != nil {
			return err
		}
		log.Println("Pegawai sembilan belas berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_klecorejo")
	if err == sql.ErrNoRows {
		adminPuskesmasKlecorejo := domainmaster.Pegawai{
			Id:          "ADMIN-KLECOREJO-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas klecorejo",
			Nip:         "admin_puskesmas_klecorejo",
			KodeOpd:     "1.02.0.00.0.00.01.0018",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasKlecorejo)
		if err != nil {
			return err
		}
		log.Println("Pegawai dua puluh berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_wonoasri")
	if err == sql.ErrNoRows {
		adminPuskesmasWonoasri := domainmaster.Pegawai{
			Id:          "ADMIN-WONOASRI-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas wonoasri",
			Nip:         "admin_puskesmas_wonoasri",
			KodeOpd:     "1.02.0.00.0.00.01.0019",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasWonoasri)
		if err != nil {
			return err
		}
		log.Println("Pegawai dua puluh satu berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_balerejo")
	if err == sql.ErrNoRows {
		adminPuskesmasBalerejo := domainmaster.Pegawai{
			Id:          "ADMIN-BALEREJO-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas balerejo",
			Nip:         "admin_puskesmas_balerejo",
			KodeOpd:     "1.02.0.00.0.00.01.0020",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasBalerejo)
		if err != nil {
			return err
		}
		log.Println("Pegawai dua puluh dua berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_simo")
	if err == sql.ErrNoRows {
		adminPuskesmasSimo := domainmaster.Pegawai{
			Id:          "ADMIN-SIMO-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas simo",
			Nip:         "admin_puskesmas_simo",
			KodeOpd:     "1.02.0.00.0.00.01.0021",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasSimo)
		if err != nil {
			return err
		}
		log.Println("Pegawai dua puluh tiga berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_madiun")
	if err == sql.ErrNoRows {
		adminPuskesmasMadiun := domainmaster.Pegawai{
			Id:          "ADMIN-MADIUN-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas madiun",
			Nip:         "admin_puskesmas_madiun",
			KodeOpd:     "1.02.0.00.0.00.01.0022",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasMadiun)
		if err != nil {
			return err
		}
		log.Println("Pegawai dua puluh empat berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_dimong")
	if err == sql.ErrNoRows {
		adminPuskesmasDimong := domainmaster.Pegawai{
			Id:          "ADMIN-DIMONG-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas dimong",
			Nip:         "admin_puskesmas_dimong",
			KodeOpd:     "1.02.0.00.0.00.01.0023",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasDimong)
		if err != nil {
			return err
		}
		log.Println("Pegawai dua puluh lima berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_sawahan")
	if err == sql.ErrNoRows {
		adminPuskesmasSawahan := domainmaster.Pegawai{
			Id:          "ADMIN-SAWAHAN-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas sawahan",
			Nip:         "admin_puskesmas_sawahan",
			KodeOpd:     "1.02.0.00.0.00.01.0024",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasSawahan)
		if err != nil {
			return err
		}
		log.Println("Pegawai dua puluh enam berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_klegenserut")
	if err == sql.ErrNoRows {
		adminPuskesmasKlegenserut := domainmaster.Pegawai{
			Id:          "ADMIN-KLEGENSERT-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas klegenserut",
			Nip:         "admin_puskesmas_klegenserut",
			KodeOpd:     "1.02.0.00.0.00.01.0025",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasKlegenserut)
		if err != nil {
			return err
		}
		log.Println("Pegawai dua puluh tujuh berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_puskesmas_jiwan")
	if err == sql.ErrNoRows {
		adminPuskesmasJiwan := domainmaster.Pegawai{
			Id:          "ADMIN-JIWAN-" + uuid.New().String()[:4],
			NamaPegawai: "admin puskesmas jiwan",
			Nip:         "admin_puskesmas_jiwan",
			KodeOpd:     "1.02.0.00.0.00.01.0026",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPuskesmasJiwan)
		if err != nil {
			return err
		}
		log.Println("Pegawai dua puluh delapan berhasil di-seed")
	}
	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_rsud_caruban")
	if err == sql.ErrNoRows {
		adminRsudCaruban := domainmaster.Pegawai{
			Id:          "ADMIN-RS-" + uuid.New().String()[:4],
			NamaPegawai: "admin rsud caruban",
			Nip:         "admin_rsud_caruban",
			KodeOpd:     "1.02.0.00.0.00.01.0027",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminRsudCaruban)
		if err != nil {
			return err
		}
		log.Println("Pegawai dua puluh sembilan berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_rsud_dolopo")
	if err == sql.ErrNoRows {
		adminRsudDolopo := domainmaster.Pegawai{
			Id:          "ADMIN-RS-" + uuid.New().String()[:4],
			NamaPegawai: "admin rsud dolopo",
			Nip:         "admin_rsud_dolopo",
			KodeOpd:     "1.02.0.00.0.00.01.0028",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminRsudDolopo)
		if err != nil {
			return err
		}
		log.Println("Pegawai tiga puluh berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_pupr")
	if err == sql.ErrNoRows {
		adminPupr := domainmaster.Pegawai{
			Id:          "ADMIN-PUPR-" + uuid.New().String()[:4],
			NamaPegawai: "admin pupr",
			Nip:         "admin_pupr",
			KodeOpd:     "1.03.0.00.0.00.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPupr)
		if err != nil {
			return err
		}
		log.Println("Pegawai tiga puluh satu berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_perkim")
	if err == sql.ErrNoRows {
		adminPerkim := domainmaster.Pegawai{
			Id:          "ADMIN-PERKIM-" + uuid.New().String()[:4],
			NamaPegawai: "admin perkim",
			Nip:         "admin_perkim",
			KodeOpd:     "1.04.2.10.0.00.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPerkim)
		if err != nil {
			return err
		}
		log.Println("Pegawai tiga puluh dua berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_satpolpp_damkar")
	if err == sql.ErrNoRows {
		adminSatpolppDamkar := domainmaster.Pegawai{
			Id:          "ADMIN-SATPOLPP-DAMKAR-" + uuid.New().String()[:4],
			NamaPegawai: "admin satpolpp damkar",
			Nip:         "admin_satpolpp_damkar",
			KodeOpd:     "1.05.0.00.0.00.02.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminSatpolppDamkar)
		if err != nil {
			return err
		}
		log.Println("Pegawai tiga puluh tiga berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_bpbd")
	if err == sql.ErrNoRows {
		adminBpbd := domainmaster.Pegawai{
			Id:          "ADMIN-BPBD-" + uuid.New().String()[:4],
			NamaPegawai: "admin bpbd",
			Nip:         "admin_bpbd",
			KodeOpd:     "1.05.0.00.0.00.03.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminBpbd)
		if err != nil {
			return err
		}
		log.Println("Pegawai tiga puluh empat berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_dinsos")
	if err == sql.ErrNoRows {
		adminDinsos := domainmaster.Pegawai{
			Id:          "ADMIN-DINSOS-" + uuid.New().String()[:4],
			NamaPegawai: "admin dinsos",
			Nip:         "admin_dinsos",
			KodeOpd:     "1.06.0.00.0.00.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminDinsos)
		if err != nil {
			return err
		}
		log.Println("Pegawai tiga puluh lima berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_disnaker")
	if err == sql.ErrNoRows {
		adminDisnaker := domainmaster.Pegawai{
			Id:          "ADMIN-DISNAKER-" + uuid.New().String()[:4],
			NamaPegawai: "admin disnaker",
			Nip:         "admin_disnaker",
			KodeOpd:     "2.07.3.31.3.32.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminDisnaker)
		if err != nil {
			return err
		}
		log.Println("Pegawai tiga puluh enam berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_dkpp")
	if err == sql.ErrNoRows {
		adminDkpp := domainmaster.Pegawai{
			Id:          "ADMIN-DKPP-" + uuid.New().String()[:4],
			NamaPegawai: "admin dkpp",
			Nip:         "admin_dkpp",
			KodeOpd:     "2.09.3.27.0.00.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminDkpp)
		if err != nil {
			return err
		}
		log.Println("Pegawai tiga puluh tujuh berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_dlh")
	if err == sql.ErrNoRows {
		adminDlh := domainmaster.Pegawai{
			Id:          "ADMIN-DLH-" + uuid.New().String()[:4],
			NamaPegawai: "admin dlh",
			Nip:         "admin_dlh",
			KodeOpd:     "2.11.0.00.0.00.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminDlh)
		if err != nil {
			return err
		}
		log.Println("Pegawai tiga puluh delapan berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_disdukcapil")
	if err == sql.ErrNoRows {
		adminDisdukcapil := domainmaster.Pegawai{
			Id:          "ADMIN-DISDUKCAPIL-" + uuid.New().String()[:4],
			NamaPegawai: "admin disdukcapil",
			Nip:         "admin_disdukcapil",
			KodeOpd:     "2.12.0.00.0.00.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminDisdukcapil)
		if err != nil {
			return err
		}
		log.Println("Pegawai tiga puluh sembilan berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_dpmd")
	if err == sql.ErrNoRows {
		adminDpmd := domainmaster.Pegawai{
			Id:          "ADMIN-DPMD-" + uuid.New().String()[:4],
			NamaPegawai: "admin dpmd",
			Nip:         "admin_dpmd",
			KodeOpd:     "2.13.0.00.0.00.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminDpmd)
		if err != nil {
			return err
		}
		log.Println("Pegawai empat puluh berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_dinas_ppkb_ppa")
	if err == sql.ErrNoRows {
		adminDinasPpkbPpa := domainmaster.Pegawai{
			Id:          "ADMIN-DINAS-PPKB-PPA-" + uuid.New().String()[:4],
			NamaPegawai: "admin dinas ppkb ppa",
			Nip:         "admin_dinas_ppkb_ppa",
			KodeOpd:     "2.14.2.08.0.00.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminDinasPpkbPpa)
		if err != nil {
			return err
		}
		log.Println("Pegawai empat puluh satu berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_dishub")
	if err == sql.ErrNoRows {
		adminDishub := domainmaster.Pegawai{
			Id:          "ADMIN-DISHUB-" + uuid.New().String()[:4],
			NamaPegawai: "admin dishub",
			Nip:         "admin_dishub",
			KodeOpd:     "2.15.0.00.0.00.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminDishub)
		if err != nil {
			return err
		}
		log.Println("Pegawai empat puluh dua berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kominfo")
	if err == sql.ErrNoRows {
		adminKominfo := domainmaster.Pegawai{
			Id:          "ADMIN-KOMINFO-" + uuid.New().String()[:4],
			NamaPegawai: "admin kominfo",
			Nip:         "admin_kominfo",
			KodeOpd:     "2.16.2.20.2.21.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKominfo)
		if err != nil {
			return err
		}
		log.Println("Pegawai empat puluh tiga berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_disperdagkop")
	if err == sql.ErrNoRows {
		adminDisperdagkop := domainmaster.Pegawai{
			Id:          "ADMIN-DISPERDAGKOP-" + uuid.New().String()[:4],
			NamaPegawai: "admin disperdagkop",
			Nip:         "admin_disperdagkop",
			KodeOpd:     "2.17.3.30.0.00.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminDisperdagkop)
		if err != nil {
			return err
		}
		log.Println("Pegawai empat puluh empat berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_dpmptsp")
	if err == sql.ErrNoRows {
		adminDpmptsp := domainmaster.Pegawai{
			Id:          "ADMIN-DPMPTSP-" + uuid.New().String()[:4],
			NamaPegawai: "admin dpmptsp",
			Nip:         "admin_dpmptsp",
			KodeOpd:     "2.18.0.00.0.00.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminDpmptsp)
		if err != nil {
			return err
		}
		log.Println("Pegawai empat puluh lima berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_disparpora")
	if err == sql.ErrNoRows {
		adminDisparpora := domainmaster.Pegawai{
			Id:          "ADMIN-DISPARPORA-" + uuid.New().String()[:4],
			NamaPegawai: "admin disparpora",
			Nip:         "admin_disparpora",
			KodeOpd:     "2.19.3.26.0.00.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminDisparpora)
		if err != nil {
			return err
		}
		log.Println("Pegawai empat puluh enam berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_perpus")
	if err == sql.ErrNoRows {
		adminPerpus := domainmaster.Pegawai{
			Id:          "ADMIN-PERPUST-" + uuid.New().String()[:4],
			NamaPegawai: "admin perpus",
			Nip:         "admin_perpus",
			KodeOpd:     "2.23.2.24.0.00.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminPerpus)
		if err != nil {
			return err
		}
		log.Println("Pegawai empat puluh tujuh berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_disperta")
	if err == sql.ErrNoRows {
		adminDisperta := domainmaster.Pegawai{
			Id:          "ADMIN-DISPERTA-" + uuid.New().String()[:4],
			NamaPegawai: "admin disperta",
			Nip:         "admin_disperta",
			KodeOpd:     "3.27.3.25.0.00.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminDisperta)
		if err != nil {
			return err
		}
		log.Println("Pegawai empat puluh delapan berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_sekda")
	if err == sql.ErrNoRows {
		adminSekda := domainmaster.Pegawai{
			Id:          "ADMIN-SEKDA-" + uuid.New().String()[:4],
			NamaPegawai: "admin sekda",
			Nip:         "admin_sekda",
			KodeOpd:     "4.01.5.06.0.00.03.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminSekda)
		if err != nil {
			return err
		}
		log.Println("Pegawai empat puluh sembilan berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_bag_adpem")
	if err == sql.ErrNoRows {
		adminBagAdpem := domainmaster.Pegawai{
			Id:          "ADMIN-BAG-ADPEM-" + uuid.New().String()[:4],
			NamaPegawai: "admin bag adpem",
			Nip:         "admin_bag_adpem",
			KodeOpd:     "4.01.5.06.0.00.03.0001",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminBagAdpem)
		if err != nil {
			return err
		}
		log.Println("Pegawai lima puluh berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_bag_hukum")
	if err == sql.ErrNoRows {
		adminBagHukum := domainmaster.Pegawai{
			Id:          "ADMIN-BAG-HUKUM-" + uuid.New().String()[:4],
			NamaPegawai: "admin bag adpem",
			Nip:         "admin_bag_hukum",
			KodeOpd:     "4.01.5.06.0.00.03.0002",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminBagHukum)
		if err != nil {
			return err
		}
		log.Println("Pegawai lima puluh satu berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_bag_pbj")
	if err == sql.ErrNoRows {
		adminBagPbj := domainmaster.Pegawai{
			Id:          "ADMIN-BAG-PBJ-" + uuid.New().String()[:4],
			NamaPegawai: "admin bag pbj",
			Nip:         "admin_bag_pbj",
			KodeOpd:     "4.01.5.06.0.00.03.0003",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminBagPbj)
		if err != nil {
			return err
		}
		log.Println("Pegawai lima puluh dua berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_bag_kesra")
	if err == sql.ErrNoRows {
		adminBagKesra := domainmaster.Pegawai{
			Id:          "ADMIN-BAG-KESRA-" + uuid.New().String()[:4],
			NamaPegawai: "admin bag kesra",
			Nip:         "admin_bag_kesra",
			KodeOpd:     "4.01.5.06.0.00.03.0005",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminBagKesra)
		if err != nil {
			return err
		}
		log.Println("Pegawai lima puluh tiga berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_bag_perekonomian")
	if err == sql.ErrNoRows {
		adminBagPerekonomian := domainmaster.Pegawai{
			Id:          "ADMIN-BAG-PEREKONOMIAN-" + uuid.New().String()[:4],
			NamaPegawai: "admin bag perekonomian",
			Nip:         "admin_bag_perekonomian",
			KodeOpd:     "4.01.5.06.0.00.03.0006",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminBagPerekonomian)
		if err != nil {
			return err
		}
		log.Println("Pegawai lima puluh empat berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_bag_umum")
	if err == sql.ErrNoRows {
		adminBagUmum := domainmaster.Pegawai{
			Id:          "ADMIN-BAG-UMUM-" + uuid.New().String()[:4],
			NamaPegawai: "admin bag umum",
			Nip:         "admin_bag_umum",
			KodeOpd:     "4.01.5.06.0.00.03.0007",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminBagUmum)
		if err != nil {
			return err
		}
		log.Println("Pegawai lima puluh lima berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_bag_organisasi")
	if err == sql.ErrNoRows {
		adminBagOrganisasi := domainmaster.Pegawai{
			Id:          "ADMIN-BAG-ORGANISASI-" + uuid.New().String()[:4],
			NamaPegawai: "admin bag organisasi",
			Nip:         "admin_bag_organisasi",
			KodeOpd:     "4.01.5.06.0.00.03.0008",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminBagOrganisasi)
		if err != nil {
			return err
		}
		log.Println("Pegawai lima puluh enam berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_bag_protokol")
	if err == sql.ErrNoRows {
		adminBagProtokol := domainmaster.Pegawai{
			Id:          "ADMIN-BAG-ORGANISASI-" + uuid.New().String()[:4],
			NamaPegawai: "admin bag protokol",
			Nip:         "admin_bag_protokol",
			KodeOpd:     "4.01.5.06.0.00.03.0009",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminBagProtokol)
		if err != nil {
			return err
		}
		log.Println("Pegawai lima puluh tujuh berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_bag_adbang")
	if err == sql.ErrNoRows {
		adminBagAdbang := domainmaster.Pegawai{
			Id:          "ADMIN-BAG-ADBANG-" + uuid.New().String()[:4],
			NamaPegawai: "admin bag adbang",
			Nip:         "admin_bag_adbang",
			KodeOpd:     "4.01.5.06.0.00.03.0010",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminBagAdbang)
		if err != nil {
			return err
		}
		log.Println("Pegawai lima puluh delapan berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_setwan")
	if err == sql.ErrNoRows {
		adminSetwan := domainmaster.Pegawai{
			Id:          "ADMIN-SETWAN-" + uuid.New().String()[:4],
			NamaPegawai: "admin setwan",
			Nip:         "admin_setwan",
			KodeOpd:     "4.02.0.00.0.00.04.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminSetwan)
		if err != nil {
			return err
		}
		log.Println("Pegawai lima puluh sembilan berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_bapperida")
	if err == sql.ErrNoRows {
		adminBapperida := domainmaster.Pegawai{
			Id:          "ADMIN-BAPPERIDA-" + uuid.New().String()[:4],
			NamaPegawai: "admin bapperida",
			Nip:         "admin_bapperida",
			KodeOpd:     "5.01.5.05.0.00.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminBapperida)
		if err != nil {
			return err
		}
		log.Println("Pegawai enam puluh berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_bpkad")
	if err == sql.ErrNoRows {
		adminBpkad := domainmaster.Pegawai{
			Id:          "ADMIN-BPKAD-" + uuid.New().String()[:4],
			NamaPegawai: "admin bpkad",
			Nip:         "admin_bpkad",
			KodeOpd:     "5.02.0.00.0.00.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminBpkad)
		if err != nil {
			return err
		}
		log.Println("Pegawai enam puluh satu berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_bapenda")
	if err == sql.ErrNoRows {
		adminBapenda := domainmaster.Pegawai{
			Id:          "ADMIN-BAPENDA-" + uuid.New().String()[:4],
			NamaPegawai: "admin bapenda",
			Nip:         "admin_bapenda",
			KodeOpd:     "5.02.0.00.0.00.02.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminBapenda)
		if err != nil {
			return err
		}
		log.Println("Pegawai enam puluh dua berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_bkpsdm")
	if err == sql.ErrNoRows {
		adminBkpsdm := domainmaster.Pegawai{
			Id:          "ADMIN-BKPSDM-" + uuid.New().String()[:4],
			NamaPegawai: "admin bkpsdm",
			Nip:         "admin_bkpsdm",
			KodeOpd:     "5.03.5.04.0.00.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminBkpsdm)
		if err != nil {
			return err
		}
		log.Println("Pegawai enam puluh tiga berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_inspektorat")
	if err == sql.ErrNoRows {
		adminInspektorat := domainmaster.Pegawai{
			Id:          "ADMIN-INSPEKTORAT-" + uuid.New().String()[:4],
			NamaPegawai: "admin inspektorat",
			Nip:         "admin_inspektorat",
			KodeOpd:     "6.01.0.00.0.00.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminInspektorat)
		if err != nil {
			return err
		}
		log.Println("Pegawai enam puluh empat berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kecamatan_balerejo")
	if err == sql.ErrNoRows {
		adminKecamatanBalerejo := domainmaster.Pegawai{
			Id:          "ADMIN-KECAMATAN-BALERJO-" + uuid.New().String()[:4],
			NamaPegawai: "admin kecamatan balerejo",
			Nip:         "admin_kecamatan_balerejo",
			KodeOpd:     "7.01.0.00.0.00.05.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKecamatanBalerejo)
		if err != nil {
			return err
		}
		log.Println("Pegawai enam puluh lima berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kecamatan_dagangan")
	if err == sql.ErrNoRows {
		adminKecamatanDagangan := domainmaster.Pegawai{
			Id:          "ADMIN-KECAMATAN-DAGANGAN-" + uuid.New().String()[:4],
			NamaPegawai: "admin kecamatan dagangan",
			Nip:         "admin_kecamatan_dagangan",
			KodeOpd:     "7.01.0.00.0.00.06.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKecamatanDagangan)
		if err != nil {
			return err
		}
		log.Println("Pegawai enam puluh enam berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kecamatan_dolopo")
	if err == sql.ErrNoRows {
		adminKecamatanDolopo := domainmaster.Pegawai{
			Id:          "ADMIN-KECAMATAN-DOLOPO-" + uuid.New().String()[:4],
			NamaPegawai: "admin kecamatan dolopo",
			Nip:         "admin_kecamatan_dolopo",
			KodeOpd:     "7.01.0.00.0.00.07.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKecamatanDolopo)
		if err != nil {
			return err
		}
		log.Println("Pegawai enam puluh tujuh berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kelurahan_bangunsari_dolopo")
	if err == sql.ErrNoRows {
		adminKelurahanBangunsariDolopo := domainmaster.Pegawai{
			Id:          "ADMIN-KEL-BANGUNSARI-DOLOPO-" + uuid.New().String()[:4],
			NamaPegawai: "admin kelurahan bangunsari dolopo",
			Nip:         "admin_kelurahan_bangunsari_dolopo",
			KodeOpd:     "7.01.0.00.0.00.07.0001",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKelurahanBangunsariDolopo)
		if err != nil {
			return err
		}
		log.Println("Pegawai enam puluh delapan berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kelurahan_mlilir")
	if err == sql.ErrNoRows {
		adminKelurahanMlilir := domainmaster.Pegawai{
			Id:          "ADMIN-KEL-MLILIR-" + uuid.New().String()[:4],
			NamaPegawai: "admin kelurahan mlilir",
			Nip:         "admin_kelurahan_mlilir",
			KodeOpd:     "7.01.0.00.0.00.07.0002",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKelurahanMlilir)
		if err != nil {
			return err
		}
		log.Println("Pegawai enam puluh sembilan berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kecamatan_geger")
	if err == sql.ErrNoRows {
		adminKecamatanGeger := domainmaster.Pegawai{
			Id:          "ADMIN-KECAMATAN-GEGER-" + uuid.New().String()[:4],
			NamaPegawai: "admin kecamatan geger",
			Nip:         "admin_kecamatan_geger",
			KodeOpd:     "7.01.0.00.0.00.08.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKecamatanGeger)
		if err != nil {
			return err
		}
		log.Println("Pegawai tujuh puluh berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kecamatan_gemarang")
	if err == sql.ErrNoRows {
		adminKecamatanGemarang := domainmaster.Pegawai{
			Id:          "ADMIN-KECAMATAN-GEMARANG-" + uuid.New().String()[:4],
			NamaPegawai: "admin kecamatan gemarang",
			Nip:         "admin_kecamatan_gemarang",
			KodeOpd:     "7.01.0.00.0.00.09.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKecamatanGemarang)
		if err != nil {
			return err
		}
		log.Println("Pegawai tujuh puluh satu berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kecamatan_jiwan")
	if err == sql.ErrNoRows {
		adminKecamatanJiwan := domainmaster.Pegawai{
			Id:          "ADMIN-KECAMATAN-JIWAN-" + uuid.New().String()[:4],
			NamaPegawai: "admin kecamatan jiwan",
			Nip:         "admin_kecamatan_jiwan",
			KodeOpd:     "7.01.0.00.0.00.10.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKecamatanJiwan)
		if err != nil {
			return err
		}
		log.Println("Pegawai tujuh puluh dua berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kecamatan_kebonsari")
	if err == sql.ErrNoRows {
		adminKecamatanKebonsari := domainmaster.Pegawai{
			Id:          "ADMIN-KECAMATAN-KEBONSAARI-" + uuid.New().String()[:4],
			NamaPegawai: "admin kecamatan kebonsari",
			Nip:         "admin_kecamatan_kebonsari",
			KodeOpd:     "7.01.0.00.0.00.11.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKecamatanKebonsari)
		if err != nil {
			return err
		}
		log.Println("Pegawai tujuh puluh tiga berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kecamatan_kare")
	if err == sql.ErrNoRows {
		adminKecamatanKare := domainmaster.Pegawai{
			Id:          "ADMIN-KECAMATAN-KARE-" + uuid.New().String()[:4],
			NamaPegawai: "admin kecamatan kare",
			Nip:         "admin_kecamatan_kare",
			KodeOpd:     "7.01.0.00.0.00.12.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKecamatanKare)
		if err != nil {
			return err
		}
		log.Println("Pegawai tujuh puluh empat berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kecamatan_madiun")
	if err == sql.ErrNoRows {
		adminKecamatanMadiun := domainmaster.Pegawai{
			Id:          "ADMIN-KECAMATAN-MADIUN-" + uuid.New().String()[:4],
			NamaPegawai: "admin kecamatan madiun",
			Nip:         "admin_kecamatan_madiun",
			KodeOpd:     "7.01.0.00.0.00.13.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKecamatanMadiun)
		if err != nil {
			return err
		}
		log.Println("Pegawai tujuh puluh lima berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kelurahan_nglames")
	if err == sql.ErrNoRows {
		adminKelurahanNglames := domainmaster.Pegawai{
			Id:          "ADMIN-KEL-NGLAMES-" + uuid.New().String()[:4],
			NamaPegawai: "admin kelurahan nglames",
			Nip:         "admin_kelurahan_nglames",
			KodeOpd:     "7.01.0.00.0.00.13.0001",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKelurahanNglames)
		if err != nil {
			return err
		}
		log.Println("Pegawai tujuh puluh enam berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kecamatan_mejayan")
	if err == sql.ErrNoRows {
		adminKecamatanMejayan := domainmaster.Pegawai{
			Id:          "ADMIN-KECAMATAN-MEJAYAN-" + uuid.New().String()[:4],
			NamaPegawai: "admin kecamatan mejayan",
			Nip:         "admin_kecamatan_mejayan",
			KodeOpd:     "7.01.0.00.0.00.14.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKecamatanMejayan)
		if err != nil {
			return err
		}
		log.Println("Pegawai tujuh puluh tujuh berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kelurahan_bangunsari_mejayan")
	if err == sql.ErrNoRows {
		adminKelurahanBangunsariMejayan := domainmaster.Pegawai{
			Id:          "ADMIN-KEL-BANGUNSARI-MEJAYAN-" + uuid.New().String()[:4],
			NamaPegawai: "admin kelurahan bangunsari mejayan",
			Nip:         "admin_kelurahan_bangunsari_mejayan",
			KodeOpd:     "7.01.0.00.0.00.14.00011",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKelurahanBangunsariMejayan)
		if err != nil {
			return err
		}
		log.Println("Pegawai tujuh puluh delapan berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kelurahan_krajan")
	if err == sql.ErrNoRows {
		adminKelurahanKrajan := domainmaster.Pegawai{
			Id:          "ADMIN-KEL-KRAJAN-" + uuid.New().String()[:4],
			NamaPegawai: "admin kelurahan krajan",
			Nip:         "admin_kelurahan_krajan",
			KodeOpd:     "7.01.0.00.0.00.14.0002",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKelurahanKrajan)
		if err != nil {
			return err
		}
		log.Println("Pegawai tujuh puluh sembilan berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kelurahan_pandean")
	if err == sql.ErrNoRows {
		adminKelurahanPandean := domainmaster.Pegawai{
			Id:          "ADMIN-KEL-PANDEAN-" + uuid.New().String()[:4],
			NamaPegawai: "admin kelurahan pandean",
			Nip:         "admin_kelurahan_pandean",
			KodeOpd:     "7.01.0.00.0.00.14.0003",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKelurahanPandean)
		if err != nil {
			return err
		}
		log.Println("Pegawai delapan puluh berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kecamatan_pilangkenceng")
	if err == sql.ErrNoRows {
		adminKecamatanPilangkenceng := domainmaster.Pegawai{
			Id:          "ADMIN-KECAMATAN-PILANGKENCENG-" + uuid.New().String()[:4],
			NamaPegawai: "admin kecamatan pilangkenceng",
			Nip:         "admin_kecamatan_pilangkenceng",
			KodeOpd:     "7.01.0.00.0.00.15.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKecamatanPilangkenceng)
		if err != nil {
			return err
		}
		log.Println("Pegawai delapan puluh satu berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kecamatan_sawahan")
	if err == sql.ErrNoRows {
		adminKecamatanSawahan := domainmaster.Pegawai{
			Id:          "ADMIN-KECAMATAN-SAWAHAN-" + uuid.New().String()[:4],
			NamaPegawai: "admin kecamatan sawahan",
			Nip:         "admin_kecamatan_sawahan",
			KodeOpd:     "7.01.0.00.0.00.16.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKecamatanSawahan)
		if err != nil {
			return err
		}
		log.Println("Pegawai delapan puluh dua berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kecamatan_saradan")
	if err == sql.ErrNoRows {
		adminKecamatanSaradan := domainmaster.Pegawai{
			Id:          "ADMIN-KECAMATAN-SARADAN-" + uuid.New().String()[:4],
			NamaPegawai: "admin kecamatan saradan",
			Nip:         "admin_kecamatan_saradan",
			KodeOpd:     "7.01.0.00.0.00.17.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKecamatanSaradan)
		if err != nil {
			return err
		}
		log.Println("Pegawai delapan puluh tiga berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kecamatan_wungu")
	if err == sql.ErrNoRows {
		adminKecamatanWungu := domainmaster.Pegawai{
			Id:          "ADMIN-KECAMATAN-WUNGU-" + uuid.New().String()[:4],
			NamaPegawai: "admin kecamatan wungu",
			Nip:         "admin_kecamatan_wungu",
			KodeOpd:     "7.01.0.00.0.00.18.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKecamatanWungu)
		if err != nil {
			return err
		}
		log.Println("Pegawai delapan puluh empat berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kelurahan_wungu")
	if err == sql.ErrNoRows {
		adminKelurahanWungu := domainmaster.Pegawai{
			Id:          "ADMIN-KEL-WUNGU-" + uuid.New().String()[:4],
			NamaPegawai: "admin kelurahan wungu",
			Nip:         "admin_kelurahan_wungu",
			KodeOpd:     "7.01.0.00.0.00.18.0001",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKelurahanWungu)
		if err != nil {
			return err
		}
		log.Println("Pegawai delapan puluh lima berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kelurahan_munggut")
	if err == sql.ErrNoRows {
		adminKelurahanMunggut := domainmaster.Pegawai{
			Id:          "ADMIN-KEL-MUNGGUT-" + uuid.New().String()[:4],
			NamaPegawai: "admin kelurahan munggut",
			Nip:         "admin_kelurahan_munggut",
			KodeOpd:     "7.01.0.00.0.00.18.0002",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKelurahanMunggut)
		if err != nil {
			return err
		}
		log.Println("Pegawai delapan puluh enam berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_kecamatan_wonoasri")
	if err == sql.ErrNoRows {
		adminKecamatanWonoasri := domainmaster.Pegawai{
			Id:          "ADMIN-KECAMATAN-WONOASRI-" + uuid.New().String()[:4],
			NamaPegawai: "admin kecamatan wonoasri",
			Nip:         "admin_kecamatan_wonoasri",
			KodeOpd:     "7.01.0.00.0.00.19.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminKecamatanWonoasri)
		if err != nil {
			return err
		}
		log.Println("Pegawai delapan puluh tujuh berhasil di-seed")
	}

	_, err = pegawai.PegawaiRepository.FindByNip(ctx, tx, "admin_bakesbangpol")
	if err == sql.ErrNoRows {
		adminBakesbangpol := domainmaster.Pegawai{
			Id:          "ADMIN-BAKESBANGPOL-" + uuid.New().String()[:4],
			NamaPegawai: "admin bakesbangpol",
			Nip:         "admin_bakesbangpol",
			KodeOpd:     "8.01.0.00.0.00.01.0000",
		}

		_, err = pegawai.PegawaiRepository.Create(ctx, tx, adminBakesbangpol)
		if err != nil {
			return err
		}
		log.Println("Pegawai delapan puluh delapan berhasil di-seed")
	}

	return nil
}
