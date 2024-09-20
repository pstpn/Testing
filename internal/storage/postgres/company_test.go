package postgres_test

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/internal/service/dto"
	"course/internal/storage"
	"course/internal/storage/utils"
)

type CompanyStorageSuite struct {
	suite.Suite

	companyStorage storage.CompanyStorage
}

func (c *CompanyStorageSuite) Test_CompanyStorage_Create(t provider.T) {
	t.Title("[Create] Create company test")
	t.Tags("storage", "postgres", "company")
	t.Parallel()
	t.WithNewStep("Create company test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.CompanyBuilder{}.
			WithName("teStik").
			WithCity("Testovik").
			ToCreateDTO()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		company, err := c.companyStorage.Create(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(company)
		sCtx.Assert().NotNil(company.ID)
		sCtx.Assert().Equal(request.Name, company.Name)
		sCtx.Assert().Equal(request.City, company.City)

		err = c.companyStorage.Delete(context.TODO(), &dto.DeleteCompanyRequest{CompanyID: company.ID.Int()})
		sCtx.Assert().NoError(err)
	})
}

func (c *CompanyStorageSuite) Test_CompanyStorage_GetByID(t provider.T) {
	t.Title("[GetByID] Get company by ID test")
	t.Tags("storage", "postgres", "company")
	t.Parallel()
	t.WithNewStep("Get company by ID test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expCompany, err := c.companyStorage.Create(
			ctx,
			utils.CompanyBuilder{}.
				WithName("teStik").
				WithCity("Testovik").
				ToCreateDTO(),
		)
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(expCompany)
		request := utils.CompanyBuilder{}.
			WithCompanyID(expCompany.ID.Int()).
			ToGetDTO()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		company, err := c.companyStorage.GetByID(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(company)
		sCtx.Assert().Equal(expCompany.ID, company.ID)
		sCtx.Assert().Equal(expCompany.Name, company.Name)
		sCtx.Assert().Equal(expCompany.City, company.City)

		err = c.companyStorage.Delete(context.TODO(), &dto.DeleteCompanyRequest{CompanyID: company.ID.Int()})
		sCtx.Assert().NoError(err)
	})
}

func (c *CompanyStorageSuite) Test_CompanyStorage_Delete(t provider.T) {
	t.Title("[Delete] Delete company test")
	t.Tags("storage", "postgres", "company")
	t.Parallel()
	t.WithNewStep("Delete company test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expCompany, err := c.companyStorage.Create(
			ctx,
			utils.CompanyBuilder{}.
				WithName("teStik").
				WithCity("Testovik").
				ToCreateDTO(),
		)
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(expCompany)
		request := utils.CompanyBuilder{}.
			WithCompanyID(expCompany.ID.Int()).
			ToDeleteDTO()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err = c.companyStorage.Delete(ctx, request)

		sCtx.Assert().NoError(err)

		company, err := c.companyStorage.GetByID(
			ctx,
			utils.CompanyBuilder{}.
				WithCompanyID(expCompany.ID.Int()).
				ToGetDTO(),
		)

		sCtx.Assert().Error(err)
		sCtx.Assert().EqualError(err, pgx.ErrNoRows.Error())
		sCtx.Assert().Nil(company)

		err = c.companyStorage.Delete(ctx, request)

		sCtx.Assert().NoError(err)
	})
}
