package postgres_test

import (
	"context"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/internal/model"
	"course/internal/storage/mocks"
	"course/internal/storage/utils"
)

type EmployeeStorageSuite struct {
	suite.Suite

	employeeMockStorage mocks.EmployeeStorage
}

func (c *EmployeeStorageSuite) BeforeAll(t provider.T) {
	t.Title("Init employee mock storage")
	c.employeeMockStorage = *mocks.NewEmployeeStorage(t)
	t.Tags("fixture", "employee")
}

func (c *EmployeeStorageSuite) Test_EmployeeStorage_Register(t provider.T) {
	t.Title("[Register] Register employee test")
	t.Tags("storage", "postgres", "employee")
	t.Parallel()
	t.WithNewStep("Register employee test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.AuthObjectMother{
			CompanyID: 1,
		}.DefaultRegisterEmployeeRequest1()
		expEmployee := &model.Employee{
			ID:             model.ToEmployeeID(1),
			FullName:       "123",
			PhoneNumber:    "123",
			CompanyID:      model.ToCompanyID(1),
			Post:           model.ToPostTypeFromInt(1),
			Password:       "123",
			RefreshToken:   "123",
			TokenExpiredAt: nil,
			DateOfBirth:    nil,
		}

		c.employeeMockStorage.
			On("Register", ctx, request).
			Return(expEmployee, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		employee, err := c.employeeMockStorage.Register(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(employee)
		sCtx.Assert().Equal(expEmployee, employee)
	})
}

func (c *EmployeeStorageSuite) Test_EmployeeStorage_UpdateRefreshToken(t provider.T) {
	t.Title("[UpdateRefreshToken] Update refresh token test")
	t.Tags("storage", "postgres", "employee")
	t.Parallel()
	t.WithNewStep("Update refresh token test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.AuthObjectMother{
			CompanyID: 1,
		}.DefaultUpdateRefreshTokenRequest()

		c.employeeMockStorage.
			On("UpdateRefreshToken", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := c.employeeMockStorage.UpdateRefreshToken(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (c *EmployeeStorageSuite) Test_EmployeeStorage_GetByPhone(t provider.T) {
	t.Title("[GetByPhone] Get employee by phone test")
	t.Tags("storage", "postgres", "employee")
	t.Parallel()
	t.WithNewStep("Get employee by phone test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.EmployeeObjectMother{
			EmployeeID: 1,
			InfoCardID: 1,
		}.DefaultGetEmployeeRequest2()
		expEmployee := &model.Employee{
			ID:             model.ToEmployeeID(1),
			FullName:       "123",
			PhoneNumber:    "123",
			CompanyID:      model.ToCompanyID(1),
			Post:           model.ToPostTypeFromInt(1),
			Password:       "123",
			RefreshToken:   "123",
			TokenExpiredAt: nil,
			DateOfBirth:    nil,
		}

		c.employeeMockStorage.
			On("GetByPhone", ctx, request).
			Return(expEmployee, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		employee, err := c.employeeMockStorage.GetByPhone(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(employee)
		sCtx.Assert().Equal(expEmployee, employee)
	})
}

func (c *EmployeeStorageSuite) Test_EmployeeStorage_GetByInfoCardID(t provider.T) {
	t.Title("[GetByInfoCardID] Get employee by info card ID test")
	t.Tags("storage", "postgres", "employee")
	t.Parallel()
	t.WithNewStep("Get employee by info card ID test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.EmployeeObjectMother{
			EmployeeID: 1,
			InfoCardID: 1,
		}.DefaultGetEmployeeByInfoCardIDRequest()
		expEmployee := &model.Employee{
			ID:             model.ToEmployeeID(1),
			FullName:       "123",
			PhoneNumber:    "123",
			CompanyID:      model.ToCompanyID(1),
			Post:           model.ToPostTypeFromInt(1),
			Password:       "123",
			RefreshToken:   "123",
			TokenExpiredAt: nil,
			DateOfBirth:    nil,
		}

		c.employeeMockStorage.
			On("GetByInfoCardID", ctx, request).
			Return(expEmployee, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		employee, err := c.employeeMockStorage.GetByInfoCardID(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(employee)
		sCtx.Assert().Equal(expEmployee, employee)
	})
}

func (c *EmployeeStorageSuite) Test_EmployeeStorage_Delete(t provider.T) {
	t.Title("[Delete] Delete employee test")
	t.Tags("storage", "postgres", "employee")
	t.Parallel()
	t.WithNewStep("Delete employee test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.EmployeeObjectMother{
			EmployeeID: 1,
			InfoCardID: 1,
		}.DefaultDeleteEmployeeRequest()

		c.employeeMockStorage.
			On("Delete", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := c.employeeMockStorage.Delete(ctx, request)

		sCtx.Assert().NoError(err)
	})
}
