package service

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/domain/domainmaster"
	"ekak_kab_sleman/model/web/jabatan"
	"ekak_kab_sleman/repository"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type JabatanServiceImpl struct {
	jabatanRepository repository.JabatanRepository
	opdRepository     repository.OpdRepository
	DB                *sql.DB
}

func NewJabatanServiceImpl(jabatanRepository repository.JabatanRepository, opdRepository repository.OpdRepository, DB *sql.DB) *JabatanServiceImpl {
	return &JabatanServiceImpl{
		jabatanRepository: jabatanRepository,
		opdRepository:     opdRepository,
		DB:                DB,
	}
}

func (service *JabatanServiceImpl) Create(ctx context.Context, request jabatan.JabatanCreateRequest) (jabatan.JabatanResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	opd, err := service.opdRepository.FindByKodeOpd(ctx, tx, request.KodeOpd)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Kode OPD %s tidak ditemukan", request.KodeOpd)
			return jabatan.JabatanResponse{}, fmt.Errorf("kode OPD %s tidak ditemukan", request.KodeOpd)
		}
		log.Printf("Gagal memeriksa kode OPD: %v", err)
		return jabatan.JabatanResponse{}, fmt.Errorf("gagal memeriksa kode OPD: %v", err)
	}

	if opd.KodeOpd == "" {
		log.Printf("Kode OPD %s tidak valid", request.KodeOpd)
		return jabatan.JabatanResponse{}, fmt.Errorf("kode OPD %s tidak valid", request.KodeOpd)
	}

	uuid := uuid.New().String()
	jabatanId := fmt.Sprintf("JBTN-%v", uuid[:4])

	jabatanDomain := domainmaster.Jabatan{
		Id:           jabatanId,
		KodeJabatan:  request.KodeJabatan,
		NamaJabatan:  request.NamaJabatan,
		KelasJabatan: helper.EmptyStringIfNull(request.KelasJabatan),
		JenisJabatan: helper.EmptyStringIfNull(request.JenisJabatan),
		NilaiJabatan: request.NilaiJabatan,
		KodeOpd:      request.KodeOpd,
		IndexJabatan: request.IndexJabatan,
		Tahun:        helper.EmptyStringIfNull(request.Tahun),
		Esselon:      helper.EmptyStringIfNull(request.Esselon),
	}

	jabatanDomain.NamaOpd = opd.NamaOpd

	jabatan := service.jabatanRepository.Create(ctx, tx, jabatanDomain)
	return helper.ToJabatanResponse(jabatan), nil
}

func (service *JabatanServiceImpl) Update(ctx context.Context, request jabatan.JabatanUpdateRequest) (jabatan.JabatanResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	opd, err := service.opdRepository.FindByKodeOpd(ctx, tx, request.KodeOpd)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Kode OPD %s tidak ditemukan", request.KodeOpd)
			return jabatan.JabatanResponse{}, fmt.Errorf("kode OPD %s tidak ditemukan", request.KodeOpd)
		}
		log.Printf("Gagal memeriksa kode OPD: %v", err)
		return jabatan.JabatanResponse{}, fmt.Errorf("gagal memeriksa kode OPD: %v", err)
	}

	if opd.KodeOpd == "" {
		log.Printf("Kode OPD %s tidak valid", request.KodeOpd)
		return jabatan.JabatanResponse{}, fmt.Errorf("kode OPD %s tidak valid", request.KodeOpd)
	}

	jabatanDomain := domainmaster.Jabatan{
		Id:           request.Id,
		KodeJabatan:  request.KodeJabatan,
		NamaJabatan:  request.NamaJabatan,
		KelasJabatan: helper.EmptyStringIfNull(request.KelasJabatan),
		JenisJabatan: helper.EmptyStringIfNull(request.JenisJabatan),
		NilaiJabatan: request.NilaiJabatan,
		KodeOpd:      request.KodeOpd,
		IndexJabatan: request.IndexJabatan,
		Tahun:        helper.EmptyStringIfNull(request.Tahun),
		Esselon:      helper.EmptyStringIfNull(request.Esselon),
	}

	jabatanDomain.NamaOpd = opd.NamaOpd

	jabatan := service.jabatanRepository.Update(ctx, tx, jabatanDomain)
	return helper.ToJabatanResponse(jabatan), nil
}

func (service *JabatanServiceImpl) Delete(ctx context.Context, id string) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.jabatanRepository.Delete(ctx, tx, id)
	helper.PanicIfError(err)
	return nil
}

func (service *JabatanServiceImpl) FindById(ctx context.Context, id string) (jabatan.JabatanResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return jabatan.JabatanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	jabatans, err := service.jabatanRepository.FindById(ctx, tx, id)
	if err != nil {
		return jabatan.JabatanResponse{}, err
	}

	opd, err := service.opdRepository.FindByKodeOpd(ctx, tx, jabatans.KodeOpd)
	if err != nil {
		return jabatan.JabatanResponse{}, err
	}

	jabatans.NamaOpd = opd.NamaOpd

	return helper.ToJabatanResponse(jabatans), nil
}
func (service *JabatanServiceImpl) FindAll(ctx context.Context, kodeOpd string, tahun string) ([]jabatan.JabatanResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	jabatans, err := service.jabatanRepository.FindAll(ctx, tx, kodeOpd, tahun)
	if err != nil {
		return nil, err
	}
	opd, err := service.opdRepository.FindByKodeOpd(ctx, tx, kodeOpd)
	if err != nil {
		log.Printf("Gagal mengambil data OPD: %v", err)
		return nil, fmt.Errorf("gagal mengambil data OPD: %v", err)
	}

	// Set nama OPD untuk setiap jabatan
	for i := range jabatans {
		jabatans[i].NamaOpd = opd.NamaOpd
	}

	return helper.ToJabatanResponses(jabatans), nil
}
