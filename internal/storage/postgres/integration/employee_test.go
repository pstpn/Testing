//go:build integration

package postgres_test

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/internal/storage/utils"
)

type EmployeeStorageSuite struct {
	suite.Suite

	employeeStorage storage.EmployeeStorage
	companyID       int64
	infoCardID      int64
	employeeID      int64
}

func (c *EmployeeStorageSuite) Test_IntegrationEmployeeStorage_Register(t provider.T) {
	t.Title("[Register] Register employee test")
	t.Tags("storage", "postgres", "employee")
	t.Parallel()
	t.WithNewStep("Register employee test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.AuthObjectMother{
			CompanyID: c.companyID,
		}.DefaultRegisterEmployeeRequest1()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		employee, err := c.employeeStorage.Register(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(employee)
		sCtx.Assert().NotNil(employee.ID)
		sCtx.Assert().Equal(request.PhoneNumber, employee.PhoneNumber)
		sCtx.Assert().Equal(request.FullName, employee.FullName)
		sCtx.Assert().Equal(request.Post, employee.Post.Int())
		sCtx.Assert().Equal(request.Password, employee.Password)
		sCtx.Assert().Equal(request.CompanyID, employee.CompanyID.Int())
		sCtx.Assert().Equal(request.TokenExpiredAt, employee.TokenExpiredAt)
		sCtx.Assert().Equal(request.DateOfBirth, employee.DateOfBirth)
		sCtx.Assert().Equal(request.RefreshToken, employee.RefreshToken)

		err = c.employeeStorage.Delete(context.TODO(), &dto.DeleteEmployeeRequest{EmployeeID: employee.ID.Int()})
		sCtx.Assert().NoError(err)
	})
}

func (c *EmployeeStorageSuite) Test_IntegrationEmployeeStorage_UpdateRefreshToken(t provider.T) {
	t.Title("[UpdateRefreshToken] Update refresh token test")
	t.Tags("storage", "postgres", "employee")
	t.Parallel()
	t.WithNewStep("Update refresh token test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expEmployee, err := c.employeeStorage.Register(
			ctx,
			utils.AuthObjectMother{
				CompanyID: c.companyID,
			}.DefaultRegisterEmployeeRequest2(),
		)
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(expEmployee)
		request := utils.AuthObjectMother{
			CompanyID:  c.companyID,
			EmployeeID: expEmployee.ID.Int(),
		}.DefaultUpdateRefreshTokenRequest()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err = c.employeeStorage.UpdateRefreshToken(ctx, request)

		sCtx.Assert().NoError(err)

		employee, err := c.employeeStorage.GetByPhone(
			ctx,
			utils.EmployeeObjectMother{
				EmployeeID: -1,
			}.DefaultGetEmployeeRequest2(),
		)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(employee)
		sCtx.Assert().Equal(request.EmployeeID, employee.ID.Int())
		sCtx.Assert().Equal(request.RefreshToken, employee.RefreshToken)
		sCtx.Assert().Equal(request.TokenExpiredAt, employee.TokenExpiredAt)

		err = c.employeeStorage.Delete(context.TODO(), &dto.DeleteEmployeeRequest{EmployeeID: employee.ID.Int()})
		sCtx.Assert().NoError(err)
	})
}

func (c *EmployeeStorageSuite) Test_IntegrationEmployeeStorage_GetByPhone(t provider.T) {
	t.Title("[GetByPhone] Get employee by phone test")
	t.Tags("storage", "postgres", "employee")
	t.Parallel()
	t.WithNewStep("Get employee by phone test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expEmployee, err := c.employeeStorage.Register(
			ctx,
			utils.AuthObjectMother{
				CompanyID: c.companyID,
			}.DefaultRegisterEmployeeRequest3(),
		)
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(expEmployee)
		request := utils.EmployeeObjectMother{
			EmployeeID: expEmployee.ID.Int(),
		}.DefaultGetEmployeeRequest3()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		employee, err := c.employeeStorage.GetByPhone(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(employee)
		sCtx.Assert().Equal(expEmployee.ID, employee.ID)
		sCtx.Assert().Equal(expEmployee.PhoneNumber, employee.PhoneNumber)
		sCtx.Assert().Equal(expEmployee.FullName, employee.FullName)
		sCtx.Assert().Equal(expEmployee.Post, employee.Post)
		sCtx.Assert().Equal(expEmployee.Password, employee.Password)
		sCtx.Assert().Equal(expEmployee.CompanyID, employee.CompanyID)
		sCtx.Assert().Equal(expEmployee.TokenExpiredAt, employee.TokenExpiredAt)
		sCtx.Assert().Equal(expEmployee.DateOfBirth, employee.DateOfBirth)
		sCtx.Assert().Equal(expEmployee.RefreshToken, employee.RefreshToken)

		err = c.employeeStorage.Delete(context.TODO(), &dto.DeleteEmployeeRequest{EmployeeID: employee.ID.Int()})
		sCtx.Assert().NoError(err)
	})
}

