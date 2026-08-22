package service

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/domain"
	"ekak_kab_sleman/model/web/permasalahan"
	"ekak_kab_sleman/repository"
)

type PermasalahanRekinServiceImpl struct {
	PermasalahanRekinRepository repository.PermasalahanRekinRepository
	DB                          *sql.DB
}

func NewPermasalahanRekinServiceImpl(permasalahanRekinRepository repository.PermasalahanRekinRepository, DB *sql.DB) *PermasalahanRekinServiceImpl {
	return &PermasalahanRekinServiceImpl{
		PermasalahanRekinRepository: permasalahanRekinRepository,
		DB:                          DB,
	}
}

func (service *PermasalahanRekinServiceImpl) Create(ctx context.Context, request permasalahan.PermasalahanRekinCreateRequest) (permasalahan.PermasalahanRekinResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return permasalahan.PermasalahanRekinResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	permasalahanDomain := domain.PermasalahanRekin{
		Id:                helper.GenerateRandomNumber(6),
		RekinId:           request.RekinId,
		Permasalahan:      request.Permasalahan,
		PenyebabInternal:  helper.EmptyStringIfNull(request.PenyebabInternal),
		PenyebabEksternal: helper.EmptyStringIfNull(request.PenyebabEksternal),
		JenisPermasalahan: request.JenisPermasalahan,
	}

	result, err := service.PermasalahanRekinRepository.Create(ctx, tx, permasalahanDomain)
	if err != nil {
		return permasalahan.PermasalahanRekinResponse{}, err
	}

	return permasalahan.PermasalahanRekinResponse{
		Id:                result.Id,
		RekinId:           result.RekinId,
		Permasalahan:      result.Permasalahan,
		PenyebabInternal:  helper.EmptyStringIfNull(result.PenyebabInternal),
		PenyebabEksternal: helper.EmptyStringIfNull(result.PenyebabEksternal),
		JenisPermasalahan: result.JenisPermasalahan,
	}, nil
}

func (service *PermasalahanRekinServiceImpl) Update(ctx context.Context, request permasalahan.PermasalahanRekinUpdateRequest) (permasalahan.PermasalahanRekinResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return permasalahan.PermasalahanRekinResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	permasalahanDomain := domain.PermasalahanRekin{
		Id:                request.Id,
		Permasalahan:      request.Permasalahan,
		PenyebabInternal:  request.PenyebabInternal,
		PenyebabEksternal: request.PenyebabEksternal,
	}

	result, err := service.PermasalahanRekinRepository.Update(ctx, tx, permasalahanDomain)
	if err != nil {
		return permasalahan.PermasalahanRekinResponse{}, err
	}

	return permasalahan.PermasalahanRekinResponse{
		Id:                result.Id,
		RekinId:           result.RekinId,
		Permasalahan:      result.Permasalahan,
		PenyebabInternal:  result.PenyebabInternal,
		PenyebabEksternal: result.PenyebabEksternal,
		JenisPermasalahan: result.JenisPermasalahan,
	}, nil
}

func (service *PermasalahanRekinServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.PermasalahanRekinRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *PermasalahanRekinServiceImpl) FindAll(ctx context.Context, rekinId *string) ([]permasalahan.PermasalahanRekinResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	permasalahanList, err := service.PermasalahanRekinRepository.FindAll(ctx, tx, rekinId)
	if err != nil {
		return nil, err
	}

	var responses []permasalahan.PermasalahanRekinResponse
	for _, p := range permasalahanList {
		responses = append(responses, permasalahan.PermasalahanRekinResponse{
			Id:                p.Id,
			RekinId:           p.RekinId,
			Permasalahan:      p.Permasalahan,
			PenyebabInternal:  p.PenyebabInternal,
			PenyebabEksternal: p.PenyebabEksternal,
			JenisPermasalahan: p.JenisPermasalahan,
		})
	}

	return responses, nil
}

func (service *PermasalahanRekinServiceImpl) FindById(ctx context.Context, id int) (permasalahan.PermasalahanRekinResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return permasalahan.PermasalahanRekinResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	result, err := service.PermasalahanRekinRepository.FindById(ctx, tx, id)
	if err != nil {
		return permasalahan.PermasalahanRekinResponse{}, err
	}

	return permasalahan.PermasalahanRekinResponse{
		Id:                result.Id,
		RekinId:           result.RekinId,
		Permasalahan:      result.Permasalahan,
		PenyebabInternal:  result.PenyebabInternal,
		PenyebabEksternal: result.PenyebabEksternal,
		JenisPermasalahan: result.JenisPermasalahan,
	}, nil
}
