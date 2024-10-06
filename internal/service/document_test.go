//go:build unit

package service_test

import (
	"context"
	"fmt"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/internal/model"
	"course/internal/service"
	"course/internal/service/utils"
	"course/internal/storage/mocks"
)

type DocumentSuite struct {
	suite.Suite
}

func (c *DocumentSuite) Test_Document_CreateDocument1(t provider.T) {
	t.Title("[CreateDocument] Incorrect info card ID")
	t.Tags("document")
	t.Parallel()
	t.WithNewStep("Incorrect info card ID", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.DocumentObjectMother{
			InfoCardID: 2,
			DocumentID: 3,
		}.IncorrectInfoCardIDCreateDocumentRequest()

		documentMockStorage := mocks.NewDocumentStorage(t)
		documentMockStorage.
			On("Create", ctx, request).
			Return(nil, fmt.Errorf("incorrect info card ID")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		document, err := service.NewDocumentService(utils.NewMockLogger(), documentMockStorage).CreateDocument(ctx, request)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(document)
	})
}

func (c *DocumentSuite) Test_Document_CreateDocument2(t provider.T) {
	t.Title("[CreateDocument] Success")
	t.Tags("document")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.DocumentObjectMother{
			InfoCardID: 2,
			DocumentID: 3,
		}.DefaultCreateDocumentRequest()
		expDocument := &model.Document{
			ID:           model.ToDocumentID(3),
			SerialNumber: "0001",
			InfoCardID:   model.ToInfoCardID(2),
			Type:         model.ToDocumentTypeFromInt(1),
		}

		documentMockStorage := mocks.NewDocumentStorage(t)
		documentMockStorage.
			On("Create", ctx, request).
			Return(expDocument, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		document, err := service.NewDocumentService(utils.NewMockLogger(), documentMockStorage).CreateDocument(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(document)
		sCtx.Assert().Equal(expDocument, document)
	})
}

func (c *DocumentSuite) Test_Document_GetDocument1(t provider.T) {
	t.Title("[GetDocument] Incorrect document ID")
	t.Tags("document")
	t.Parallel()
	t.WithNewStep("Incorrect document ID", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.DocumentObjectMother{
			InfoCardID: 2,
			DocumentID: 3,
		}.IncorrectDocumentIDGetDocumentRequest()

		documentMockStorage := mocks.NewDocumentStorage(t)
		documentMockStorage.
			On("GetByID", ctx, request).
			Return(nil, fmt.Errorf("incorrect document ID")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		document, err := service.NewDocumentService(utils.NewMockLogger(), documentMockStorage).GetDocument(ctx, request)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(document)
	})
}

func (c *DocumentSuite) Test_Document_GetDocument2(t provider.T) {
	t.Title("[GetDocument] Success")
	t.Tags("document")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.DocumentObjectMother{
			InfoCardID: 2,
			DocumentID: 3,
		}.DefaultGetDocumentRequest()
		expDocument := &model.Document{
			ID:           model.ToDocumentID(3),
			SerialNumber: "0001",
			InfoCardID:   model.ToInfoCardID(2),
			Type:         model.ToDocumentTypeFromInt(1),
		}

		documentMockStorage := mocks.NewDocumentStorage(t)
		documentMockStorage.
			On("GetByID", ctx, request).
			Return(expDocument, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		document, err := service.NewDocumentService(utils.NewMockLogger(), documentMockStorage).GetDocument(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(document)
		sCtx.Assert().Equal(expDocument, document)
	})
}

func (c *DocumentSuite) Test_Document_GetDocumentByInfoCardID1(t provider.T) {
	t.Title("[GetDocumentByInfoCardID] Incorrect info card ID")
	t.Tags("document")
	t.Parallel()
	t.WithNewStep("Incorrect info card ID", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.DocumentObjectMother{
			InfoCardID: 2,
			DocumentID: 3,
		}.IncorrectInfoCardIDGetDocumentByInfoCardIDRequest()

		documentMockStorage := mocks.NewDocumentStorage(t)
		documentMockStorage.
			On("GetByInfoCardID", ctx, request).
			Return(nil, fmt.Errorf("incorrect info card ID")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		document, err := service.NewDocumentService(utils.NewMockLogger(), documentMockStorage).GetDocumentByInfoCard(ctx, request)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(document)
	})
}

func (c *DocumentSuite) Test_Document_GetDocumentByInfoCardID2(t provider.T) {
	t.Title("[GetDocumentByInfoCardID] Success")
	t.Tags("document")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.DocumentObjectMother{
			InfoCardID: 2,
			DocumentID: 3,
		}.DefaultGetDocumentByInfoCardIDRequest()
		expDocument := &model.Document{
			ID:           model.ToDocumentID(3),
			SerialNumber: "0001",
			InfoCardID:   model.ToInfoCardID(2),
			Type:         model.ToDocumentTypeFromInt(1),
		}

		documentMockStorage := mocks.NewDocumentStorage(t)
		documentMockStorage.
			On("GetByInfoCardID", ctx, request).
			Return(expDocument, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		document, err := service.NewDocumentService(utils.NewMockLogger(), documentMockStorage).GetDocumentByInfoCard(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(document)
		sCtx.Assert().Equal(expDocument, document)
	})
}

func (c *DocumentSuite) Test_Document_DeleteDocument1(t provider.T) {
	t.Title("[DeleteDocument] Incorrect document ID")
	t.Tags("document")
	t.Parallel()
	t.WithNewStep("Incorrect document ID", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.DocumentObjectMother{
			InfoCardID: 2,
			DocumentID: 3,
		}.IncorrectDocumentIDDeleteDocumentRequest()

		documentMockStorage := mocks.NewDocumentStorage(t)
		documentMockStorage.
			On("Delete", ctx, request).
			Return(fmt.Errorf("incorrect document ID")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.NewDocumentService(utils.NewMockLogger(), documentMockStorage).DeleteDocument(ctx, request)

		sCtx.Assert().Error(err)
	})
}

func (c *DocumentSuite) Test_Document_DeleteDocument2(t provider.T) {
	t.Title("[DeleteDocument] Success")
	t.Tags("document")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.DocumentObjectMother{
			InfoCardID: 2,
			DocumentID: 3,
		}.DefaultDeleteDocumentRequest()

		documentMockStorage := mocks.NewDocumentStorage(t)
		documentMockStorage.
			On("Delete", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.NewDocumentService(utils.NewMockLogger(), documentMockStorage).DeleteDocument(ctx, request)

		sCtx.Assert().NoError(err)
	})
}
