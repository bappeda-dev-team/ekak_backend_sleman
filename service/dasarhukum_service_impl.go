package service

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/domain"
	"ekak_kab_sleman/model/web/dasarhukum"
	"ekak_kab_sleman/repository"
	"fmt"

	"github.com/google/uuid"
)

type DasarHukumServiceImpl struct {
	DasarHukumRepository repository.DasarHukumRepository
	DB                   *sql.DB
}

func NewDasarHukumServiceImpl(dasarHukumRepository repository.DasarHukumRepository, DB *sql.DB) *DasarHukumServiceImpl {
	return &DasarHukumServiceImpl{
		DasarHukumRepository: dasarHukumRepository,
		DB:                   DB,
	}
}

func (service *DasarHukumServiceImpl) Create(ctx context.Context, request dasarhukum.DasarHukumCreateRequest) (dasarhukum.DasarHukumResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return dasarhukum.DasarHukumResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	// Membuat UUID dengan format yang diinginkan
	randomDigits := fmt.Sprintf("%05d", uuid.New().ID()%100000)
	uuId := fmt.Sprintf("DASHU-REKIN-%s", randomDigits)

	// Mendapatkan urutan terakhir
	lastUrutan, err := service.DasarHukumRepository.GetLastUrutanByRekinId(ctx, tx, request.RekinId)
	if err != nil {
		return dasarhukum.DasarHukumResponse{}, err
	}

	newUrutan := 1
	if lastUrutan > 0 {
		newUrutan = lastUrutan + 1
	}

	dasarHukum := domain.DasarHukum{
		Id:               uuId,
		RekinId:          request.RekinId,
		KodeOpd:          request.KodeOpd,
		Urutan:           newUrutan,
		PeraturanTerkait: request.PeraturanTerkait,
		Uraian:           request.Uraian,
	}

	dasarHukum, err = service.DasarHukumRepository.Create(ctx, tx, dasarHukum)
	if err != nil {
		return dasarhukum.DasarHukumResponse{}, err
	}

	return helper.ToDasarHukumResponse(dasarHukum), nil
}

func (service *DasarHukumServiceImpl) Update(ctx context.Context, request dasarhukum.DasarHukumUpdateRequest) (dasarhukum.DasarHukumResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return dasarhukum.DasarHukumResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	dasarHukum, err := service.DasarHukumRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return dasarhukum.DasarHukumResponse{}, err
	}

	dasarHukum.PeraturanTerkait = request.PeraturanTerkait
	dasarHukum.Uraian = request.Uraian

	updatedDasarHukum, err := service.DasarHukumRepository.Update(ctx, tx, dasarHukum)
	if err != nil {
		return dasarhukum.DasarHukumResponse{}, err
	}

	return helper.ToDasarHukumResponse(updatedDasarHukum), nil
}

func (service *DasarHukumServiceImpl) Delete(ctx context.Context, id string) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.DasarHukumRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *DasarHukumServiceImpl) FindAll(ctx context.Context, rekinId string) ([]dasarhukum.DasarHukumResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dasarHukums, err := service.DasarHukumRepository.FindAll(ctx, tx, rekinId)
	if err != nil {
		return []dasarhukum.DasarHukumResponse{}, err
	}

	return helper.ToDasarHukumResponses(dasarHukums), nil
}

func (service *DasarHukumServiceImpl) FindById(ctx context.Context, id string) (dasarhukum.DasarHukumResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dasarHukum, err := service.DasarHukumRepository.FindById(ctx, tx, id)
	if err != nil {
		return dasarhukum.DasarHukumResponse{}, err
	}

	return helper.ToDasarHukumResponse(dasarHukum), nil
}
