package service

import (
	"context"
	"fmt"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/logger"
)

type FieldService interface {
	CreateDocumentField(ctx context.Context, request *dto.CreateDocumentFieldRequest) (*model.Field, error)
	GetDocumentField(ctx context.Context, request *dto.GetDocumentFieldRequest) (*model.Field, error)
	ListDocumentFields(ctx context.Context, request *dto.ListDocumentFieldsRequest) ([]*model.Field, error)
	DeleteDocumentField(ctx context.Context, request *dto.DeleteDocumentFieldRequest) error
}

type fieldServiceImpl struct {
	logger       logger.Interface
	fieldStorage storage.FieldStorage
}

func NewFieldService(logger logger.Interface, fieldStorage storage.FieldStorage) FieldService {
	return &fieldServiceImpl{
		logger:       logger,
		fieldStorage: fieldStorage,
	}
}

func (f *fieldServiceImpl) CreateDocumentField(ctx context.Context, request *dto.CreateDocumentFieldRequest) (*model.Field, error) {
	f.logger.Infof("create field for document with ID %d", request.DocumentID)

	field, err := f.fieldStorage.Create(ctx, request)
	if err != nil {
		f.logger.Errorf("create info card field: %s", err.Error())
		return nil, fmt.Errorf("create info card field: %w", err)
	}

	return field, nil
}

func (f *fieldServiceImpl) GetDocumentField(ctx context.Context, request *dto.GetDocumentFieldRequest) (*model.Field, error) {
	f.logger.Infof("get info card field by document ID %d", request.DocumentID)

	field, err := f.fieldStorage.Get(ctx, request)
	if err != nil {
		f.logger.Errorf("get info card field: %s", err.Error())
		return nil, fmt.Errorf("get info card field: %w", err)
	}

	return field, nil
}

func (f *fieldServiceImpl) ListDocumentFields(ctx context.Context, request *dto.ListDocumentFieldsRequest) ([]*model.Field, error) {
	f.logger.Infof("list info card fields by document ID %d", request.DocumentID)

	fields, err := f.fieldStorage.ListCardFields(ctx, request)
	if err != nil {
		f.logger.Errorf("list info card fields: %s", err.Error())
		return nil, fmt.Errorf("list info card fields: %w", err)
	}

	return fields, nil
}

func (f *fieldServiceImpl) DeleteDocumentField(ctx context.Context, request *dto.DeleteDocumentFieldRequest) error {
	f.logger.Infof("delete info card field by ID %d", request.FieldID)

	err := f.fieldStorage.Delete(ctx, request)
	if err != nil {
		f.logger.Errorf("delete info card field: %s", err.Error())
		return fmt.Errorf("delete info card field: %w", err)
	}

	return nil
}
