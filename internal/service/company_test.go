package service_test

import (
	"context"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/internal/service"
	"course/internal/service/utils"
)

type CompanySuite struct {
	suite.Suite

	companyService service.CompanyService
	companyID      int64
}

func (s *CompanySuite) Test_Company_GetCompany1(t provider.T) {
	t.Title("[GetCompany] Incorrect company ID")
	t.Tags("company")
	t.Parallel()
	t.WithNewStep("Incorrect company ID", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.CompanyBuilder{}.
			WithCompanyID(-1).
			ToGetDTO()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		company, err := s.companyService.GetCompany(ctx, request)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(company)
	})
}

func (s *CompanySuite) Test_Company_GetCompany2(t provider.T) {
	t.Title("[GetCompany] Success")
	t.Tags("company")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.CompanyBuilder{}.
			WithCompanyID(s.companyID).
			ToGetDTO()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		company, err := s.companyService.GetCompany(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(company)
		sCtx.Assert().NotNil(company.ID)
		sCtx.Assert().Equal("Test", company.Name)
		sCtx.Assert().Equal("Test", company.City)
	})
}
