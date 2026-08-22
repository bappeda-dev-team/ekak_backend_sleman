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

type UsulanPokokPikiranServiceImpl struct {
	UsulanPokokPikiranRepository repository.UsulanPokokPikiranRepository
	RencanaKinerjaRepository     repository.RencanaKinerjaRepository
	OpdRepository                repository.OpdRepository
	DB                           *sql.DB
}

func NewUsulanPokokPikiranServiceImpl(usulanPokokPikiranRepository repository.UsulanPokokPikiranRepository, rencanaKinerjaRepository repository.RencanaKinerjaRepository, opdRepository repository.OpdRepository, DB *sql.DB) *UsulanPokokPikiranServiceImpl {
	return &UsulanPokokPikiranServiceImpl{
		UsulanPokokPikiranRepository: usulanPokokPikiranRepository,
		RencanaKinerjaRepository:     rencanaKinerjaRepository,
		OpdRepository:                opdRepository,
		DB:                           DB,
	}
}

func (service *UsulanPokokPikiranServiceImpl) Create(ctx context.Context, request usulan.UsulanPokokPikiranCreateRequest) (usulan.UsulanPokokPikiranResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return usulan.UsulanPokokPikiranResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	randomDigits := fmt.Sprintf("%05d", uuid.New().ID()%100000)
	uuId := fmt.Sprintf("USU-POKIR-%s", randomDigits)

	domainUsulanPokokPikiran := domain.UsulanPokokPikiran{
		Id:      uuId,
		Usulan:  request.Usulan,
		Alamat:  request.Alamat,
		Uraian:  request.Uraian,
		Tahun:   request.Tahun,
		RekinId: request.RekinId,
		KodeOpd: request.KodeOpd,
		Status:  "belum_diambil",
	}

	usulanPokokPikiran, err := service.UsulanPokokPikiranRepository.Create(ctx, tx, domainUsulanPokokPikiran)
	if err != nil {
		return usulan.UsulanPokokPikiranResponse{}, err
	}

	return helper.ToUsulanPokokPikiranResponse(usulanPokokPikiran), nil
}

func (service *UsulanPokokPikiranServiceImpl) Update(ctx context.Context, request usulan.UsulanPokokPikiranUpdateRequest) (usulan.UsulanPokokPikiranResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return usulan.UsulanPokokPikiranResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	usulans, err := service.UsulanPokokPikiranRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return usulan.UsulanPokokPikiranResponse{}, err
	}

	usulans.Usulan = request.Usulan
	usulans.Alamat = request.Alamat
	usulans.Uraian = request.Uraian
	usulans.Tahun = request.Tahun
	usulans.KodeOpd = request.KodeOpd
	usulans.Status = request.Status

	updatedUsulan, err := service.UsulanPokokPikiranRepository.Update(ctx, tx, usulans)
	if err != nil {
		return usulan.UsulanPokokPikiranResponse{}, err
	}

	return helper.ToUsulanPokokPikiranResponse(updatedUsulan), nil
}

func (service *UsulanPokokPikiranServiceImpl) FindById(ctx context.Context, idUsulan string) (usulan.UsulanPokokPikiranResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return usulan.UsulanPokokPikiranResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	usulanPokokPikiran, err := service.UsulanPokokPikiranRepository.FindById(ctx, tx, idUsulan)
	if err != nil {
		return usulan.UsulanPokokPikiranResponse{}, err
	}

	return helper.ToUsulanPokokPikiranResponse(usulanPokokPikiran), nil
}

