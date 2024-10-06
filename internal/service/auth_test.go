package service_test

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/mock"

	"course/internal/model"
	"course/internal/service"
	"course/internal/service/dto"
	"course/internal/service/utils"
	"course/internal/storage/mocks"
	"course/pkg/jwt"
)

type AuthSuite struct {
	suite.Suite
}

func (s *AuthSuite) Test_Auth_RegisterEmployee(t provider.T) {
	t.Title("[RegisterEmployee] Incorrect company ID")
	t.Tags("auth", "register")
	t.Parallel()
	t.WithNewStep("Incorrect company ID", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.AuthObjectMother{CompanyID: -1}.IncorrectCompanyIDRegisterEmployeeRequest()

		employeeMockStorage := mocks.NewEmployeeStorage(t)
		employeeMockStorage.
			On("Register", ctx, mock.Anything).
			Return(nil, fmt.Errorf("incorrect company ID")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		tokens, err := service.NewAuthService(
			utils.NewMockLogger(),
			employeeMockStorage,
			mocks.NewInfoCardStorage(t),
			&jwt.Manager{},
			time.Hour,
			time.Hour,
		).RegisterEmployee(ctx, request)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(tokens)
	})
}

func (s *AuthSuite) Test_Auth_LoginEmployee1(t provider.T) {
	t.Title("[LoginEmployee] Incorrect phone number")
	t.Tags("auth", "login")
	t.Parallel()
	t.WithNewStep("Incorrect phone number", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.AuthObjectMother{CompanyID: 1}.IncorrectPhoneNumberLoginEmployeeRequest()

		employeeMockStorage := mocks.NewEmployeeStorage(t)
		employeeMockStorage.
			On("GetByPhone", ctx, &dto.GetEmployeeRequest{PhoneNumber: request.PhoneNumber}).
			Return(nil, fmt.Errorf("incorrect phone number")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		tokens, err := service.NewAuthService(
			utils.NewMockLogger(),
			employeeMockStorage,
			mocks.NewInfoCardStorage(t),
			&jwt.Manager{},
			time.Hour,
			time.Hour,
		).LoginEmployee(ctx, request)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(tokens)
	})
}

func (s *AuthSuite) Test_Auth_LoginEmployee2(t provider.T) {
	t.Title("[LoginEmployee] Incorrect password")
	t.Tags("auth", "login")
	t.Parallel()
	t.WithNewStep("Incorrect password", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.AuthObjectMother{CompanyID: 1}.IncorrectPasswordLoginEmployeeRequest()

		employeeMockStorage := mocks.NewEmployeeStorage(t)
		employeeMockStorage.
			On("GetByPhone", ctx, &dto.GetEmployeeRequest{PhoneNumber: request.PhoneNumber}).
			Return(
				&model.Employee{
					ID:             nil,
					FullName:       "",
					PhoneNumber:    "",
					CompanyID:      nil,
					Post:           nil,
					Password:       "87h9437fh832",
					RefreshToken:   "",
					TokenExpiredAt: nil,
					DateOfBirth:    nil,
				}, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		tokens, err := service.NewAuthService(
			utils.NewMockLogger(),
			employeeMockStorage,
			mocks.NewInfoCardStorage(t),
			&jwt.Manager{},
			time.Hour,
			time.Hour,
		).LoginEmployee(ctx, request)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(tokens)
	})
}

func (s *AuthSuite) Test_Auth_LoginEmployee3(t provider.T) {
	t.Title("[LoginEmployee] Correct request")
	t.Tags("auth", "login")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.AuthObjectMother{CompanyID: 1}.DefaultLoginEmployeeRequest()
		tm := time.Now()
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		sCtx.Assert().NoError(err)

		employeeMockStorage := mocks.NewEmployeeStorage(t)
		employeeMockStorage.
			On("GetByPhone", ctx, &dto.GetEmployeeRequest{PhoneNumber: request.PhoneNumber}).
			Return(
				&model.Employee{
					ID:             model.ToEmployeeID(123),
					FullName:       "123",
					PhoneNumber:    "123",
					CompanyID:      model.ToCompanyID(1),
					Post:           model.ToPostTypeFromInt(1),
					Password:       string(hashedPassword),
					RefreshToken:   "123",
					TokenExpiredAt: &tm,
					DateOfBirth:    &tm,
				}, nil).
			Once()

		infoCardMockStorage := mocks.NewInfoCardStorage(t)
		infoCardMockStorage.
			On("GetByEmployeeID", ctx, &dto.GetInfoCardByEmployeeIDRequest{EmployeeID: 123}).
			Return(
				&model.InfoCard{
					ID:                model.ToInfoCardID(123),
					CreatedEmployeeID: model.ToEmployeeID(123),
					IsConfirmed:       false,
					CreatedDate:       &tm,
				}, nil).
			Once()

		employeeMockStorage.
			On("UpdateRefreshToken", ctx, mock.Anything).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		tokens, err := service.NewAuthService(
			utils.NewMockLogger(),
			employeeMockStorage,
			infoCardMockStorage,
			&jwt.Manager{},
			time.Hour,
			time.Hour,
		).LoginEmployee(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(tokens)
		sCtx.Assert().NotEmpty(tokens.RefreshToken)
		sCtx.Assert().NotEmpty(tokens.RefreshToken)
		sCtx.Assert().False(tokens.IsAdmin)
	})
}
