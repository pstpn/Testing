package postgres_test

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/internal/service/dto"
	"course/internal/storage"
	"course/internal/storage/utils"
)

type FieldStorageSuite struct {
	suite.Suite

	fieldStorage storage.FieldStorage
	documentID   int64
}

func (c *FieldStorageSuite) Test_IntegrationFieldStorage_Create(t provider.T) {
	t.Title("[Create] Create document field test")
	t.Tags("storage", "postgres", "document", "field")
	t.Parallel()
	t.WithNewStep("Create document field test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.FieldObjectMother{
			DocumentID: c.documentID,
			FieldID:    -1,
		}.DefaultCreateDocumentFieldRequest3()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		field, err := c.fieldStorage.Create(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(field)
		sCtx.Assert().NotNil(field.ID)
		sCtx.Assert().Equal(request.DocumentID, field.DocumentID.Int())
		sCtx.Assert().Equal(request.Type, field.Type.Int())
		sCtx.Assert().Equal(request.Value, field.Value)

		err = c.fieldStorage.Delete(context.TODO(), &dto.DeleteDocumentFieldRequest{FieldID: field.ID.Int()})
		sCtx.Assert().NoError(err)
	})
}

func (c *FieldStorageSuite) Test_IntegrationFieldStorage_Get(t provider.T) {
	t.Title("[Get] Get document field test")
	t.Tags("storage", "postgres", "document", "field")
	t.Parallel()
	t.WithNewStep("Get document field test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expField, err := c.fieldStorage.Create(
			ctx,
			utils.FieldObjectMother{
				DocumentID: c.documentID,
				FieldID:    -1,
			}.DefaultCreateDocumentFieldRequest4(),
		)
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(expField)
		request := utils.FieldObjectMother{
			DocumentID: c.documentID,
			FieldID:    -1,
		}.DefaultGetDocumentFieldRequest4()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		field, err := c.fieldStorage.Get(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(field)
		sCtx.Assert().Equal(expField.ID, field.ID)
		sCtx.Assert().Equal(expField.DocumentID, field.DocumentID)
		sCtx.Assert().Equal(expField.Type, field.Type)
		sCtx.Assert().Equal(expField.Value, field.Value)

		err = c.fieldStorage.Delete(context.TODO(), &dto.DeleteDocumentFieldRequest{FieldID: field.ID.Int()})
		sCtx.Assert().NoError(err)
	})
}

func (c *FieldStorageSuite) Test_IntegrationFieldStorage_List(t provider.T) {
	t.Title("[List] List document fields test")
	t.Tags("storage", "postgres", "document", "field")
	t.Parallel()
	t.WithNewStep("List document fields test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expField, err := c.fieldStorage.Create(
			ctx,
			utils.FieldObjectMother{
				DocumentID: c.documentID,
				FieldID:    -1,
			}.DefaultCreateDocumentFieldRequest5(),
		)
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(expField)
		request := utils.FieldObjectMother{
			DocumentID: c.documentID,
			FieldID:    -1,
		}.DefaultListDocumentFieldsRequest()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		fields, err := c.fieldStorage.ListCardFields(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(fields)
		sCtx.Assert().GreaterOrEqual(len(fields), 1)

		err = c.fieldStorage.Delete(context.TODO(), &dto.DeleteDocumentFieldRequest{FieldID: expField.ID.Int()})
		sCtx.Assert().NoError(err)
	})
}

func (c *FieldStorageSuite) Test_IntegrationFieldStorage_Delete(t provider.T) {
	t.Title("[Delete] Delete document field test")
	t.Tags("storage", "postgres", "document", "field")
	t.Parallel()
	t.WithNewStep("Delete document field test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expField, err := c.fieldStorage.Create(
			ctx,
			utils.FieldObjectMother{
				DocumentID: c.documentID,
				FieldID:    -1,
			}.DefaultCreateDocumentFieldRequest6(),
		)
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(expField)
		request := utils.FieldObjectMother{
			DocumentID: -1,
			FieldID:    expField.ID.Int(),
		}.DefaultDeleteDocumentFieldRequest()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err = c.fieldStorage.Delete(ctx, request)

		sCtx.Assert().NoError(err)

		field, err := c.fieldStorage.Get(
			ctx,
			utils.FieldObjectMother{
				DocumentID: c.documentID,
				FieldID:    -1,
			}.DefaultGetDocumentFieldRequest6(),
		)

		sCtx.Assert().Error(err)
		sCtx.Assert().EqualError(err, pgx.ErrNoRows.Error())
		sCtx.Assert().Nil(field)

		err = c.fieldStorage.Delete(context.TODO(), &dto.DeleteDocumentFieldRequest{FieldID: expField.ID.Int()})
		sCtx.Assert().NoError(err)
	})
}
