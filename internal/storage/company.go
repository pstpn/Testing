package storage

import (
	"context"

	"course/internal/model"
	"course/internal/service/dto"
)

//go:generate mockery --name=CompanyStorage
type CompanyStorage interface {
	Create(ctx context.Context, request *dto.CreateCompanyRequest) (*model.Company, error)
	GetByID(ctx context.Context, request *dto.GetCompanyRequest) (*model.Company, error)
	Delete(ctx context.Context, request *dto.DeleteCompanyRequest) error
}
