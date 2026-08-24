package service

import (
	"context"
	"ekak_kab_sleman/model/web/isustrategis"
)

type CSFService interface {
	FindByTahun(ctx context.Context, tahun string) ([]isustrategis.CSFResponse, error)
	FindById(ctx context.Context, csfID int) (isustrategis.CSFResponse, error)
}
