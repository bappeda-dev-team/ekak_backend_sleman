package service

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/domain"
	"ekak_kab_sleman/model/web/inovasi"
	"ekak_kab_sleman/repository"
	"fmt"

	"github.com/google/uuid"
)

type InovasiServiceImpl struct {
	InovasiRepository repository.InovasiRepository
	DB                *sql.DB
}

func NewInovasiServiceImpl(inovasiRepository repository.InovasiRepository, DB *sql.DB) *InovasiServiceImpl {
	return &InovasiServiceImpl{
		InovasiRepository: inovasiRepository,
		DB:                DB,
	}
}

func (service *InovasiServiceImpl) Create(ctx context.Context, request inovasi.InovasiCreateRequest) (inovasi.InovasiResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return inovasi.InovasiResponse{}, fmt.Errorf("gagal memulai transaksi: %v", err)
	}
	defer helper.CommitOrRollback(tx)

	randomDigits := fmt.Sprintf("%05d", uuid.New().ID()%100000)
	uuId := fmt.Sprintf("INOV-REKIN-%s", randomDigits)

	domainInovasi := domain.Inovasi{
		Id:                    uuId,
		RekinId:               request.RekinId,
		KodeOpd:               request.KodeOpd,
		JudulInovasi:          request.JudulInovasi,
		JenisInovasi:          request.JenisInovasi,
		GambaranNilaiKebaruan: request.GambaranNilaiKebaruan,
	}

	inovasis, err := service.InovasiRepository.Create(ctx, tx, domainInovasi)
	if err != nil {
		return inovasi.InovasiResponse{}, err
	}

	response := helper.ToInovasiResponse(inovasis)
	return response, nil
}

func (service *InovasiServiceImpl) Update(ctx context.Context, request inovasi.InovasiUpdateRequest) (inovasi.InovasiResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return inovasi.InovasiResponse{}, fmt.Errorf("gagal memulai transaksi: %v", err)
	}
	defer helper.CommitOrRollback(tx)

	domainInovasi := domain.Inovasi{
		Id:                    request.Id,
		RekinId:               request.RekinId,
		JudulInovasi:          request.JudulInovasi,
		JenisInovasi:          request.JenisInovasi,
		GambaranNilaiKebaruan: request.GambaranNilaiKebaruan,
	}

	inovasis, err := service.InovasiRepository.Update(ctx, tx, domainInovasi)
	if err != nil {
		return inovasi.InovasiResponse{}, err
	}

	response := helper.ToInovasiResponse(inovasis)
	return response, nil
}

func (service *InovasiServiceImpl) FindAll(ctx context.Context, rekinId string) ([]inovasi.InovasiResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []inovasi.InovasiResponse{}, fmt.Errorf("gagal memulai transaksi: %v", err)
	}
	defer helper.CommitOrRollback(tx)

	inovasis, err := service.InovasiRepository.FindAll(ctx, tx, rekinId)
	if err != nil {
		return []inovasi.InovasiResponse{}, err
	}

	response := helper.ToInovasiResponses(inovasis)
	return response, nil
}

func (service *InovasiServiceImpl) FindById(ctx context.Context, inovasiId string) (inovasi.InovasiResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return inovasi.InovasiResponse{}, fmt.Errorf("gagal memulai transaksi: %v", err)
	}
	defer helper.CommitOrRollback(tx)

	inovasis, err := service.InovasiRepository.FindById(ctx, tx, inovasiId)
	if err != nil {
		if err == sql.ErrNoRows {
			return inovasi.InovasiResponse{}, fmt.Errorf("id tidak ditemukan")
		}
		return inovasi.InovasiResponse{}, err
	}

	response := helper.ToInovasiResponse(inovasis)
	return response, nil
}

func (service *InovasiServiceImpl) Delete(ctx context.Context, inovasiId string) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return fmt.Errorf("gagal memulai transaksi: %v", err)
	}
	defer helper.CommitOrRollback(tx)

	// Periksa apakah inovasi dengan ID tersebut ada
	_, err = service.InovasiRepository.FindById(ctx, tx, inovasiId)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("id tidak ditemukan")
		}
		return err
	}

	err = service.InovasiRepository.Delete(ctx, tx, inovasiId)
	if err != nil {
		return err
	}
	return nil
}
