package service

import (
	"context"
	"fmt"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/logger"
)

type CompanyService interface {
	GetCompany(ctx context.Context, request *dto.GetCompanyRequest) (*model.Company, error)
}

type companyServiceImpl struct {
	logger         logger.Interface
	companyStorage storage.CompanyStorage
}

func NewCompanyService(logger logger.Interface, companyStorage storage.CompanyStorage) CompanyService {
	return &companyServiceImpl{
		logger:         logger,
		companyStorage: companyStorage,
	}
}

func (c *companyServiceImpl) GetCompany(ctx context.Context, request *dto.GetCompanyRequest) (*model.Company, error) {
	c.logger.Infof("get company by ID %d", request.CompanyID)

	company, err := c.companyStorage.GetByID(ctx, request)
	if err != nil {
		c.logger.Errorf("get company: %s", err.Error())
		return nil, fmt.Errorf("get company: %w", err)
	}

	return company, nil
}
