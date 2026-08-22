package service

import (
	"context"
	"ekak_kab_sleman/model/web/gambaranumum"
)

type GambaranUmumService interface {
	Create(ctx context.Context, request gambaranumum.GambaranUmumCreateRequest) (gambaranumum.GambaranUmumResponse, error)
	Update(ctx context.Context, request gambaranumum.GambaranUmumUpdateRequest) (gambaranumum.GambaranUmumResponse, error)
	FindById(ctx context.Context, id string) (gambaranumum.GambaranUmumResponse, error)
	FindAll(ctx context.Context, rekinId string) ([]gambaranumum.GambaranUmumResponse, error)
	Delete(ctx context.Context, id string) error
}
