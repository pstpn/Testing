package storage

import (
	"context"

	"course/internal/model"
	"course/internal/service/dto"
)

type PhotoMetaStorage interface {
	SaveKey(ctx context.Context, request *dto.CreatePhotoKeyRequest) (*model.PhotoMeta, error)
	GetKey(ctx context.Context, request *dto.GetPhotoRequest) (*model.PhotoMeta, error)
	DeleteKey(ctx context.Context, request *dto.DeletePhotoRequest) error
}

type PhotoDataStorage interface {
	Save(ctx context.Context, request *dto.CreatePhotoRequest) (*model.PhotoKey, error)
	Get(ctx context.Context, key *model.PhotoKey) ([]byte, error)
	Delete(ctx context.Context, key *model.PhotoKey) error
}

//go:generate mockery --name=PhotoStorages
type PhotoStorages interface {
	PhotoDataStorage
	PhotoMetaStorage
}

type PhotoStorage struct {
	PhotoDataStorage
	PhotoMetaStorage
}
