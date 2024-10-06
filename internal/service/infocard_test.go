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

type InfoCardSuite struct {
	suite.Suite
}

func (s *InfoCardSuite) Test_InfoCard_CreateInfoCard1(t provider.T) {
	t.Title("[CreateInfoCard] Incorrect employee ID")
	t.Tags("infocard")
	t.Parallel()
	t.WithNewStep("Incorrect employee ID", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		tm, _ := time.Parse(time.RFC3339, "2012-01-02T00:00:00Z")
		request := utils.InfoCardBuilder{}.
			WithEmployeeID(-11).
			WithIsConfirmed(false).
			WithCreatedDate(tm).
			ToCreateDTO()

		infoCardMockStorage := mocks.NewInfoCardStorage(t)
		infoCardMockStorage.
			On("Create", ctx, request).
			Return(nil, fmt.Errorf("incorrect employee ID")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		infoCard, err := service.NewInfoCardService(utils.NewMockLogger(), infoCardMockStorage).CreateInfoCard(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Nil(infoCard)
	})
}

func (s *InfoCardSuite) Test_InfoCard_CreateInfoCard2(t provider.T) {
	t.Title("[CreateInfoCard] Success")
	t.Tags("infocard")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		tm, _ := time.Parse(time.RFC3339, "2012-01-02T00:00:00Z")
		request := utils.InfoCardBuilder{}.
			WithEmployeeID(123).
			WithIsConfirmed(false).
			WithCreatedDate(tm).
			ToCreateDTO()
		expInfoCard := &model.InfoCard{
			ID:                model.ToInfoCardID(1313),
			CreatedEmployeeID: model.ToEmployeeID(123),
			IsConfirmed:       false,
			CreatedDate:       &tm,
		}

		infoCardMockStorage := mocks.NewInfoCardStorage(t)
		infoCardMockStorage.
			On("Create", ctx, request).
			Return(expInfoCard, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		infoCard, err := service.NewInfoCardService(utils.NewMockLogger(), infoCardMockStorage).CreateInfoCard(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(infoCard)
		sCtx.Assert().Equal(expInfoCard, infoCard)
	})
}

func (s *InfoCardSuite) Test_InfoCard_ValidateInfoCard1(t provider.T) {
	t.Title("[ValidateInfoCard] Incorrect info card ID")
	t.Tags("infocard")
	t.Parallel()
	t.WithNewStep("Incorrect info card ID", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.InfoCardBuilder{}.
			WithInfoCardID(-423).
			WithIsConfirmed(true).
			ToValidateDTO()

		infoCardMockStorage := mocks.NewInfoCardStorage(t)
		infoCardMockStorage.
			On("Validate", ctx, request).
			Return(fmt.Errorf("incorrect info card ID")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.NewInfoCardService(utils.NewMockLogger(), infoCardMockStorage).ValidateInfoCard(ctx, request)

		sCtx.Assert().Error(err)
	})
}

func (s *InfoCardSuite) Test_InfoCard_ValidateInfoCard2(t provider.T) {
	t.Title("[ValidateInfoCard] Success")
	t.Tags("infocard")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.InfoCardBuilder{}.
			WithEmployeeID(483).
			WithIsConfirmed(true).
			ToValidateDTO()

		infoCardMockStorage := mocks.NewInfoCardStorage(t)
		infoCardMockStorage.
			On("Validate", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.NewInfoCardService(utils.NewMockLogger(), infoCardMockStorage).ValidateInfoCard(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (s *InfoCardSuite) Test_InfoCard_GetInfoCard1(t provider.T) {
	t.Title("[GetInfoCard] Incorrect info card ID")
	t.Tags("infocard")
	t.Parallel()
	t.WithNewStep("Incorrect info card ID", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.InfoCardBuilder{}.
			WithInfoCardID(-11).
			ToGetByIDDTO()

		infoCardMockStorage := mocks.NewInfoCardStorage(t)
		infoCardMockStorage.
			On("GetByID", ctx, request).
			Return(nil, fmt.Errorf("incorrect info card ID")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		infoCard, err := service.NewInfoCardService(utils.NewMockLogger(), infoCardMockStorage).GetInfoCard(ctx, request)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(infoCard)
	})
}

func (s *InfoCardSuite) Test_InfoCard_GetInfoCard2(t provider.T) {
	t.Title("[GetInfoCard] Success")
	t.Tags("infocard")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		tm, _ := time.Parse(time.RFC3339, "2012-01-02T00:00:00Z")
		request := utils.InfoCardBuilder{}.
			WithInfoCardID(123).
			ToGetByIDDTO()
		expInfoCard := &model.InfoCard{
			ID:                model.ToInfoCardID(123),
			CreatedEmployeeID: model.ToEmployeeID(123),
			IsConfirmed:       false,
			CreatedDate:       &tm,
		}

		infoCardMockStorage := mocks.NewInfoCardStorage(t)
		infoCardMockStorage.
			On("GetByID", ctx, request).
			Return(expInfoCard, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		infoCard, err := service.NewInfoCardService(utils.NewMockLogger(), infoCardMockStorage).GetInfoCard(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(infoCard)
		sCtx.Assert().Equal(expInfoCard, infoCard)
	})
}

func (s *InfoCardSuite) Test_InfoCard_ListInfoCards1(t provider.T) {
	t.Title("[ListInfoCards] Incorrect request")
	t.Tags("infocard")
	t.Parallel()
	t.WithNewStep("Incorrect request", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.InfoCardBuilder{}.
			ToIncorrectListDTO()

		infoCardMockStorage := mocks.NewInfoCardStorage(t)
		infoCardMockStorage.
			On("List", ctx, request).
			Return(nil, fmt.Errorf("incorrect request")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		infoCards, err := service.NewInfoCardService(utils.NewMockLogger(), infoCardMockStorage).ListInfoCards(ctx, request)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(infoCards)
	})
}

func (s *InfoCardSuite) Test_InfoCard_ListInfoCards2(t provider.T) {
	t.Title("[ListInfoCards] Success")
	t.Tags("infocard")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		tm, _ := time.Parse(time.RFC3339, "2012-01-02T00:00:00Z")
		request := utils.InfoCardBuilder{}.
			ToListDTO()
		expInfoCards := []*model.FullInfoCard{
			{
				ID:                model.ToInfoCardID(123),
				CreatedEmployeeID: model.ToEmployeeID(123),
				IsConfirmed:       false,
				CreatedDate:       &tm,
				FullName:          "Stepaaaaa",
				PhoneNumber:       "88412",
				CompanyID:         model.ToCompanyID(123),
				Post:              "Test",
				DateOfBirth:       &tm,
			},
		}

		infoCardMockStorage := mocks.NewInfoCardStorage(t)
		infoCardMockStorage.
			On("List", ctx, request).
			Return(expInfoCards, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		infoCards, err := service.NewInfoCardService(utils.NewMockLogger(), infoCardMockStorage).ListInfoCards(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(infoCards)
		sCtx.Assert().Equal(expInfoCards, infoCards)
	})
}

func (s *InfoCardSuite) Test_InfoCard_DeleteInfoCard1(t provider.T) {
	t.Title("[DeleteInfoCard] Incorrect info card ID")
	t.Tags("infocard")
	t.Parallel()
	t.WithNewStep("Incorrect info card ID", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.InfoCardBuilder{}.
			WithInfoCardID(-123).
			ToDeleteDTO()

		infoCardMockStorage := mocks.NewInfoCardStorage(t)
		infoCardMockStorage.
			On("Delete", ctx, request).
			Return(fmt.Errorf("incorrect info card ID")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.NewInfoCardService(utils.NewMockLogger(), infoCardMockStorage).DeleteInfoCard(ctx, request)

		sCtx.Assert().Error(err)
	})
}

func (s *InfoCardSuite) Test_InfoCard_DeleteInfoCard2(t provider.T) {
	t.Title("[DeleteInfoCard] Success")
	t.Tags("infocard")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.InfoCardBuilder{}.
			WithInfoCardID(123).
			ToDeleteDTO()

		infoCardMockStorage := mocks.NewInfoCardStorage(t)
		infoCardMockStorage.
			On("Delete", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.NewInfoCardService(utils.NewMockLogger(), infoCardMockStorage).DeleteInfoCard(ctx, request)

		sCtx.Assert().NoError(err)
	})
}
