package storage

import (
	"context"

	"course/internal/model"
	"course/internal/service/dto"
)

//go:generate mockery --name=CheckpointStorage
type CheckpointStorage interface {
	CreatePassage(ctx context.Context, request *dto.CreatePassageRequest) (*model.Passage, error)
	GetPassage(ctx context.Context, request *dto.GetPassageRequest) (*model.Passage, error)
	ListPassages(ctx context.Context, request *dto.ListPassagesRequest) ([]*model.Passage, error)
	DeletePassage(ctx context.Context, request *dto.DeletePassageRequest) error

	CreateCheckpoint(ctx context.Context, request *dto.CreateCheckpointRequest) (*model.Checkpoint, error)
	GetCheckpoint(ctx context.Context, request *dto.GetCheckpointRequest) (*model.Checkpoint, error)
	DeleteCheckpoint(ctx context.Context, request *dto.DeleteCheckpointRequest) error
}
