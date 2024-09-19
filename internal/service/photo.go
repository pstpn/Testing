package service

import (
	"context"
	"fmt"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/logger"
)

type PhotoService interface {
	CreatePhoto(ctx context.Context, request *dto.CreatePhotoRequest) (*model.PhotoMeta, error)
	GetPhoto(ctx context.Context, request *dto.GetPhotoRequest) (*model.Photo, error)
	DeletePhoto(ctx context.Context, request *dto.DeletePhotoRequest) error
}

type photoServiceImpl struct {
	logger       logger.Interface
	photoStorage storage.PhotoStorages
}

func NewPhotoService(logger logger.Interface, photoStorage storage.PhotoStorages) PhotoService {
	return &photoServiceImpl{
		logger:       logger,
		photoStorage: photoStorage,
	}
}

func (p *photoServiceImpl) CreatePhoto(ctx context.Context, request *dto.CreatePhotoRequest) (*model.PhotoMeta, error) {
	// FIXME: Crop face from document
	p.logger.Infof("create photo by document ID %d", request.DocumentID)

	key, err := p.photoStorage.Save(ctx, request)
	if err != nil {
		p.logger.Errorf("save photo: %s", err.Error())
		return nil, fmt.Errorf("save photo: %w", err)
	}

	photoMeta, err := p.photoStorage.SaveKey(ctx, &dto.CreatePhotoKeyRequest{
		DocumentID: model.ToDocumentID(request.DocumentID),
		Key:        key,
	})
	if err != nil {
		p.logger.Errorf("save photo key: %s", err.Error())
		return nil, fmt.Errorf("save photo key: %w", err)
	}

	return photoMeta, nil
}

func (p *photoServiceImpl) GetPhoto(ctx context.Context, request *dto.GetPhotoRequest) (*model.Photo, error) {
	p.logger.Infof("get photo by document ID %d", request.DocumentID)

	meta, err := p.photoStorage.GetKey(ctx, request)
	if err != nil {
		p.logger.Errorf("get photo key: %s", err.Error())
		return nil, fmt.Errorf("get photo key: %w", err)
	}

	data, err := p.photoStorage.Get(ctx, meta.PhotoKey)
	if err != nil {
		p.logger.Errorf("get photo by key: %s", err.Error())
		return nil, fmt.Errorf("get photo by key: %w", err)
	}

	return &model.Photo{
		Meta: meta,
		Data: data,
	}, nil
}

func (p *photoServiceImpl) DeletePhoto(ctx context.Context, request *dto.DeletePhotoRequest) error {
	p.logger.Infof("delete photo by document ID %d", request.DocumentID)

	meta, err := p.photoStorage.GetKey(ctx, &dto.GetPhotoRequest{DocumentID: request.DocumentID})
	if err != nil {
		p.logger.Errorf("get photo key: %s", err.Error())
		return fmt.Errorf("get photo key: %w", err)
	}

	err = p.photoStorage.Delete(ctx, meta.PhotoKey)
	if err != nil {
		p.logger.Errorf("delete photo by key: %s", err.Error())
		return fmt.Errorf("delete photo by key: %w", err)
	}

	err = p.photoStorage.DeleteKey(ctx, request)
	if err != nil {
		p.logger.Errorf("delete photo key: %s", err.Error())
		return fmt.Errorf("delete photo key: %w", err)
	}

	return nil
}
