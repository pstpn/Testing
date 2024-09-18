package storage

import (
	"context"

	"course/internal/model"
	"course/internal/service/dto"
)

//go:generate mockery --name=DocumentStorage
type DocumentStorage interface {
	Create(ctx context.Context, request *dto.CreateDocumentRequest) (*model.Document, error)
	GetByID(ctx context.Context, request *dto.GetDocumentByIDRequest) (*model.Document, error)
	GetByInfoCardID(ctx context.Context, request *dto.GetDocumentByInfoCardIDRequest) (*model.Document, error)
	Delete(ctx context.Context, request *dto.DeleteDocumentRequest) error
}
