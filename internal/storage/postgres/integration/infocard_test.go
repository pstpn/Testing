//go:build integration

package postgres_test

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/internal/service/dto"
	"course/internal/storage"
	"course/internal/storage/utils"
)

type InfoCardStorageSuite struct {
	suite.Suite

	infoCardStorage storage.InfoCardStorage
	employeeID      int64
}

func (c *InfoCardStorageSuite) Test_IntegrationInfoCardStorage_Create(t provider.T) {
	t.Title("[Create] Create info card test")
	t.Tags("storage", "postgres", "infocard")
	t.Parallel()
	t.WithNewStep("Create info card test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		tm, _ := time.Parse(time.RFC3339, "2012-01-02T00:00:00Z")
		request := utils.InfoCardBuilder{}.
			WithEmployeeID(c.employeeID).
			WithIsConfirmed(false).
			WithCreatedDate(tm).
			ToCreateDTO()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		infoCard, err := c.infoCardStorage.Create(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(infoCard)
		sCtx.Assert().NotNil(infoCard.ID)
		sCtx.Assert().Equal(request.EmployeeID, infoCard.CreatedEmployeeID.Int())
		sCtx.Assert().Equal(request.IsConfirmed, infoCard.IsConfirmed)
		sCtx.Assert().Equal(request.CreatedDate, infoCard.CreatedDate)

		err = c.infoCardStorage.Delete(context.TODO(), &dto.DeleteInfoCardRequest{InfoCardID: infoCard.ID.Int()})
		sCtx.Assert().NoError(err)
	})
}

func (c *InfoCardStorageSuite) Test_IntegrationInfoCardStorage_Validate(t provider.T) {
	t.Title("[Validate] Validate info card test")
	t.Tags("storage", "postgres", "infocard")
	t.Parallel()
	t.WithNewStep("Validate info card test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		tm, _ := time.Parse(time.RFC3339, "2012-01-02T00:00:00Z")
		expInfoCard, err := c.infoCardStorage.Create(
			ctx,
			utils.InfoCardBuilder{}.
				WithEmployeeID(c.employeeID).
				WithIsConfirmed(false).
				WithCreatedDate(tm).
				ToCreateDTO(),
		)
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(expInfoCard)

		request := utils.InfoCardBuilder{}.
			WithInfoCardID(expInfoCard.ID.Int()).
			WithIsConfirmed(true).
			ToValidateDTO()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err = c.infoCardStorage.Validate(ctx, request)

		sCtx.Assert().NoError(err)

		infoCard, err := c.infoCardStorage.GetByID(
			ctx,
			utils.InfoCardBuilder{}.
				WithInfoCardID(expInfoCard.ID.Int()).
				ToGetByIDDTO(),
		)
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(infoCard)
		sCtx.Assert().True(infoCard.IsConfirmed)

		err = c.infoCardStorage.Delete(context.TODO(), &dto.DeleteInfoCardRequest{InfoCardID: infoCard.ID.Int()})
		sCtx.Assert().NoError(err)
	})
}

func (c *InfoCardStorageSuite) Test_IntegrationInfoCardStorage_GetByID(t provider.T) {
	t.Title("[GetByID] Get info card by ID test")
	t.Tags("storage", "postgres", "infocard")
	t.Parallel()
	t.WithNewStep("Get info card by ID test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		tm, _ := time.Parse(time.RFC3339, "2012-01-02T00:00:00Z")
		expInfoCard, err := c.infoCardStorage.Create(
			ctx,
			utils.InfoCardBuilder{}.
				WithEmployeeID(c.employeeID).
				WithIsConfirmed(false).
				WithCreatedDate(tm).
				ToCreateDTO(),
		)
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(expInfoCard)

		request := utils.InfoCardBuilder{}.
			WithInfoCardID(expInfoCard.ID.Int()).
			ToGetByIDDTO()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		infoCard, err := c.infoCardStorage.GetByID(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(infoCard)
		sCtx.Assert().Equal(expInfoCard.ID, infoCard.ID)
		sCtx.Assert().Equal(expInfoCard.CreatedEmployeeID, infoCard.CreatedEmployeeID)
		sCtx.Assert().Equal(expInfoCard.CreatedDate, infoCard.CreatedDate)
		sCtx.Assert().Equal(expInfoCard.IsConfirmed, infoCard.IsConfirmed)

		err = c.infoCardStorage.Delete(context.TODO(), &dto.DeleteInfoCardRequest{InfoCardID: infoCard.ID.Int()})
		sCtx.Assert().NoError(err)
	})
}

