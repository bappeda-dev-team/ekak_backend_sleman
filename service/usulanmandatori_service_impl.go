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

type UsulanMandatoriServiceImpl struct {
	usulanMandatoriRepository repository.UsulanMandatoriRepository
	pegawaiRepository         repository.PegawaiRepository
	opdRepository             repository.OpdRepository
	DB                        *sql.DB
}

func NewUsulanMandatoriServiceImpl(usulanMandatoriRepository repository.UsulanMandatoriRepository, pegawaiRepository repository.PegawaiRepository, opdRepository repository.OpdRepository, DB *sql.DB) *UsulanMandatoriServiceImpl {
	return &UsulanMandatoriServiceImpl{
		usulanMandatoriRepository: usulanMandatoriRepository,
		pegawaiRepository:         pegawaiRepository,
		opdRepository:             opdRepository,
		DB:                        DB,
	}
}

func (service *UsulanMandatoriServiceImpl) Create(ctx context.Context, request usulan.UsulanMandatoriCreateRequest) (usulan.UsulanMandatoriResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return usulan.UsulanMandatoriResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	randomDigits := fmt.Sprintf("%05d", uuid.New().ID()%100000)
	uuId := fmt.Sprintf("USU-MAND-%s", randomDigits)

	pegawai, err := service.pegawaiRepository.FindByNip(ctx, tx, request.PegawaiId)
	if err != nil {
		if err == sql.ErrNoRows {
			return usulan.UsulanMandatoriResponse{}, fmt.Errorf("pegawai dengan NIP %s tidak ditemukan dalam database", request.PegawaiId)
		}
		return usulan.UsulanMandatoriResponse{}, fmt.Errorf("terjadi kesalahan saat mencari data pegawai: %v", err)
	}

	opd, err := service.opdRepository.FindByKodeOpd(ctx, tx, request.KodeOpd)
	if err != nil {
		if err == sql.ErrNoRows {
			return usulan.UsulanMandatoriResponse{}, fmt.Errorf("OPD dengan kode %s tidak ditemukan dalam database", request.KodeOpd)
		}
		return usulan.UsulanMandatoriResponse{}, fmt.Errorf("terjadi kesalahan saat mencari data OPD: %v", err)
	}

	domainUsulanMandatori := domain.UsulanMandatori{
		Id:               uuId,
		Usulan:           request.Usulan,
		PeraturanTerkait: request.PeraturanTerkait,
		Uraian:           request.Uraian,
		Tahun:            request.Tahun,
		RekinId:          request.RekinId,
		PegawaiId:        pegawai.Nip,
		KodeOpd:          opd.KodeOpd,
		Status:           request.Status,
	}

	usulanMandatori, err := service.usulanMandatoriRepository.Create(ctx, tx, domainUsulanMandatori)
	if err != nil {
		return usulan.UsulanMandatoriResponse{}, err
	}

	return helper.ToUsulanMandatoriResponse(usulanMandatori), nil
}

func (service *UsulanMandatoriServiceImpl) Update(ctx context.Context, request usulan.UsulanMandatoriUpdateRequest) (usulan.UsulanMandatoriResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return usulan.UsulanMandatoriResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	existingUsulan, err := service.usulanMandatoriRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return usulan.UsulanMandatoriResponse{}, err
	}

	pegawai, err := service.pegawaiRepository.FindByNip(ctx, tx, request.PegawaiId)
	if err != nil {
		return usulan.UsulanMandatoriResponse{}, err
	}

	existingUsulan.Usulan = request.Usulan
	existingUsulan.PeraturanTerkait = request.PeraturanTerkait
	existingUsulan.Uraian = request.Uraian
	existingUsulan.Tahun = request.Tahun
	existingUsulan.PegawaiId = pegawai.Nip
	existingUsulan.KodeOpd = request.KodeOpd
	existingUsulan.Status = request.Status

	updatedUsulan, err := service.usulanMandatoriRepository.Update(ctx, tx, existingUsulan)
	if err != nil {
		return usulan.UsulanMandatoriResponse{}, err
	}

	return helper.ToUsulanMandatoriResponse(updatedUsulan), nil
}

func (service *UsulanMandatoriServiceImpl) FindById(ctx context.Context, idUsulan string) (usulan.UsulanMandatoriResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return usulan.UsulanMandatoriResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	usulanMandatori, err := service.usulanMandatoriRepository.FindById(ctx, tx, idUsulan)
	if err != nil {
		return usulan.UsulanMandatoriResponse{}, err
	}

	return helper.ToUsulanMandatoriResponse(usulanMandatori), nil
}

func (service *UsulanMandatoriServiceImpl) FindAll(ctx context.Context, kodeOpd *string, pegawaiId *string, isActive *bool, rekinId *string) ([]usulan.UsulanMandatoriResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []usulan.UsulanMandatoriResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	usulanMandatoris, err := service.usulanMandatoriRepository.FindAll(ctx, tx, kodeOpd, pegawaiId, isActive, rekinId)
	if err != nil {
		return []usulan.UsulanMandatoriResponse{}, err
	}

	// Jika tidak ada data, langsung kembalikan slice kosong
	if len(usulanMandatoris) == 0 {
		return []usulan.UsulanMandatoriResponse{}, nil
	}

	// Hanya cari data pegawai jika ada usulan mandatori
	pegawai, err := service.pegawaiRepository.FindByNip(ctx, tx, usulanMandatoris[0].PegawaiId)
	if err != nil {
		return []usulan.UsulanMandatoriResponse{}, err
	}

	usulanMandatoris[0].NamaPegawai = pegawai.NamaPegawai

	return helper.ToUsulanMandatoriResponses(usulanMandatoris), nil
}

func (service *UsulanMandatoriServiceImpl) Delete(ctx context.Context, idUsulan string) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.usulanMandatoriRepository.FindById(ctx, tx, idUsulan)
	if err != nil {
		return fmt.Errorf("usulan dengan ID %s tidak ditemukan", idUsulan)
	}

	err = service.usulanMandatoriRepository.Delete(ctx, tx, idUsulan)
	helper.PanicIfError(err)

	return nil
}
