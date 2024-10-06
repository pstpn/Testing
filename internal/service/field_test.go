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

type FieldSuite struct {
	suite.Suite
}

func (s *FieldSuite) Test_Field_CreateDocumentField1(t provider.T) {
	t.Title("[CreateDocumentField] Incorrect document ID")
	t.Tags("field")
	t.Parallel()
	t.WithNewStep("Incorrect document ID", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.FieldObjectMother{DocumentID: 123}.IncorrectDocumentIDCreateDocumentFieldRequest()

		fieldMockStorage := mocks.NewFieldStorage(t)
		fieldMockStorage.
			On("Create", ctx, request).
			Return(nil, fmt.Errorf("incorrect document ID")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		field, err := service.NewFieldService(utils.NewMockLogger(), fieldMockStorage).CreateDocumentField(ctx, request)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(field)
	})
}

func (s *FieldSuite) Test_Field_CreateDocumentField2(t provider.T) {
	t.Title("[CreateDocumentField] Success")
	t.Tags("field")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.FieldObjectMother{DocumentID: 123}.DefaultCreateDocumentFieldRequest1()
		expField := &model.Field{
			ID:         model.ToFieldID(191),
			DocumentID: model.ToDocumentID(123),
			Type:       model.ToFieldTypeFromInt(1),
			Value:      "123",
		}

		fieldMockStorage := mocks.NewFieldStorage(t)
		fieldMockStorage.
			On("Create", ctx, request).
			Return(expField, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		field, err := service.NewFieldService(utils.NewMockLogger(), fieldMockStorage).CreateDocumentField(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(field)
		sCtx.Assert().Equal(expField, field)
	})
}

func (s *FieldSuite) Test_Field_GetDocumentField1(t provider.T) {
	t.Title("[GetDocumentField] Incorrect document ID")
	t.Tags("field")
	t.Parallel()
	t.WithNewStep("Incorrect document ID", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.FieldObjectMother{DocumentID: 123}.IncorrectDocumentIDGetDocumentFieldRequest()

		fieldMockStorage := mocks.NewFieldStorage(t)
		fieldMockStorage.
			On("Get", ctx, request).
			Return(nil, fmt.Errorf("incorrect document ID")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		field, err := service.NewFieldService(utils.NewMockLogger(), fieldMockStorage).GetDocumentField(ctx, request)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(field)
	})
}

func (s *FieldSuite) Test_Field_GetDocumentField2(t provider.T) {
	t.Title("[GetDocumentField] Success")
	t.Tags("field")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.FieldObjectMother{DocumentID: 123}.DefaultGetDocumentFieldRequest2()
		expField := &model.Field{
			ID:         model.ToFieldID(191),
			DocumentID: model.ToDocumentID(123),
			Type:       model.ToFieldTypeFromInt(1),
			Value:      "123",
		}

		fieldMockStorage := mocks.NewFieldStorage(t)
		fieldMockStorage.
			On("Get", ctx, request).
			Return(expField, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		field, err := service.NewFieldService(utils.NewMockLogger(), fieldMockStorage).GetDocumentField(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(field)
		sCtx.Assert().Equal(expField, field)
	})
}

func (s *FieldSuite) Test_Field_ListDocumentFields1(t provider.T) {
	t.Title("[ListDocumentFields] Incorrect document ID")
	t.Tags("field")
	t.Parallel()
	t.WithNewStep("Incorrect document ID", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.FieldObjectMother{DocumentID: 123}.IncorrectDocumentIDListDocumentFieldsRequest()

		fieldMockStorage := mocks.NewFieldStorage(t)
		fieldMockStorage.
			On("ListCardFields", ctx, request).
			Return(nil, fmt.Errorf("incorrect document ID")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		field, err := service.NewFieldService(utils.NewMockLogger(), fieldMockStorage).ListDocumentFields(ctx, request)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(field)
	})
}

func (s *FieldSuite) Test_Field_ListDocumentFields2(t provider.T) {
	t.Title("[ListDocumentFields] Success")
	t.Tags("field")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.FieldObjectMother{DocumentID: 123}.DefaultListDocumentFieldsRequest()
		expFields := []*model.Field{{
			ID:         model.ToFieldID(191),
			DocumentID: model.ToDocumentID(123),
			Type:       model.ToFieldTypeFromInt(1),
			Value:      "123",
		}}

		fieldMockStorage := mocks.NewFieldStorage(t)
		fieldMockStorage.
			On("ListCardFields", ctx, request).
			Return(expFields, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		fields, err := service.NewFieldService(utils.NewMockLogger(), fieldMockStorage).ListDocumentFields(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(fields)
		sCtx.Assert().Len(fields, 1)
		sCtx.Assert().Equal(expFields[0], fields[0])
	})
}

func (s *FieldSuite) Test_Field_DeleteDocumentField1(t provider.T) {
	t.Title("[DeleteDocumentField] Incorrect field ID")
	t.Tags("field")
	t.Parallel()
	t.WithNewStep("Incorrect field ID", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.FieldObjectMother{DocumentID: 123}.IncorrectFieldIDDeleteDocumentFieldRequest()

		fieldMockStorage := mocks.NewFieldStorage(t)
		fieldMockStorage.
			On("Delete", ctx, request).
			Return(fmt.Errorf("incorrect field ID")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.NewFieldService(utils.NewMockLogger(), fieldMockStorage).DeleteDocumentField(ctx, request)

		sCtx.Assert().Error(err)
	})
}

func (s *FieldSuite) Test_Field_DeleteDocumentField2(t provider.T) {
	t.Title("[DeleteDocumentField] Success")
	t.Tags("field")
	t.Parallel()
	t.WithNewStep("Delete", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.FieldObjectMother{DocumentID: 123}.DefaultDeleteDocumentFieldRequest()

		fieldMockStorage := mocks.NewFieldStorage(t)
		fieldMockStorage.
			On("Delete", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.NewFieldService(utils.NewMockLogger(), fieldMockStorage).DeleteDocumentField(ctx, request)

		sCtx.Assert().NoError(err)
	})
}
