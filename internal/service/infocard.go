package service

import (
	"context"
	"fmt"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/logger"
)

type InfoCardService interface {
	CreateInfoCard(ctx context.Context, request *dto.CreateInfoCardRequest) (*model.InfoCard, error)
	ValidateInfoCard(ctx context.Context, request *dto.ValidateInfoCardRequest) error
	GetInfoCard(ctx context.Context, request *dto.GetInfoCardByIDRequest) (*model.InfoCard, error)
	ListInfoCards(ctx context.Context, request *dto.ListInfoCardsRequest) ([]*model.FullInfoCard, error)
	DeleteInfoCard(ctx context.Context, request *dto.DeleteInfoCardRequest) error
}

type infoCardServiceImpl struct {
	logger          logger.Interface
	infoCardStorage storage.InfoCardStorage
}

func NewInfoCardService(logger logger.Interface, infoCardStorage storage.InfoCardStorage) InfoCardService {
	return &infoCardServiceImpl{
		logger:          logger,
		infoCardStorage: infoCardStorage,
	}
}

func (i *infoCardServiceImpl) CreateInfoCard(ctx context.Context, request *dto.CreateInfoCardRequest) (*model.InfoCard, error) {
	i.logger.Infof("create info card for employee with %d id", request.EmployeeID)

	infoCard, err := i.infoCardStorage.Create(ctx, request)
	if err != nil {
		i.logger.Errorf("create info card: %s", err.Error())
		return nil, fmt.Errorf("create info card: %w", err)
	}

	return infoCard, nil
}

func (i *infoCardServiceImpl) ValidateInfoCard(ctx context.Context, request *dto.ValidateInfoCardRequest) error {
	i.logger.Infof("validate info card by ID %d", request.InfoCardID)

	err := i.infoCardStorage.Validate(ctx, request)
	if err != nil {
		i.logger.Errorf("validate info card: %s", err.Error())
		return fmt.Errorf("validate info card: %w", err)
	}

	return nil
}

func (i *infoCardServiceImpl) GetInfoCard(ctx context.Context, request *dto.GetInfoCardByIDRequest) (*model.InfoCard, error) {
	i.logger.Infof("get info card by ID %d", request.InfoCardID)

	infoCard, err := i.infoCardStorage.GetByID(ctx, request)
	if err != nil {
		i.logger.Errorf("get info card: %s", err.Error())
		return nil, fmt.Errorf("get info card: %w", err)
	}

	return infoCard, nil
}

func (i *infoCardServiceImpl) ListInfoCards(ctx context.Context, request *dto.ListInfoCardsRequest) ([]*model.FullInfoCard, error) {
	i.logger.Infof("list full info cards with pagination")

	fullInfoCards, err := i.infoCardStorage.List(ctx, request)
	if err != nil {
		i.logger.Errorf("list full info cards: %s", err.Error())
		return nil, fmt.Errorf("list full info cards: %w", err)
	}

	return fullInfoCards, nil
}

func (i *infoCardServiceImpl) DeleteInfoCard(ctx context.Context, request *dto.DeleteInfoCardRequest) error {
	i.logger.Infof("delete info card by ID %d", request.InfoCardID)

	err := i.infoCardStorage.Delete(ctx, request)
	if err != nil {
		i.logger.Errorf("delete info card: %s", err.Error())
		return fmt.Errorf("delete info card: %w", err)
	}

	return nil
}
