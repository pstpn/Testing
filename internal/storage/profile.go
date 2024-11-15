package storage

import (
	"context"

	"bee/internal/model"
	"bee/internal/service/dto"
)

type ProfileStorage interface {
	Create(ctx context.Context, request *dto.CreateProfileRequest) error
	Get(ctx context.Context, request *dto.GetProfileRequest) (*model.Profile, error)
	Update(ctx context.Context, request *dto.UpdateProfileRequest) error
	Delete(ctx context.Context, request *dto.DeleteProfileRequest) error
}