func (service *UsulanPokokPikiranServiceImpl) FindAll(ctx context.Context, kodeOpd *string, is_active *bool, rekinId *string, status *string) ([]usulan.UsulanPokokPikiranResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []usulan.UsulanPokokPikiranResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	usulanPokokPikiran, err := service.UsulanPokokPikiranRepository.FindAll(ctx, tx, kodeOpd, is_active, rekinId, status)
	if err != nil {
		return []usulan.UsulanPokokPikiranResponse{}, err
	}

	// Jika tidak ada data usulan, kembalikan array kosong
	if len(usulanPokokPikiran) == 0 {
		return []usulan.UsulanPokokPikiranResponse{}, nil
	}

	// Ambil data OPD untuk usulan pertama
	opd, err := service.OpdRepository.FindByKodeOpd(ctx, tx, usulanPokokPikiran[0].KodeOpd)
	if err != nil {
		return []usulan.UsulanPokokPikiranResponse{}, nil // Kembalikan array kosong jika OPD tidak ditemukan
	}

	// Set nama OPD untuk semua usulan dengan kode OPD yang sama
	for i := range usulanPokokPikiran {
		if usulanPokokPikiran[i].KodeOpd == usulanPokokPikiran[0].KodeOpd {
			usulanPokokPikiran[i].NamaOpd = opd.NamaOpd
		} else {
			// Jika ada kode OPD berbeda, ambil data OPD-nya
			opdLain, err := service.OpdRepository.FindByKodeOpd(ctx, tx, usulanPokokPikiran[i].KodeOpd)
			if err == nil { // Hanya set jika berhasil mendapatkan data OPD
				usulanPokokPikiran[i].NamaOpd = opdLain.NamaOpd
			}
		}
	}

	return helper.ToUsulanPokokPikiranResponses(usulanPokokPikiran), nil
}
func (service *UsulanPokokPikiranServiceImpl) Delete(ctx context.Context, idUsulan string) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.UsulanPokokPikiranRepository.FindById(ctx, tx, idUsulan)
	if err != nil {
		return fmt.Errorf("usulan dengan ID %s tidak ditemukan", idUsulan)
	}

	err = service.UsulanPokokPikiranRepository.Delete(ctx, tx, idUsulan)
	helper.PanicIfError(err)

	return nil
}

func (service *UsulanPokokPikiranServiceImpl) CreateRekin(ctx context.Context, request usulan.UsulanPokokPikiranCreateRekinRequest) ([]usulan.UsulanPokokPikiranResponse, error) {
	// Konversi single ID ke array
	idUsulanArray := []string{request.IdUsulan}

	tx, err := service.DB.Begin()
	if err != nil {
		return nil, fmt.Errorf("gagal memulai transaksi: %v", err)
	}
	defer helper.CommitOrRollback(tx)

	// Cek apakah rencana kinerja dengan ID yang diberikan ada
	_, err = service.RencanaKinerjaRepository.FindById(ctx, tx, request.RekinId, "", "")
	if err != nil {
		return nil, fmt.Errorf("rencana kinerja dengan id %s tidak ditemukan: %v", request.RekinId, err)
	}

	var updatedUsulans []domain.UsulanPokokPikiran

	// Proses setiap ID usulan
	for _, idUsulan := range idUsulanArray {
		// Cek apakah usulan dengan ID yang diberikan ada
		existingUsulan, err := service.UsulanPokokPikiranRepository.FindById(ctx, tx, idUsulan)
		if err != nil {
			return nil, fmt.Errorf("usulan pokok pikiran dengan id %s tidak ditemukan: %v", idUsulan, err)
		}

		// Cek apakah usulan sudah memiliki rekin_id
		if existingUsulan.RekinId != "" {
			return nil, fmt.Errorf("usulan pokok pikiran dengan id %s sudah memiliki rencana kinerja", idUsulan)
		}

		// Update rekin_id dan status
		err = service.UsulanPokokPikiranRepository.CreateRekin(ctx, tx, idUsulan, request.RekinId)
		if err != nil {
			return nil, fmt.Errorf("gagal mengupdate rekin untuk usulan %s: %v", idUsulan, err)
		}

		// Ambil data usulan yang sudah diupdate
		updatedUsulan, err := service.UsulanPokokPikiranRepository.FindById(ctx, tx, idUsulan)
		if err != nil {
			return nil, fmt.Errorf("gagal mengambil data usulan yang diupdate: %v", err)
		}

		// Ambil data OPD
		opd, err := service.OpdRepository.FindByKodeOpd(ctx, tx, updatedUsulan.KodeOpd)
		if err == nil { // Hanya set jika berhasil mendapatkan data OPD
			updatedUsulan.NamaOpd = opd.NamaOpd
		}

		updatedUsulans = append(updatedUsulans, updatedUsulan)
	}

	// Konversi ke response
	responses := helper.ToUsulanPokokPikiranResponses(updatedUsulans)
	return responses, nil
}

func (service *UsulanPokokPikiranServiceImpl) DeleteUsulanTerpilih(ctx context.Context, idUsulan string) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return fmt.Errorf("gagal memulai transaksi: %v", err)
	}
	defer helper.CommitOrRollback(tx)

	err = service.UsulanPokokPikiranRepository.DeleteUsulanTerpilih(ctx, tx, idUsulan)
	if err != nil {
		return fmt.Errorf("gagal menghapus usulan terpilih: %v", err)
	}

	return nil
}
