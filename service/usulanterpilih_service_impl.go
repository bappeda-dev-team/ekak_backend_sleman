package service

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/domain"
	"ekak_kab_sleman/model/web/usulan"
	"ekak_kab_sleman/repository"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UsulanTerpilihServiceImpl struct {
	UsulanTerpilihRepository repository.UsulanTerpilihRepository
	DB                       *sql.DB
	Validate                 *validator.Validate
}

func NewUsulanTerpilihServiceImpl(usulanTerpilihRepository repository.UsulanTerpilihRepository, DB *sql.DB, validate *validator.Validate) *UsulanTerpilihServiceImpl {
	return &UsulanTerpilihServiceImpl{
		UsulanTerpilihRepository: usulanTerpilihRepository,
		DB:                       DB,
		Validate:                 validate,
	}
}

func (service *UsulanTerpilihServiceImpl) Create(ctx context.Context, request usulan.UsulanTerpilihCreateRequest) (usulan.UsulanTerpilihResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Memeriksa apakah jenis_usulan dan usulan_id sesuai dengan case
	isValid, err := service.UsulanTerpilihRepository.ValidateJenisAndUsulanId(ctx, tx, request.JenisUsulan, request.UsulanId)
	if err != nil {
		return usulan.UsulanTerpilihResponse{}, fmt.Errorf("gagal memvalidasi jenis usulan dan usulan ID: %v", err)
	}
	if !isValid {
		return usulan.UsulanTerpilihResponse{}, fmt.Errorf("jenis usulan atau usulan ID tidak sesuai dengan yang mandatori")
	}

	// Memeriksa apakah usulan dengan jenis dan id yang sama sudah ada
	exists, err := service.UsulanTerpilihRepository.ExistsByJenisAndUsulanId(ctx, tx, request.JenisUsulan, request.UsulanId)
	if err != nil {
		return usulan.UsulanTerpilihResponse{}, fmt.Errorf("gagal memeriksa keberadaan usulan: %v", err)
	}
	if exists {
		return usulan.UsulanTerpilihResponse{}, fmt.Errorf("usulan dengan jenis dan id yang sama sudah ada")
	}

	// Membuat UUID baru
	randomDigits := fmt.Sprintf("%05d", uuid.New().ID()%100000)
	uuId := fmt.Sprintf("USU-TERP-%s", randomDigits)

	usulanTerpilih := domain.UsulanTerpilih{
		Id:          uuId,
		Keterangan:  request.Keterangan,
		JenisUsulan: request.JenisUsulan,
		UsulanId:    request.UsulanId,
		RekinId:     request.RekinId,
		Tahun:       request.Tahun,
		KodeOpd:     request.KodeOpd,
	}

	usulanTerpilih, err = service.UsulanTerpilihRepository.Create(ctx, tx, usulanTerpilih)
	if err != nil {
		return usulan.UsulanTerpilihResponse{}, fmt.Errorf("gagal membuat usulan terpilih: %v", err)
	}

	return helper.ToUsulanTerpilihResponse(usulanTerpilih), nil
}

func (service *UsulanTerpilihServiceImpl) Delete(ctx context.Context, idUsulan string) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.UsulanTerpilihRepository.Delete(ctx, tx, idUsulan)
	helper.PanicIfError(err)

	return nil
}
