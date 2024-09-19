package service

import (
	"context"
	"fmt"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/logger"
)

type CheckpointService interface {
	CreatePassage(ctx context.Context, request *dto.CreatePassageRequest) (*model.Passage, error)
	ListPassages(ctx context.Context, request *dto.ListPassagesRequest) ([]*model.Passage, error)
}

type checkpointServiceImpl struct {
	logger            logger.Interface
	checkpointStorage storage.CheckpointStorage
}

func NewCheckpointService(logger logger.Interface, checkpointStorage storage.CheckpointStorage) CheckpointService {
	return &checkpointServiceImpl{
		logger:            logger,
		checkpointStorage: checkpointStorage,
	}
}

func (c *checkpointServiceImpl) CreatePassage(ctx context.Context, request *dto.CreatePassageRequest) (*model.Passage, error) {
	c.logger.Infof("create passage through %d checkpoint with %d document ID", request.CheckpointID, request.DocumentID)

	passage, err := c.checkpointStorage.CreatePassage(ctx, request)
	if err != nil {
		c.logger.Errorf("create passage: %s", err.Error())
		return nil, fmt.Errorf("create passage: %w", err)
	}

	return passage, nil
}

func (c *checkpointServiceImpl) ListPassages(ctx context.Context, request *dto.ListPassagesRequest) ([]*model.Passage, error) {
	c.logger.Infof("list passages by %d document ID", request.DocumentID)

	passages, err := c.checkpointStorage.ListPassages(ctx, request)
	if err != nil {
		c.logger.Errorf("list passages: %s", err.Error())
		return nil, fmt.Errorf("list passages: %w", err)
	}

	return passages, nil
}
