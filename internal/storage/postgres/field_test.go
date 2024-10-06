//go:build unit

package postgres_test

import (
	"context"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/internal/model"
	"course/internal/storage/mocks"
	"course/internal/storage/utils"
)

type FieldStorageSuite struct {
	suite.Suite

	fieldMockStorage mocks.FieldStorage
}

func (c *FieldStorageSuite) BeforeAll(t provider.T) {
	t.Title("Init field mock storage")
	c.fieldMockStorage = *mocks.NewFieldStorage(t)
	t.Tags("fixture", "field")
}

func (c *FieldStorageSuite) Test_FieldStorage_Create(t provider.T) {
	t.Title("[Create] Create document field test")
	t.Tags("storage", "postgres", "field")
	t.Parallel()
	t.WithNewStep("Create document field test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.FieldObjectMother{
			DocumentID: 1,
			FieldID:    1,
		}.DefaultCreateDocumentFieldRequest3()
		expField := &model.Field{
			ID:         model.ToFieldID(1),
			DocumentID: model.ToDocumentID(1),
			Type:       model.ToFieldTypeFromInt(1),
			Value:      "123",
		}

		c.fieldMockStorage.
			On("Create", ctx, request).
			Return(expField, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		field, err := c.fieldMockStorage.Create(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(field)
		sCtx.Assert().Equal(expField, field)
	})
}

func (c *FieldStorageSuite) Test_FieldStorage_Get(t provider.T) {
	t.Title("[Get] Get document field test")
	t.Tags("storage", "postgres", "field")
	t.Parallel()
	t.WithNewStep("Get document field test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.FieldObjectMother{
			DocumentID: 1,
			FieldID:    1,
		}.DefaultGetDocumentFieldRequest3()
		expField := &model.Field{
			ID:         model.ToFieldID(1),
			DocumentID: model.ToDocumentID(1),
			Type:       model.ToFieldTypeFromInt(1),
			Value:      "123",
		}

		c.fieldMockStorage.
			On("Get", ctx, request).
			Return(expField, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		field, err := c.fieldMockStorage.Get(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(field)
		sCtx.Assert().Equal(expField, field)
	})
}

func (c *FieldStorageSuite) Test_FieldStorage_ListCardFields(t provider.T) {
	t.Title("[ListCardFields] List document fields test")
	t.Tags("storage", "postgres", "field")
	t.Parallel()
	t.WithNewStep("List document fields test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.FieldObjectMother{
			DocumentID: 1,
			FieldID:    1,
		}.DefaultListDocumentFieldsRequest()
		expFields := []*model.Field{
			{
				ID:         model.ToFieldID(1),
				DocumentID: model.ToDocumentID(1),
				Type:       model.ToFieldTypeFromInt(1),
				Value:      "123",
			},
		}

		c.fieldMockStorage.
			On("ListCardFields", ctx, request).
			Return(expFields, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		fields, err := c.fieldMockStorage.ListCardFields(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(fields)
		sCtx.Assert().Equal(expFields, fields)
	})
}

func (c *FieldStorageSuite) Test_FieldStorage_Delete(t provider.T) {
	t.Title("[Delete] Delete document field test")
	t.Tags("storage", "postgres", "field")
	t.Parallel()
	t.WithNewStep("Delete document field test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.FieldObjectMother{
			DocumentID: 1,
			FieldID:    1,
		}.DefaultDeleteDocumentFieldRequest()

		c.fieldMockStorage.
			On("Delete", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := c.fieldMockStorage.Delete(ctx, request)

		sCtx.Assert().NoError(err)
	})
}
