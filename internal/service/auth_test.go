package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/internal/service"
	"course/internal/service/utils"
	"course/internal/storage/postgres"
	"course/pkg/jwt"
)

type AuthSuite struct {
	suite.Suite

	authService  service.AuthService
	objectMother *utils.ObjectMother
}

func (s *AuthSuite) Test_Auth_RegisterEmployee1(t provider.T) {
	t.Title("[RegisterEmployee] Incorrect company ID")
	t.Tags("auth", "register")
	t.Parallel()
	t.WithNewStep("Incorrect company ID", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := s.objectMother.IncorrectCompanyIDRegisterEmployeeRequest()
		sCtx.WithNewParameters("ctx", ctx, "request", request)

		tokens, err := s.authService.RegisterEmployee(ctx, request)

		sCtx.Require().Error(err)
		sCtx.Require().Nil(tokens)
	})
}

func (s *AuthSuite) Test_Auth_RegisterEmployee2(t provider.T) {
	t.Title("[RegisterEmployee] Correct request")
	t.Tags("auth", "register")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := s.objectMother.DefaultRegisterEmployeeRequest()
		sCtx.WithNewParameters("ctx", ctx, "request", request)

		tokens, err := s.authService.RegisterEmployee(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(tokens)
		sCtx.Assert().NotEmpty(tokens.RefreshToken)
		sCtx.Assert().NotEmpty(tokens.RefreshToken)
		sCtx.Assert().False(tokens.IsAdmin)
	})
}

func (s *AuthSuite) Test_Auth_LoginEmployee1(t provider.T) {
	t.Title("[LoginEmployee] Incorrect phone number")
	t.Tags("auth", "login")
	t.Parallel()
	t.WithNewStep("Incorrect phone number", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := s.objectMother.IncorrectPhoneNumberLoginEmployeeRequest()
		sCtx.WithNewParameters("ctx", ctx, "request", request)

		tokens, err := s.authService.LoginEmployee(ctx, request)

		sCtx.Require().Error(err)
		sCtx.Require().Nil(tokens)
	})
}

func (s *AuthSuite) Test_Auth_LoginEmployee2(t provider.T) {
	t.Title("[LoginEmployee] Incorrect password")
	t.Tags("auth", "login")
	t.Parallel()
	t.WithNewStep("Incorrect password", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := s.objectMother.IncorrectPasswordLoginEmployeeRequest()
		sCtx.WithNewParameters("ctx", ctx, "request", request)

		tokens, err := s.authService.LoginEmployee(ctx, request)

		sCtx.Require().Error(err)
		sCtx.Require().Nil(tokens)
	})
}

func (s *AuthSuite) Test_Auth_LoginEmployee3(t provider.T) {
	t.Title("[LoginEmployee] Correct request")
	t.Tags("auth", "login")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := s.objectMother.DefaultLoginEmployeeRequest()
		sCtx.WithNewParameters("ctx", ctx, "request", request)

		tokens, err := s.authService.LoginEmployee(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(tokens)
		sCtx.Assert().NotEmpty(tokens.RefreshToken)
		sCtx.Assert().NotEmpty(tokens.RefreshToken)
		sCtx.Assert().False(tokens.IsAdmin)
	})
}

func TestRunner(t *testing.T) {
	db, ids := utils.NewTestStorage()
	defer utils.DropTestStorage(db)

	tm, err := jwt.NewManager("test")
	if err != nil {
		panic(err)
	}

	suite.RunSuite(t, &AuthSuite{
		authService: service.NewAuthService(
			utils.NewMockLogger(),
			postgres.NewEmployeeStorage(db),
			postgres.NewInfoCardStorage(db),
			tm,
			time.Hour,
			time.Hour,
		),
		objectMother: &utils.ObjectMother{
			CompanyID: ids["companyID"],
		},
	})
}
