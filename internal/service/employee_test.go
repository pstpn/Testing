package service_test

import (
	"context"
	"fmt"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/internal/model"
	"course/internal/service"
	"course/internal/service/utils"
	"course/internal/storage/mocks"
)

type EmployeeSuite struct {
	suite.Suite
}

func (s *EmployeeSuite) Test_Employee_GetEmployee1(t provider.T) {
	t.Title("[GetEmployee] Incorrect phone number")
	t.Tags("employee")
	t.Parallel()
	t.WithNewStep("Incorrect phone number", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.EmployeeObjectMother{EmployeeID: 123}.IncorrectPhoneNumberGetEmployeeRequest()

		employeeMockStorage := mocks.NewEmployeeStorage(t)
		employeeMockStorage.
			On("GetByPhone", ctx, request).
			Return(nil, fmt.Errorf("incorrect phone number")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		employee, err := service.NewEmployeeService(utils.NewMockLogger(), employeeMockStorage).GetEmployee(ctx, request)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(employee)
	})
}

func (s *EmployeeSuite) Test_Employee_GetEmployee2(t provider.T) {
	t.Title("[GetEmployee] Success")
	t.Tags("employee")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.EmployeeObjectMother{EmployeeID: 123}.DefaultGetEmployeeRequest()
		tm := time.Now()
		expEmployee := &model.Employee{
			ID:             model.ToEmployeeID(1),
			FullName:       "Stepa Stepan Stepanovich",
			PhoneNumber:    "123",
			CompanyID:      model.ToCompanyID(1),
			Post:           model.ToPostTypeFromInt(1),
			Password:       "OHiuoup98u",
			RefreshToken:   "123321",
			TokenExpiredAt: &tm,
			DateOfBirth:    &tm,
		}

		employeeMockStorage := mocks.NewEmployeeStorage(t)
		employeeMockStorage.
			On("GetByPhone", ctx, request).
			Return(expEmployee, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		employee, err := service.NewEmployeeService(utils.NewMockLogger(), employeeMockStorage).GetEmployee(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(employee)
		sCtx.Assert().Equal(expEmployee, employee)
	})
}

func (s *EmployeeSuite) Test_Employee_DeleteEmployee1(t provider.T) {
	t.Title("[DeleteEmployee] Incorrect info card ID")
	t.Tags("employee")
	t.Parallel()
	t.WithNewStep("Incorrect info card ID", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.EmployeeObjectMother{EmployeeID: 123}.IncorrectEmployeeIDDeleteEmployeeRequest()

		employeeMockStorage := mocks.NewEmployeeStorage(t)
		employeeMockStorage.
			On("Delete", ctx, request).
			Return(fmt.Errorf("incorrect info card ID")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.NewEmployeeService(utils.NewMockLogger(), employeeMockStorage).DeleteEmployee(ctx, request)

		sCtx.Assert().Error(err)
	})
}

func (s *EmployeeSuite) Test_Employee_DeleteEmployee2(t provider.T) {
	t.Title("[DeleteEmployee] Success")
	t.Tags("employee")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.EmployeeObjectMother{EmployeeID: 123}.DefaultDeleteEmployeeRequest()

		employeeMockStorage := mocks.NewEmployeeStorage(t)
		employeeMockStorage.
			On("Delete", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.NewEmployeeService(utils.NewMockLogger(), employeeMockStorage).DeleteEmployee(ctx, request)

		sCtx.Assert().NoError(err)
	})
}
