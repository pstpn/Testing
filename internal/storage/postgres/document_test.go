package postgres_test

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/internal/storage/utils"
)

type DocumentStorageSuite struct {
	suite.Suite

	documentStorage storage.DocumentStorage
	infoCardID      int64
	documentID      int64
}

func (c *DocumentStorageSuite) Test_DocumentStorage_Create(t provider.T) {
	t.Title("[Create] Create document test")
	t.Tags("storage", "postgres", "document")
	t.Parallel()
	t.WithNewStep("Create document test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.DocumentObjectMother{
			InfoCardID: c.infoCardID,
			DocumentID: -1,
		}.DefaultCreateDocumentRequest()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		document, err := c.documentStorage.Create(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(document)
		sCtx.Assert().NotNil(document.ID)
		sCtx.Assert().Equal(request.InfoCardID, document.InfoCardID.Int())
		sCtx.Assert().Equal(request.SerialNumber, document.SerialNumber)
		sCtx.Assert().Equal(request.DocumentType, document.Type.Int())

		err = c.documentStorage.Delete(context.TODO(), &dto.DeleteDocumentRequest{DocumentID: document.ID.Int()})
		sCtx.Assert().NoError(err)
	})
}

func (c *DocumentStorageSuite) Test_DocumentStorage_GetByID(t provider.T) {
	t.Title("[GetByID] Get document by ID test")
	t.Tags("storage", "postgres", "document")
	t.Parallel()
	t.WithNewStep("Get document by ID test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expDocument, err := c.documentStorage.Create(
			ctx,
			utils.DocumentObjectMother{
				InfoCardID: c.infoCardID,
				DocumentID: -1,
			}.DefaultCreateDocumentRequest(),
		)
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(expDocument)
		request := utils.DocumentObjectMother{
			InfoCardID: c.infoCardID,
			DocumentID: expDocument.ID.Int(),
		}.DefaultGetDocumentRequest()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		document, err := c.documentStorage.GetByID(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(document)
		sCtx.Assert().Equal(expDocument.ID, document.ID)
		sCtx.Assert().Equal(expDocument.InfoCardID, document.InfoCardID)
		sCtx.Assert().Equal(expDocument.Type, document.Type)
		sCtx.Assert().Equal(expDocument.SerialNumber, document.SerialNumber)

		err = c.documentStorage.Delete(context.TODO(), &dto.DeleteDocumentRequest{DocumentID: document.ID.Int()})
		sCtx.Assert().NoError(err)
	})
}

func (c *DocumentStorageSuite) Test_DocumentStorage_GetByInfoCardID(t provider.T) {
	t.Title("[GetByInfoCardID] Get document by info card ID test")
	t.Tags("storage", "postgres", "document")
	t.Parallel()
	t.WithNewStep("Get document by info card ID test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expDocument := &model.Document{
			ID:           model.ToDocumentID(c.documentID),
			SerialNumber: "123923",
			InfoCardID:   model.ToInfoCardID(c.infoCardID),
			Type:         model.ToDocumentTypeFromInt(1),
		}
		request := utils.DocumentObjectMother{
			InfoCardID: c.infoCardID,
			DocumentID: -1,
		}.DefaultGetDocumentByInfoCardIDRequest()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		document, err := c.documentStorage.GetByInfoCardID(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(document)
		sCtx.Assert().Equal(expDocument.ID, document.ID)
		sCtx.Assert().Equal(expDocument.InfoCardID, document.InfoCardID)
		sCtx.Assert().Equal(expDocument.Type, document.Type)
		sCtx.Assert().Equal(expDocument.SerialNumber, document.SerialNumber)
	})
}

func (c *DocumentStorageSuite) Test_DocumentStorage_Delete(t provider.T) {
	t.Title("[Delete] Delete document test")
	t.Tags("storage", "postgres", "document")
	t.Parallel()
	t.WithNewStep("Delete document test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expDocument, err := c.documentStorage.Create(
			ctx,
			utils.DocumentObjectMother{
				InfoCardID: c.infoCardID,
				DocumentID: -1,
			}.DefaultCreateDocumentRequest(),
		)
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(expDocument)
		request := utils.DocumentObjectMother{
			InfoCardID: c.infoCardID,
			DocumentID: expDocument.ID.Int(),
		}.DefaultDeleteDocumentRequest()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err = c.documentStorage.Delete(ctx, request)

		sCtx.Assert().NoError(err)

		document, err := c.documentStorage.GetByID(
			ctx,
			utils.DocumentObjectMother{
				InfoCardID: c.infoCardID,
				DocumentID: expDocument.ID.Int(),
			}.DefaultGetDocumentRequest(),
		)

		sCtx.Assert().Error(err)
		sCtx.Assert().EqualError(err, pgx.ErrNoRows.Error())
		sCtx.Assert().Nil(document)

		err = c.documentStorage.Delete(ctx, request)

		sCtx.Assert().NoError(err)
	})
}
