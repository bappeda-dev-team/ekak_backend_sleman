package service

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/domain"
	"ekak_kab_sleman/model/web/periodetahun"
	"ekak_kab_sleman/repository"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type PeriodeServiceImpl struct {
	PeriodeRepository repository.PeriodeRepository
	DB                *sql.DB
}

func NewPeriodeServiceImpl(periodeRepository repository.PeriodeRepository, DB *sql.DB) *PeriodeServiceImpl {
	return &PeriodeServiceImpl{
		PeriodeRepository: periodeRepository,
		DB:                DB,
	}
}

func (service *PeriodeServiceImpl) generateRandomId(ctx context.Context, tx *sql.Tx) int {
	rand.Seed(time.Now().UnixNano())
	for {
		// Generate random number between 10000-99999
		id := rand.Intn(90000) + 10000
		if !service.PeriodeRepository.IsIdExists(ctx, tx, id) {
			return id
		}
	}
}

func (service *PeriodeServiceImpl) validatePeriodeOverlap(ctx context.Context, tx *sql.Tx, tahunAwal, tahunAkhir, jenisPeriode string) error {
	existingPeriodes, err := service.PeriodeRepository.FindOverlappingPeriodes(ctx, tx, tahunAwal, tahunAkhir, jenisPeriode)
	if err != nil {
		return err
	}

	if len(existingPeriodes) > 0 {
		return fmt.Errorf("periode tahun %s-%s dengan jenis periode %s overlap dengan periode yang sudah ada: %s-%s (jenis periode: %s)",
			tahunAwal, tahunAkhir, jenisPeriode, existingPeriodes[0].TahunAwal, existingPeriodes[0].TahunAkhir, existingPeriodes[0].JenisPeriode)
	}
	return nil
}

func (service *PeriodeServiceImpl) validatePeriodeOverlapExcludeCurrent(ctx context.Context, tx *sql.Tx, currentId int, tahunAwal, tahunAkhir, jenisPeriode string) error {
	existingPeriodes, err := service.PeriodeRepository.FindOverlappingPeriodesExcludeCurrent(ctx, tx, currentId, tahunAwal, tahunAkhir, jenisPeriode)
	if err != nil {
		return err
	}

	if len(existingPeriodes) > 0 {
		return fmt.Errorf("periode tahun %s-%s dengan jenis periode %s overlap dengan periode yang sudah ada: %s-%s (jenis periode: %s)",
			tahunAwal, tahunAkhir, jenisPeriode, existingPeriodes[0].TahunAwal, existingPeriodes[0].TahunAkhir, existingPeriodes[0].JenisPeriode)
	}
	return nil
}

func (service *PeriodeServiceImpl) Create(ctx context.Context, request periodetahun.PeriodeCreateRequest) (periodetahun.PeriodeResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return periodetahun.PeriodeResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	// Validasi overlap periode dengan mempertimbangkan jenis periode
	if err := service.validatePeriodeOverlap(ctx, tx, request.TahunAwal, request.TahunAkhir, request.JenisPeriode); err != nil {
		return periodetahun.PeriodeResponse{}, err
	}

	randomId := service.generateRandomId(ctx, tx)

	periode := domain.Periode{
		Id:           randomId,
		TahunAwal:    request.TahunAwal,
		TahunAkhir:   request.TahunAkhir,
		JenisPeriode: request.JenisPeriode,
	}

	// Simpan periode
	newPeriode, err := service.PeriodeRepository.Save(ctx, tx, periode)
	if err != nil {
		return periodetahun.PeriodeResponse{}, err
	}

	// Generate dan simpan tahun-tahun periode
	tahunAwal, err := strconv.Atoi(request.TahunAwal)
	if err != nil {
		return periodetahun.PeriodeResponse{}, fmt.Errorf("invalid tahun awal format: %v", err)
	}

	tahunAkhir, err := strconv.Atoi(request.TahunAkhir)
	if err != nil {
		return periodetahun.PeriodeResponse{}, fmt.Errorf("invalid tahun akhir format: %v", err)
	}

	if tahunAkhir < tahunAwal {
		return periodetahun.PeriodeResponse{}, fmt.Errorf("tahun akhir (%d) harus lebih besar dari tahun awal (%d)", tahunAkhir, tahunAwal)
	}

	var tahunList []string
	for tahun := tahunAwal; tahun <= tahunAkhir; tahun++ {
		tahunStr := strconv.Itoa(tahun)
		tahunPeriode := domain.TahunPeriode{
			IdPeriode: newPeriode.Id,
			Tahun:     tahunStr,
		}

		if err := service.PeriodeRepository.SaveTahunPeriode(ctx, tx, tahunPeriode); err != nil {
			return periodetahun.PeriodeResponse{}, err
		}
		tahunList = append(tahunList, tahunStr)
	}

	response := periodetahun.PeriodeResponse{
		Id:           newPeriode.Id,
		TahunAwal:    newPeriode.TahunAwal,
		TahunAkhir:   newPeriode.TahunAkhir,
		JenisPeriode: newPeriode.JenisPeriode,
		TahunList:    tahunList,
	}

	return response, nil
}

