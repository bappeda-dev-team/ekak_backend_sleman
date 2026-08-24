package service

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/domain"
	"ekak_kab_sleman/model/web/usulan"
	"ekak_kab_sleman/repository"
	"fmt"

	"github.com/google/uuid"
)

type UsulanMusrebangServiceImpl struct {
	usulanMusrebangRepository repository.UsulanMusrebangRepository
	rencanaKinerjaRepository  repository.RencanaKinerjaRepository
	opdRepository             repository.OpdRepository
	DB                        *sql.DB
}

func NewUsulanMusrebangServiceImpl(usulanMusrebangRepository repository.UsulanMusrebangRepository, rencanaKinerjaRepository repository.RencanaKinerjaRepository, opdRepository repository.OpdRepository, DB *sql.DB) *UsulanMusrebangServiceImpl {
	return &UsulanMusrebangServiceImpl{
		usulanMusrebangRepository: usulanMusrebangRepository,
		rencanaKinerjaRepository:  rencanaKinerjaRepository,
		opdRepository:             opdRepository,
		DB:                        DB,
	}
}

func (service *UsulanMusrebangServiceImpl) Create(ctx context.Context, request usulan.UsulanMusrebangCreateRequest) (usulan.UsulanMusrebangResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return usulan.UsulanMusrebangResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	randomDigits := fmt.Sprintf("%05d", uuid.New().ID()%100000)
	uuId := fmt.Sprintf("USU-MUS-%s", randomDigits)

	// Konversi request ke domain.UsulanMusrebang
	domainUsulanMusrebang := domain.UsulanMusrebang{
		Id:      uuId,
		Usulan:  request.Usulan,
		Alamat:  request.Alamat,
		Uraian:  request.Uraian,
		Tahun:   request.Tahun,
		RekinId: request.RekinId,
		KodeOpd: request.KodeOpd,
		Status:  "belum_diambil",
	}

	usulanMusrebang, err := service.usulanMusrebangRepository.Create(ctx, tx, domainUsulanMusrebang)
	if err != nil {
		return usulan.UsulanMusrebangResponse{}, err
	}

	return helper.ToUsulanMusrebangResponse(usulanMusrebang), nil
}

func (service *UsulanMusrebangServiceImpl) Update(ctx context.Context, request usulan.UsulanMusrebangUpdateRequest) (usulan.UsulanMusrebangResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return usulan.UsulanMusrebangResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	// Cek apakah usulan dengan ID yang diberikan ada
	existingUsulan, err := service.usulanMusrebangRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return usulan.UsulanMusrebangResponse{}, fmt.Errorf("usulan musrebang tidak ditemukan: %v", err)
	}

	// Update data usulan
	existingUsulan.Usulan = request.Usulan
	existingUsulan.Alamat = request.Alamat
	existingUsulan.Uraian = request.Uraian
	existingUsulan.Tahun = request.Tahun
	existingUsulan.KodeOpd = request.KodeOpd
	existingUsulan.Status = request.Status

	updatedUsulan, err := service.usulanMusrebangRepository.Update(ctx, tx, existingUsulan)
	if err != nil {
		return usulan.UsulanMusrebangResponse{}, err
	}

	// Menggunakan helper untuk menghasilkan respons
	return helper.ToUsulanMusrebangResponse(updatedUsulan), nil
}

func (service *UsulanMusrebangServiceImpl) FindById(ctx context.Context, idUsulan string) (usulan.UsulanMusrebangResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return usulan.UsulanMusrebangResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	usulanMusrebang, err := service.usulanMusrebangRepository.FindById(ctx, tx, idUsulan)
	if err != nil {
		return usulan.UsulanMusrebangResponse{}, err
	}

	opd, err := service.opdRepository.FindByKodeOpd(ctx, tx, usulanMusrebang.KodeOpd)
	if err != nil {
		return usulan.UsulanMusrebangResponse{}, err
	}
	usulanMusrebang.NamaOpd = opd.NamaOpd

	return helper.ToUsulanMusrebangResponse(usulanMusrebang), nil
}