func (c *InfoCardStorageSuite) Test_IntegrationInfoCardStorage_GetByEmployeeID(t provider.T) {
	t.Title("[GetByEmployeeID] Get info card by employee ID test")
	t.Tags("storage", "postgres", "infocard")
	t.Parallel()
	t.WithNewStep("Get info card by employee ID test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		tm, _ := time.Parse(time.RFC3339, "2012-01-02T00:00:00Z")
		expInfoCard, err := c.infoCardStorage.Create(
			ctx,
			utils.InfoCardBuilder{}.
				WithEmployeeID(c.employeeID).
				WithIsConfirmed(false).
				WithCreatedDate(tm).
				ToCreateDTO(),
		)
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(expInfoCard)

		request := utils.InfoCardBuilder{}.
			WithInfoCardID(expInfoCard.ID.Int()).
			ToGetByIDDTO()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		infoCard, err := c.infoCardStorage.GetByID(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(infoCard)
		sCtx.Assert().Equal(expInfoCard.ID, infoCard.ID)
		sCtx.Assert().Equal(expInfoCard.CreatedEmployeeID, infoCard.CreatedEmployeeID)
		sCtx.Assert().Equal(expInfoCard.CreatedDate, infoCard.CreatedDate)
		sCtx.Assert().Equal(expInfoCard.IsConfirmed, infoCard.IsConfirmed)

		err = c.infoCardStorage.Delete(context.TODO(), &dto.DeleteInfoCardRequest{InfoCardID: infoCard.ID.Int()})
		sCtx.Assert().NoError(err)
	})
}

func (c *InfoCardStorageSuite) Test_IntegrationInfoCardStorage_List(t provider.T) {
	t.Title("[List] List info cards test")
	t.Tags("storage", "postgres", "infocard")
	t.Parallel()
	t.WithNewStep("List info cards test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		tm, _ := time.Parse(time.RFC3339, "2012-01-02T00:00:00Z")
		expInfoCard, err := c.infoCardStorage.Create(
			ctx,
			utils.InfoCardBuilder{}.
				WithEmployeeID(c.employeeID).
				WithIsConfirmed(false).
				WithCreatedDate(tm).
				ToCreateDTO(),
		)
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(expInfoCard)

		request := utils.InfoCardBuilder{}.
			ToListDTO()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		infoCards, err := c.infoCardStorage.List(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(infoCards)
		sCtx.Assert().GreaterOrEqual(len(infoCards), 1)

		err = c.infoCardStorage.Delete(context.TODO(), &dto.DeleteInfoCardRequest{InfoCardID: expInfoCard.ID.Int()})
		sCtx.Assert().NoError(err)
	})
}

func (c *InfoCardStorageSuite) Test_IntegrationInfoCardStorage_Delete(t provider.T) {
	t.Title("[Delete] Delete info card test")
	t.Tags("storage", "postgres", "infocard")
	t.Parallel()
	t.WithNewStep("Delete info card test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		tm, _ := time.Parse(time.RFC3339, "2012-01-02T00:00:00Z")
		expInfoCard, err := c.infoCardStorage.Create(
			ctx,
			utils.InfoCardBuilder{}.
				WithEmployeeID(c.employeeID).
				WithIsConfirmed(false).
				WithCreatedDate(tm).
				ToCreateDTO(),
		)
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(expInfoCard)

		request := utils.InfoCardBuilder{}.
			WithInfoCardID(expInfoCard.ID.Int()).
			ToDeleteDTO()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err = c.infoCardStorage.Delete(ctx, request)

		sCtx.Assert().NoError(err)

		infoCard, err := c.infoCardStorage.GetByID(
			ctx,
			utils.InfoCardBuilder{}.
				WithInfoCardID(expInfoCard.ID.Int()).
				ToGetByIDDTO(),
		)

		sCtx.Assert().Error(err)
		sCtx.Assert().EqualError(err, pgx.ErrNoRows.Error())
		sCtx.Assert().Nil(infoCard)

		err = c.infoCardStorage.Delete(context.TODO(), &dto.DeleteInfoCardRequest{InfoCardID: expInfoCard.ID.Int()})
		sCtx.Assert().NoError(err)
	})
}