func (service *PeriodeServiceImpl) Update(ctx context.Context, request periodetahun.PeriodeUpdateRequest) (periodetahun.PeriodeResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return periodetahun.PeriodeResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	// Validasi periode exists
	_, err = service.PeriodeRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return periodetahun.PeriodeResponse{}, err
	}

	// Validasi overlap periode dengan mempertimbangkan jenis periode
	if err := service.validatePeriodeOverlapExcludeCurrent(ctx, tx, request.Id, request.TahunAwal, request.TahunAkhir, request.JenisPeriode); err != nil {
		return periodetahun.PeriodeResponse{}, err
	}

	// Update periode dengan menyertakan jenis_periode
	periode := domain.Periode{
		Id:           request.Id,
		TahunAwal:    request.TahunAwal,
		TahunAkhir:   request.TahunAkhir,
		JenisPeriode: request.JenisPeriode,
	}

	updatedPeriode, err := service.PeriodeRepository.Update(ctx, tx, periode)
	if err != nil {
		return periodetahun.PeriodeResponse{}, err
	}

	// Hapus tahun periode lama
	if err := service.PeriodeRepository.DeleteTahunPeriode(ctx, tx, periode.Id); err != nil {
		return periodetahun.PeriodeResponse{}, err
	}

	// Generate dan simpan tahun-tahun periode baru
	tahunAwal, err := strconv.Atoi(request.TahunAwal)
	if err != nil {
		return periodetahun.PeriodeResponse{}, fmt.Errorf("invalid tahun awal format: %v", err)
	}

	tahunAkhir, err := strconv.Atoi(request.TahunAkhir)
	if err != nil {
		return periodetahun.PeriodeResponse{}, fmt.Errorf("invalid tahun akhir format: %v", err)
	}

	if tahunAkhir < tahunAwal {
		return periodetahun.PeriodeResponse{}, fmt.Errorf("tahun akhir (%d) harus lebih besar dari tahun awal (%d)", tahunAkhir, tahunAwal)
	}

	var tahunList []string
	for tahun := tahunAwal; tahun <= tahunAkhir; tahun++ {
		tahunStr := strconv.Itoa(tahun)
		tahunPeriode := domain.TahunPeriode{
			IdPeriode: updatedPeriode.Id,
			Tahun:     tahunStr,
		}

		if err := service.PeriodeRepository.SaveTahunPeriode(ctx, tx, tahunPeriode); err != nil {
			return periodetahun.PeriodeResponse{}, err
		}
		tahunList = append(tahunList, tahunStr)
	}

	return periodetahun.PeriodeResponse{
		Id:           updatedPeriode.Id,
		TahunAwal:    updatedPeriode.TahunAwal,
		TahunAkhir:   updatedPeriode.TahunAkhir,
		JenisPeriode: updatedPeriode.JenisPeriode,
		TahunList:    tahunList,
	}, nil
}

func (service *PeriodeServiceImpl) FindByTahun(ctx context.Context, tahun string) (periodetahun.PeriodeResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return periodetahun.PeriodeResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	periode, err := service.PeriodeRepository.FindByTahun(ctx, tx, tahun)
	if err != nil {
		return periodetahun.PeriodeResponse{}, err
	}

	tahunAwal, err := strconv.Atoi(periode.TahunAwal)
	if err != nil {
		return periodetahun.PeriodeResponse{}, fmt.Errorf("invalid tahun awal format: %v", err)
	}

	tahunAkhir, err := strconv.Atoi(periode.TahunAkhir)
	if err != nil {
		return periodetahun.PeriodeResponse{}, fmt.Errorf("invalid tahun akhir format: %v", err)
	}

	var tahunList []string
	for tahun := tahunAwal; tahun <= tahunAkhir; tahun++ {
		tahunList = append(tahunList, strconv.Itoa(tahun))
	}

	return periodetahun.PeriodeResponse{
		Id:           periode.Id,
		TahunAwal:    periode.TahunAwal,
		TahunAkhir:   periode.TahunAkhir,
		JenisPeriode: periode.JenisPeriode,
		TahunList:    tahunList,
	}, nil
}

func (service *PeriodeServiceImpl) FindAll(ctx context.Context, jenis_periode string) ([]periodetahun.PeriodeResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	periodes, err := service.PeriodeRepository.FindAll(ctx, tx, jenis_periode)
	if err != nil {
		return nil, err
	}

	var periodesResponse []periodetahun.PeriodeResponse
	for _, periode := range periodes {
		// Generate tahunList seperti di FindByTahun
		tahunAwal, err := strconv.Atoi(periode.TahunAwal)
		if err != nil {
			return nil, fmt.Errorf("invalid tahun awal format: %v", err)
		}

		tahunAkhir, err := strconv.Atoi(periode.TahunAkhir)
		if err != nil {
			return nil, fmt.Errorf("invalid tahun akhir format: %v", err)
		}

		var tahunList []string
		for tahun := tahunAwal; tahun <= tahunAkhir; tahun++ {
			tahunList = append(tahunList, strconv.Itoa(tahun))
		}

		periodesResponse = append(periodesResponse, periodetahun.PeriodeResponse{
			Id:           periode.Id,
			TahunAwal:    periode.TahunAwal,
			TahunAkhir:   periode.TahunAkhir,
			JenisPeriode: periode.JenisPeriode,
			TahunList:    tahunList,
		})
	}

	return periodesResponse, nil
}

func (service *PeriodeServiceImpl) FindById(ctx context.Context, id int) (periodetahun.PeriodeResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return periodetahun.PeriodeResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	periode, err := service.PeriodeRepository.FindById(ctx, tx, id)
	if err != nil {
		return periodetahun.PeriodeResponse{}, err
	}

	return periodetahun.PeriodeResponse{
		Id:           periode.Id,
		TahunAwal:    periode.TahunAwal,
		TahunAkhir:   periode.TahunAkhir,
		JenisPeriode: periode.JenisPeriode,
	}, nil
}

func (service *PeriodeServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.PeriodeRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}
	return nil
}