func (service *UsulanMusrebangServiceImpl) FindAll(ctx context.Context, kodeOpd *string, is_active *bool, rekinId *string, status *string) ([]usulan.UsulanMusrebangResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []usulan.UsulanMusrebangResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	usulanMusrebang, err := service.usulanMusrebangRepository.FindAll(ctx, tx, kodeOpd, is_active, rekinId, status)
	if err != nil {
		return []usulan.UsulanMusrebangResponse{}, err
	}

	// Buat map untuk menyimpan data OPD yang sudah diambil
	opdCache := make(map[string]string)

	// Isi nama OPD untuk setiap usulan
	for i := range usulanMusrebang {
		kodeOpd := usulanMusrebang[i].KodeOpd

		// Cek apakah nama OPD sudah ada di cache
		if namaOpd, exists := opdCache[kodeOpd]; exists {
			usulanMusrebang[i].NamaOpd = namaOpd
		} else {

			opd, err := service.opdRepository.FindByKodeOpd(ctx, tx, kodeOpd)
			if err != nil {
				continue
			}
			// Simpan ke cache
			opdCache[kodeOpd] = opd.NamaOpd
			usulanMusrebang[i].NamaOpd = opd.NamaOpd
		}
	}

	// Konversi ke response setelah nama OPD diisi
	usulanMusrebangResponses := helper.ToUsulanMusrebangResponses(usulanMusrebang)
	return usulanMusrebangResponses, nil
}

func (service *UsulanMusrebangServiceImpl) Delete(ctx context.Context, idUsulan string) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return fmt.Errorf("gagal memulai transaksi: %v", err)
	}
	defer helper.CommitOrRollback(tx)

	// Cari usulan berdasarkan ID
	_, err = service.usulanMusrebangRepository.FindById(ctx, tx, idUsulan)
	if err != nil {
		return fmt.Errorf("usulan tidak ditemukan: %v", err)
	}

	// Jika usulan ditemukan, lanjutkan dengan penghapusan
	err = service.usulanMusrebangRepository.Delete(ctx, tx, idUsulan)
	if err != nil {
		return fmt.Errorf("gagal menghapus usulan: %v", err)
	}

	return nil
}

func (service *UsulanMusrebangServiceImpl) CreateRekin(ctx context.Context, request usulan.UsulanMusrebangCreateRekinRequest) ([]usulan.UsulanMusrebangResponse, error) {
	// Konversi single ID ke array
	idUsulanArray := []string{request.IdUsulan}

	tx, err := service.DB.Begin()
	if err != nil {
		return nil, fmt.Errorf("gagal memulai transaksi: %v", err)
	}
	defer helper.CommitOrRollback(tx)

	// Cek apakah rencana kinerja dengan ID yang diberikan ada
	_, err = service.rencanaKinerjaRepository.FindById(ctx, tx, request.RekinId, "", "")
	if err != nil {
		return nil, fmt.Errorf("rencana kinerja dengan id %s tidak ditemukan: %v", request.RekinId, err)
	}

	var updatedUsulans []domain.UsulanMusrebang

	// Proses setiap ID usulan
	for _, idUsulan := range idUsulanArray {
		// Cek apakah usulan dengan ID yang diberikan ada
		existingUsulan, err := service.usulanMusrebangRepository.FindById(ctx, tx, idUsulan)
		if err != nil {
			return nil, fmt.Errorf("usulan musrebang dengan id %s tidak ditemukan: %v", idUsulan, err)
		}

		// Cek apakah usulan sudah memiliki rekin_id
		if existingUsulan.RekinId != "" {
			return nil, fmt.Errorf("usulan musrebang dengan id %s sudah memiliki rencana kinerja", idUsulan)
		}

		// Update rekin_id dan status
		err = service.usulanMusrebangRepository.CreateRekin(ctx, tx, idUsulan, request.RekinId)
		if err != nil {
			return nil, fmt.Errorf("gagal mengupdate rekin untuk usulan %s: %v", idUsulan, err)
		}

		// Ambil data usulan yang sudah diupdate
		updatedUsulan, err := service.usulanMusrebangRepository.FindById(ctx, tx, idUsulan)
		if err != nil {
			return nil, fmt.Errorf("gagal mengambil data usulan yang diupdate: %v", err)
		}

		// Ambil data OPD
		opd, err := service.opdRepository.FindByKodeOpd(ctx, tx, updatedUsulan.KodeOpd)
		if err == nil { // Hanya set jika berhasil mendapatkan data OPD
			updatedUsulan.NamaOpd = opd.NamaOpd
		}

		updatedUsulans = append(updatedUsulans, updatedUsulan)
	}

	// Konversi ke response
	responses := helper.ToUsulanMusrebangResponses(updatedUsulans)
	return responses, nil
}

func (service *UsulanMusrebangServiceImpl) DeleteUsulanTerpilih(ctx context.Context, idUsulan string) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return fmt.Errorf("gagal memulai transaksi: %v", err)
	}
	defer helper.CommitOrRollback(tx)

	err = service.usulanMusrebangRepository.DeleteUsulanTerpilih(ctx, tx, idUsulan)
	if err != nil {
		return fmt.Errorf("gagal menghapus usulan terpilih: %v", err)
	}

	return nil
}
