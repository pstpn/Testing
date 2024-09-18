package storage

import (
	"context"

	"course/internal/model"
	"course/internal/service/dto"
)

//go:generate mockery --name=InfoCardStorage
type InfoCardStorage interface {
	Create(ctx context.Context, request *dto.CreateInfoCardRequest) (*model.InfoCard, error)
	Validate(ctx context.Context, request *dto.ValidateInfoCardRequest) error
	GetByID(ctx context.Context, request *dto.GetInfoCardByIDRequest) (*model.InfoCard, error)
	GetByEmployeeID(ctx context.Context, request *dto.GetInfoCardByEmployeeIDRequest) (*model.InfoCard, error)
	List(ctx context.Context, request *dto.ListInfoCardsRequest) ([]*model.FullInfoCard, error)
	Delete(ctx context.Context, request *dto.DeleteInfoCardRequest) error
}
