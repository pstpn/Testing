package service_test

import (
	"context"
	"fmt"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/internal/model"
	"course/internal/service"
	"course/internal/service/utils"
	"course/internal/storage/mocks"
)

type CompanySuite struct {
	suite.Suite
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

		companyMockStorage := mocks.NewCompanyStorage(t)
		companyMockStorage.
			On("GetByID", ctx, request).
			Return(nil, fmt.Errorf("incorrect company ID"))

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		company, err := service.NewCompanyService(utils.NewMockLogger(), companyMockStorage).GetCompany(ctx, request)

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
			WithCompanyID(1).
			ToGetDTO()

		companyMockStorage := mocks.NewCompanyStorage(t)
		companyMockStorage.
			On("GetByID", ctx, request).
			Return(&model.Company{
				ID:   model.ToCompanyID(1),
				Name: "test",
				City: "test",
			}, nil)

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		company, err := service.NewCompanyService(utils.NewMockLogger(), companyMockStorage).GetCompany(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(company)
		sCtx.Assert().NotNil(company.ID)
		sCtx.Assert().Equal(int64(1), company.ID.Int())
		sCtx.Assert().Equal("test", company.Name)
		sCtx.Assert().Equal("test", company.City)
	})
}
