//go:build unit

package postgres_test

import (
	"context"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/internal/model"
	"course/internal/storage/mocks"
	"course/internal/storage/utils"
)

type CompanyStorageSuite struct {
	suite.Suite

	companyMockStorage mocks.CompanyStorage
}

func (c *CompanyStorageSuite) BeforeAll(t provider.T) {
	t.Title("Init company mock storage")
	c.companyMockStorage = *mocks.NewCompanyStorage(t)
	t.Tags("fixture", "company")
}

func (c *CompanyStorageSuite) Test_CheckpointStorage_CreateCompany(t provider.T) {
	t.Title("[Create] Create company test")
	t.Tags("storage", "postgres", "company")
	t.Parallel()
	t.WithNewStep("Create company test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.CompanyBuilder{}.
			WithName("teStik").
			WithCity("Testovik").
			ToCreateDTO()
		expCompany := &model.Company{
			ID:   model.ToCompanyID(1),
			Name: "13",
			City: "13",
		}

		c.companyMockStorage.
			On("Create", ctx, request).
			Return(expCompany, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		company, err := c.companyMockStorage.Create(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(company)
		sCtx.Assert().Equal(expCompany, company)
	})
}

func (c *CompanyStorageSuite) Test_CompanyStorage_GetByID(t provider.T) {
	t.Title("[GetByID] Get company by ID test")
	t.Tags("storage", "postgres", "company")
	t.Parallel()
	t.WithNewStep("Get company by ID test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.CompanyBuilder{}.
			WithCompanyID(1).
			ToGetDTO()
		expCompany := &model.Company{
			ID:   model.ToCompanyID(1),
			Name: "13",
			City: "13",
		}

		c.companyMockStorage.
			On("GetByID", ctx, request).
			Return(expCompany, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		company, err := c.companyMockStorage.GetByID(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(company)
		sCtx.Assert().Equal(expCompany, company)
	})
}

func (c *CompanyStorageSuite) Test_CompanyStorage_Delete(t provider.T) {
	t.Title("[Delete] Delete company test")
	t.Tags("storage", "postgres", "company")
	t.Parallel()
	t.WithNewStep("Delete company test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.CompanyBuilder{}.
			WithCompanyID(1).
			ToDeleteDTO()

		c.companyMockStorage.
			On("Delete", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := c.companyMockStorage.Delete(ctx, request)

		sCtx.Assert().NoError(err)
	})
}
