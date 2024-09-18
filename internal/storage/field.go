package storage

import (
	"context"

	"course/internal/model"
	"course/internal/service/dto"
)

//go:generate mockery --name=FieldStorage
type FieldStorage interface {
	Create(ctx context.Context, request *dto.CreateDocumentFieldRequest) (*model.Field, error)
	Get(ctx context.Context, request *dto.GetDocumentFieldRequest) (*model.Field, error)
	ListCardFields(ctx context.Context, request *dto.ListDocumentFieldsRequest) ([]*model.Field, error)
	Delete(ctx context.Context, request *dto.DeleteDocumentFieldRequest) error
}
