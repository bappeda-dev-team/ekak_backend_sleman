package service

import (
	"context"
	"database/sql"
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/domain"
	"ekak_kab_sleman/model/web/user"
	"ekak_kab_sleman/repository"
)

type RoleServiceImpl struct {
	RoleRepository repository.RoleRepository
	DB             *sql.DB
}

func NewRoleServiceImpl(roleRepository repository.RoleRepository, DB *sql.DB) *RoleServiceImpl {
	return &RoleServiceImpl{
		RoleRepository: roleRepository,
		DB:             DB,
	}
}

func (service *RoleServiceImpl) Create(ctx context.Context, request user.RoleCreateRequest) (user.RoleResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return user.RoleResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	role := domain.Roles{
		Role: request.Role,
	}

	roleResponse, err := service.RoleRepository.Create(ctx, tx, role)
	if err != nil {
		return user.RoleResponse{}, err
	}
	return user.RoleResponse{
		Id:   roleResponse.Id,
		Role: roleResponse.Role,
	}, nil
}

func (service *RoleServiceImpl) Update(ctx context.Context, request user.RoleUpdateRequest) (user.RoleResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return user.RoleResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	role := domain.Roles{
		Id:   request.Id,
		Role: request.Role,
	}

	roleResponse, err := service.RoleRepository.Update(ctx, tx, role)
	if err != nil {
		return user.RoleResponse{}, err
	}
	return user.RoleResponse{
		Id:   roleResponse.Id,
		Role: roleResponse.Role,
	}, nil
}

func (service *RoleServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.RoleRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}
	return nil
}

func (service *RoleServiceImpl) FindById(ctx context.Context, id int) (user.RoleResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return user.RoleResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	role, err := service.RoleRepository.FindById(ctx, tx, id)
	if err != nil {
		return user.RoleResponse{}, err
	}
	return user.RoleResponse{
		Id:   role.Id,
		Role: role.Role,
	}, nil
}

func (service *RoleServiceImpl) FindAll(ctx context.Context) ([]user.RoleResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []user.RoleResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	roles, err := service.RoleRepository.FindAll(ctx, tx)
	if err != nil {
		return []user.RoleResponse{}, err
	}

	var roleResponses []user.RoleResponse
	for _, role := range roles {
		roleResponses = append(roleResponses, user.RoleResponse{
			Id:   role.Id,
			Role: role.Role,
		})
	}
	return roleResponses, nil
}
