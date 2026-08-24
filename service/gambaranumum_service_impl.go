package service

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/domain"
	"ekak_kab_sleman/model/web/gambaranumum"
	"ekak_kab_sleman/repository"
	"fmt"

	"github.com/google/uuid"
)

type GambaranUmumServiceImpl struct {
	gambaranUmumRepository repository.GambaranUmumRepository
	DB                     *sql.DB
}

func NewGambaranUmumServiceImpl(gambaranUmumRepository repository.GambaranUmumRepository, DB *sql.DB) *GambaranUmumServiceImpl {
	return &GambaranUmumServiceImpl{
		gambaranUmumRepository: gambaranUmumRepository,
		DB:                     DB,
	}
}

func (service *GambaranUmumServiceImpl) Create(ctx context.Context, request gambaranumum.GambaranUmumCreateRequest) (gambaranumum.GambaranUmumResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return gambaranumum.GambaranUmumResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	// Membuat UUID dengan format yang diinginkan
	randomDigits := fmt.Sprintf("%05d", uuid.New().ID()%100000)
	uuId := fmt.Sprintf("GMBRUMUM-REKIN-%s", randomDigits)

	// Mendapatkan urutan terakhir
	lastUrutan, err := service.gambaranUmumRepository.GetLastUrutanByRekinId(ctx, tx, request.RekinId)
	if err != nil {
		return gambaranumum.GambaranUmumResponse{}, err
	}

	// Menambahkan 1 ke urutan terakhir
	newUrutan := lastUrutan + 1

	gambaranUmum := domain.GambaranUmum{
		Id:           uuId,
		RekinId:      request.RekinId,
		KodeOpd:      request.KodeOpd,
		Urutan:       newUrutan,
		GambaranUmum: request.GambaranUmum,
	}

	gambaranUmum, err = service.gambaranUmumRepository.Create(ctx, tx, gambaranUmum)
	if err != nil {
		return gambaranumum.GambaranUmumResponse{}, err
	}

	return helper.ToGambaranUmumResponse(gambaranUmum), nil
}

func (service *GambaranUmumServiceImpl) Update(ctx context.Context, request gambaranumum.GambaranUmumUpdateRequest) (gambaranumum.GambaranUmumResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	gambaranUmum := domain.GambaranUmum{
		Id:           request.Id,
		Urutan:       request.Urutan,
		GambaranUmum: request.GambaranUmum,
	}

	gambaranUmum, err = service.gambaranUmumRepository.Update(ctx, tx, gambaranUmum)
	if err != nil {
		return gambaranumum.GambaranUmumResponse{}, err
	}

	return helper.ToGambaranUmumResponse(gambaranUmum), nil
}

func (service *GambaranUmumServiceImpl) Delete(ctx context.Context, id string) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	return service.gambaranUmumRepository.Delete(ctx, tx, id)

}

func (service *GambaranUmumServiceImpl) FindAll(ctx context.Context, rekinId string) ([]gambaranumum.GambaranUmumResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, fmt.Errorf("gagal memulai transaksi: %v", err)
	}
	defer tx.Rollback() // Hanya melakukan rollback jika belum di-commit

	gambaranUmums, err := service.gambaranUmumRepository.FindAll(ctx, tx, rekinId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("rekin dengan ID %s tidak ditemukan", rekinId)
		}
		return nil, fmt.Errorf("gagal mengambil data: %v", err)
	}

	// if len(gambaranUmums) == 0 {
	// 	return nil, fmt.Errorf("tidak ada gambaran umum untuk rekin dengan ID %s", rekinId)
	// }

	// Commit transaksi jika berhasil
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("gagal melakukan commit transaksi: %v", err)
	}

	return helper.ToGambaranUmumResponses(gambaranUmums), nil
}

func (service *GambaranUmumServiceImpl) FindById(ctx context.Context, id string) (gambaranumum.GambaranUmumResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	gambaranUmum, err := service.gambaranUmumRepository.FindById(ctx, tx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return gambaranumum.GambaranUmumResponse{}, fmt.Errorf("gambaran umum dengan ID %s tidak ditemukan", id)
		}
		return gambaranumum.GambaranUmumResponse{}, err
	}

	return helper.ToGambaranUmumResponse(gambaranUmum), nil
}
