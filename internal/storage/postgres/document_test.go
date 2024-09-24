package postgres_test

import (
	"context"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/internal/model"
	"course/internal/storage/mocks"
	"course/internal/storage/utils"
)

type DocumentStorageSuite struct {
	suite.Suite

	documentMockStorage mocks.DocumentStorage
}

func (c *DocumentStorageSuite) BeforeAll(t provider.T) {
	t.Title("Init document mock storage")
	c.documentMockStorage = *mocks.NewDocumentStorage(t)
	t.Tags("fixture", "document")
}

func (c *DocumentStorageSuite) Test_DocumentStorage_Create(t provider.T) {
	t.Title("[Create] Create document test")
	t.Tags("storage", "postgres", "document")
	t.Parallel()
	t.WithNewStep("Create document test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.DocumentObjectMother{
			InfoCardID: 1,
			DocumentID: -1,
		}.DefaultCreateDocumentRequest()
		expDocument := &model.Document{
			ID:           model.ToDocumentID(1),
			SerialNumber: "123",
			InfoCardID:   model.ToInfoCardID(1),
			Type:         model.ToDocumentTypeFromInt(1),
		}

		c.documentMockStorage.
			On("Create", ctx, request).
			Return(expDocument, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		document, err := c.documentMockStorage.Create(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(document)
		sCtx.Assert().Equal(expDocument, document)
	})
}

func (c *DocumentStorageSuite) Test_DocumentStorage_GetByID(t provider.T) {
	t.Title("[GetByID] Get document by ID test")
	t.Tags("storage", "postgres", "document")
	t.Parallel()
	t.WithNewStep("Get document by ID test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.DocumentObjectMother{
			InfoCardID: 1,
			DocumentID: 1,
		}.DefaultGetDocumentRequest()
		expDocument := &model.Document{
			ID:           model.ToDocumentID(1),
			SerialNumber: "123",
			InfoCardID:   model.ToInfoCardID(1),
			Type:         model.ToDocumentTypeFromInt(1),
		}

		c.documentMockStorage.
			On("GetByID", ctx, request).
			Return(expDocument, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		document, err := c.documentMockStorage.GetByID(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(document)
		sCtx.Assert().Equal(expDocument, document)
	})
}

func (c *DocumentStorageSuite) Test_DocumentStorage_GetByInfoCardID(t provider.T) {
	t.Title("[GetByInfoCardID] Get document by info card ID test")
	t.Tags("storage", "postgres", "document")
	t.Parallel()
	t.WithNewStep("Get document by info card ID test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.DocumentObjectMother{
			InfoCardID: 1,
			DocumentID: 1,
		}.DefaultGetDocumentByInfoCardIDRequest()
		expDocument := &model.Document{
			ID:           model.ToDocumentID(1),
			SerialNumber: "123",
			InfoCardID:   model.ToInfoCardID(1),
			Type:         model.ToDocumentTypeFromInt(1),
		}

		c.documentMockStorage.
			On("GetByInfoCardID", ctx, request).
			Return(expDocument, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		document, err := c.documentMockStorage.GetByInfoCardID(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(document)
		sCtx.Assert().Equal(expDocument, document)
	})
}

func (c *DocumentStorageSuite) Test_DocumentStorage_Delete(t provider.T) {
	t.Title("[Delete] Delete document test")
	t.Tags("storage", "postgres", "document")
	t.Parallel()
	t.WithNewStep("Delete document test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.DocumentObjectMother{
			InfoCardID: 1,
			DocumentID: 1,
		}.DefaultDeleteDocumentRequest()

		c.documentMockStorage.
			On("Delete", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := c.documentMockStorage.Delete(ctx, request)

		sCtx.Assert().NoError(err)
	})
}
