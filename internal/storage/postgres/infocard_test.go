//go:build unit

package postgres_test

import (
	"context"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/internal/model"
	"course/internal/storage/mocks"
	"course/internal/storage/utils"
)

type InfoCardStorageSuite struct {
	suite.Suite

	infoCardMockStorage mocks.InfoCardStorage
}

func (c *InfoCardStorageSuite) BeforeAll(t provider.T) {
	t.Title("Init infocard mock storage")
	c.infoCardMockStorage = *mocks.NewInfoCardStorage(t)
	t.Tags("fixture", "infocard")
}

func (c *InfoCardStorageSuite) Test_InfoCardStorage_Create(t provider.T) {
	t.Title("[Create] Create info card test")
	t.Tags("storage", "postgres", "infocard")
	t.Parallel()
	t.WithNewStep("Create info card test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		tm, _ := time.Parse(time.RFC3339, "2012-01-02T00:00:00Z")
		request := utils.InfoCardBuilder{}.
			WithEmployeeID(1).
			WithIsConfirmed(false).
			WithCreatedDate(tm).
			ToCreateDTO()
		expInfoCard := &model.InfoCard{
			ID:                model.ToInfoCardID(1),
			CreatedEmployeeID: model.ToEmployeeID(1),
			IsConfirmed:       false,
			CreatedDate:       &tm,
		}

		c.infoCardMockStorage.
			On("Create", ctx, request).
			Return(expInfoCard, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		infoCard, err := c.infoCardMockStorage.Create(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(infoCard)
		sCtx.Assert().Equal(expInfoCard, infoCard)
	})
}

func (c *InfoCardStorageSuite) Test_InfoCardStorage_Validate(t provider.T) {
	t.Title("[Validate] Validate info card test")
	t.Tags("storage", "postgres", "infocard")
	t.Parallel()
	t.WithNewStep("Validate info card test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.InfoCardBuilder{}.
			WithInfoCardID(1).
			WithIsConfirmed(true).
			ToValidateDTO()

		c.infoCardMockStorage.
			On("Validate", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := c.infoCardMockStorage.Validate(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (c *InfoCardStorageSuite) Test_InfoCardStorage_GetByID(t provider.T) {
	t.Title("[GetByID] Get info card by ID test")
	t.Tags("storage", "postgres", "infocard")
	t.Parallel()
	t.WithNewStep("Get info card by ID test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.InfoCardBuilder{}.
			WithInfoCardID(1).
			ToGetByIDDTO()
		expInfoCard := &model.InfoCard{
			ID:                model.ToInfoCardID(1),
			CreatedEmployeeID: model.ToEmployeeID(1),
			IsConfirmed:       false,
			CreatedDate:       nil,
		}

		c.infoCardMockStorage.
			On("GetByID", ctx, request).
			Return(expInfoCard, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		infoCard, err := c.infoCardMockStorage.GetByID(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(infoCard)
		sCtx.Assert().Equal(expInfoCard, infoCard)
	})
}

func (c *InfoCardStorageSuite) Test_InfoCardStorage_GetByEmployeeID(t provider.T) {
	t.Title("[GetByEmployeeID] Get info card by employee ID test")
	t.Tags("storage", "postgres", "infocard")
	t.Parallel()
	t.WithNewStep("Get info card by employee ID test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.InfoCardBuilder{}.
			WithInfoCardID(1).
			ToGetByEmployeeIDDTO()
		expInfoCard := &model.InfoCard{
			ID:                model.ToInfoCardID(1),
			CreatedEmployeeID: model.ToEmployeeID(1),
			IsConfirmed:       false,
			CreatedDate:       nil,
		}

		c.infoCardMockStorage.
			On("GetByEmployeeID", ctx, request).
			Return(expInfoCard, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		infoCard, err := c.infoCardMockStorage.GetByEmployeeID(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(infoCard)
		sCtx.Assert().Equal(expInfoCard, infoCard)
	})
}

func (c *InfoCardStorageSuite) Test_InfoCardStorage_List(t provider.T) {
	t.Title("[List] List info cards test")
	t.Tags("storage", "postgres", "infocard")
	t.Parallel()
	t.WithNewStep("List info cards test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.InfoCardBuilder{}.
			ToListDTO()
		expInfoCards := []*model.FullInfoCard{
			{
				ID:                model.ToInfoCardID(1),
				CreatedEmployeeID: model.ToEmployeeID(1),
				IsConfirmed:       false,
				CreatedDate:       nil,
				FullName:          "123",
				PhoneNumber:       "123",
				CompanyID:         model.ToCompanyID(1),
				Post:              "123",
				DateOfBirth:       nil,
			},
		}

		c.infoCardMockStorage.
			On("List", ctx, request).
			Return(expInfoCards, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		infoCards, err := c.infoCardMockStorage.List(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(infoCards)
		sCtx.Assert().Equal(expInfoCards, infoCards)
	})
}

func (c *InfoCardStorageSuite) Test_InfoCardStorage_Delete(t provider.T) {
	t.Title("[Delete] Delete info card test")
	t.Tags("storage", "postgres", "infocard")
	t.Parallel()
	t.WithNewStep("Delete info card test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.InfoCardBuilder{}.
			WithInfoCardID(1).
			ToDeleteDTO()

		c.infoCardMockStorage.
			On("Delete", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := c.infoCardMockStorage.Delete(ctx, request)

		sCtx.Assert().NoError(err)
	})
}