func (c *EmployeeStorageSuite) Test_IntegrationEmployeeStorage_GetByInfoCardID(t provider.T) {
	t.Title("[GetByInfoCardID] Get employee by info card ID test")
	t.Tags("storage", "postgres", "employee")
	t.Parallel()
	t.WithNewStep("Get employee by info card ID test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		pass, _ := bcrypt.GenerateFromPassword([]byte("21e12"), bcrypt.DefaultCost)
		expiredAt, _ := time.Parse(time.RFC3339, "2008-01-02T00:00:00Z")
		tm, _ := time.Parse(time.RFC3339, "2010-01-02T00:00:00Z")
		expEmployee := &model.Employee{
			ID:             model.ToEmployeeID(c.employeeID),
			PhoneNumber:    "500500500",
			FullName:       "123",
			CompanyID:      model.ToCompanyID(c.companyID),
			Post:           model.ToPostTypeFromInt(1),
			Password:       string(pass),
			RefreshToken:   "974998",
			TokenExpiredAt: &expiredAt,
			DateOfBirth:    &tm,
		}
		request := utils.EmployeeObjectMother{
			EmployeeID: -1,
			InfoCardID: c.infoCardID,
		}.DefaultGetEmployeeByInfoCardIDRequest()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		employee, err := c.employeeStorage.GetByInfoCardID(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(employee)
		sCtx.Assert().Equal(expEmployee.ID, employee.ID)
		sCtx.Assert().Equal(expEmployee.PhoneNumber, employee.PhoneNumber)
		sCtx.Assert().Equal(expEmployee.FullName, employee.FullName)
		sCtx.Assert().Equal(expEmployee.Post, employee.Post)
		sCtx.Assert().Equal(expEmployee.CompanyID, employee.CompanyID)
		sCtx.Assert().Equal(expEmployee.TokenExpiredAt, employee.TokenExpiredAt)
		sCtx.Assert().Equal(expEmployee.DateOfBirth, employee.DateOfBirth)
		sCtx.Assert().Equal(expEmployee.RefreshToken, employee.RefreshToken)
	})
}

func (c *EmployeeStorageSuite) Test_IntegrationEmployeeStorage_Delete(t provider.T) {
	t.Title("[Delete] Delete employee test")
	t.Tags("storage", "postgres", "employee")
	t.Parallel()
	t.WithNewStep("Delete employee test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expEmployee, err := c.employeeStorage.Register(
			ctx,
			utils.AuthObjectMother{
				CompanyID: c.companyID,
			}.DefaultRegisterEmployeeRequest4(),
		)
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(expEmployee)
		request := utils.EmployeeObjectMother{
			EmployeeID: expEmployee.ID.Int(),
			InfoCardID: -1,
		}.DefaultDeleteEmployeeRequest()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err = c.employeeStorage.Delete(ctx, request)

		sCtx.Assert().NoError(err)

		document, err := c.employeeStorage.GetByPhone(
			ctx,
			utils.EmployeeObjectMother{
				EmployeeID: -1,
				InfoCardID: -1,
			}.DefaultGetEmployeeRequest4(),
		)

		sCtx.Assert().Error(err)
		sCtx.Assert().EqualError(err, pgx.ErrNoRows.Error())
		sCtx.Assert().Nil(document)

		err = c.employeeStorage.Delete(ctx, request)

		sCtx.Assert().NoError(err)
	})
}
