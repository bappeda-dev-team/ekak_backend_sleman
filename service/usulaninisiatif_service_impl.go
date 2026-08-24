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

type UsulanInisiatifServiceImpl struct {
	UsulanInisiatifRepository repository.UsulanInisiatifRepository
	pegawaiRepository         repository.PegawaiRepository
	opdRepository             repository.OpdRepository
	DB                        *sql.DB
}

func NewUsulanInisiatifServiceImpl(usulanInisiatifRepository repository.UsulanInisiatifRepository, pegawaiRepository repository.PegawaiRepository, opdRepository repository.OpdRepository, DB *sql.DB) *UsulanInisiatifServiceImpl {
	return &UsulanInisiatifServiceImpl{
		UsulanInisiatifRepository: usulanInisiatifRepository,
		pegawaiRepository:         pegawaiRepository,
		opdRepository:             opdRepository,
		DB:                        DB,
	}
}

func (service *UsulanInisiatifServiceImpl) Create(ctx context.Context, request usulan.UsulanInisiatifCreateRequest) (usulan.UsulanInisiatifResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return usulan.UsulanInisiatifResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	randomDigits := fmt.Sprintf("%05d", uuid.New().ID()%100000)
	uuId := fmt.Sprintf("USU-INIS-%s", randomDigits)
	pegawai, err := service.pegawaiRepository.FindByNip(ctx, tx, request.PegawaiId)
	if err != nil {
		if err == sql.ErrNoRows {
			return usulan.UsulanInisiatifResponse{}, fmt.Errorf("pegawai dengan NIP %s tidak ditemukan dalam database", request.PegawaiId)
		}
		return usulan.UsulanInisiatifResponse{}, fmt.Errorf("terjadi kesalahan saat mencari data pegawai: %v", err)
	}

	// Cek OPD
	opd, err := service.opdRepository.FindByKodeOpd(ctx, tx, request.KodeOpd)
	if err != nil {
		if err == sql.ErrNoRows {
			return usulan.UsulanInisiatifResponse{}, fmt.Errorf("OPD dengan kode %s tidak ditemukan dalam database", request.KodeOpd)
		}
		return usulan.UsulanInisiatifResponse{}, fmt.Errorf("terjadi kesalahan saat mencari data OPD: %v", err)
	}

	domainUsulanInisiatif := domain.UsulanInisiatif{
		Id:          uuId,
		Usulan:      request.Usulan,
		Manfaat:     request.Manfaat,
		Uraian:      request.Uraian,
		Tahun:       request.Tahun,
		RekinId:     request.RekinId,
		PegawaiId:   pegawai.Nip,
		NamaPegawai: pegawai.NamaPegawai,
		KodeOpd:     opd.KodeOpd,
		NamaOpd:     opd.NamaOpd,
		Status:      request.Status,
	}

	usulanInisiatif, err := service.UsulanInisiatifRepository.Create(ctx, tx, domainUsulanInisiatif)
	if err != nil {
		return usulan.UsulanInisiatifResponse{}, err
	}

	response := helper.ToUsulanInisiatifResponse(usulanInisiatif)
	return response, nil
}

func (service *UsulanInisiatifServiceImpl) Update(ctx context.Context, request usulan.UsulanInisiatifUpdateRequest) (usulan.UsulanInisiatifResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return usulan.UsulanInisiatifResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	// Cari usulan inisiatif berdasarkan ID
	existingUsulan, err := service.UsulanInisiatifRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return usulan.UsulanInisiatifResponse{}, fmt.Errorf("usulan inisiatif dengan ID %s tidak ditemukan", request.Id)
	}

	existingUsulan.Usulan = request.Usulan
	existingUsulan.Manfaat = request.Manfaat
	existingUsulan.Uraian = request.Uraian
	existingUsulan.Tahun = request.Tahun
	existingUsulan.PegawaiId = request.PegawaiId
	existingUsulan.KodeOpd = request.KodeOpd
	existingUsulan.Status = request.Status
	usulanInisiatif, err := service.UsulanInisiatifRepository.Update(ctx, tx, existingUsulan)
	if err != nil {
		return usulan.UsulanInisiatifResponse{}, err
	}

	response := helper.ToUsulanInisiatifResponse(usulanInisiatif)
	return response, nil
}

func (service *UsulanInisiatifServiceImpl) FindById(ctx context.Context, idUsulan string) (usulan.UsulanInisiatifResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return usulan.UsulanInisiatifResponse{}, fmt.Errorf("gagal memulai transaksi: %v", err)
	}
	defer helper.CommitOrRollback(tx)

	usulanInisiatif, err := service.UsulanInisiatifRepository.FindById(ctx, tx, idUsulan)
	if err != nil {
		return usulan.UsulanInisiatifResponse{}, fmt.Errorf("usulan inisiatif dengan ID %s tidak ditemukan", idUsulan)
	}

	response := helper.ToUsulanInisiatifResponse(usulanInisiatif)
	return response, nil
}

func (service *UsulanInisiatifServiceImpl) FindAll(ctx context.Context, pegawaiId *string, isActive *bool, rekinId *string) ([]usulan.UsulanInisiatifResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []usulan.UsulanInisiatifResponse{}, fmt.Errorf("gagal memulai transaksi: %v", err)
	}
	defer helper.CommitOrRollback(tx)

	usulanInisiatif, err := service.UsulanInisiatifRepository.FindAll(ctx, tx, pegawaiId, isActive, rekinId)
	if err != nil {
		return []usulan.UsulanInisiatifResponse{}, err
	}

	response := helper.ToUsulanInisiatifResponses(usulanInisiatif)
	return response, nil
}

func (service *UsulanInisiatifServiceImpl) Delete(ctx context.Context, idUsulan string) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return fmt.Errorf("gagal memulai transaksi: %v", err)
	}
	defer helper.CommitOrRollback(tx)

	// Cari usulan inisiatif berdasarkan ID
	_, err = service.UsulanInisiatifRepository.FindById(ctx, tx, idUsulan)
	if err != nil {
		return fmt.Errorf("usulan inisiatif dengan ID %s tidak ditemukan", idUsulan)
	}

	// Jika usulan inisiatif ditemukan, lanjutkan dengan penghapusan
	err = service.UsulanInisiatifRepository.Delete(ctx, tx, idUsulan)
	if err != nil {
		return fmt.Errorf("gagal menghapus usulan inisiatif: %v", err)
	}

	return nil
}
