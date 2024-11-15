package service

import (
	"context"
	"fmt"

	"bee/internal/model"
	"bee/internal/service/dto"
	"bee/internal/storage"
	"bee/pkg/logger"
)

type ProfileService interface {
	CreateProfile(ctx context.Context, request *dto.CreateProfileRequest) error
	GetProfile(ctx context.Context, request *dto.GetProfileRequest) (*model.Profile, error)
	UpdateProfile(ctx context.Context, request *dto.UpdateProfileRequest) error
	DeleteProfile(ctx context.Context, request *dto.DeleteProfileRequest) error
}

type profileServiceImpl struct {
	logger         logger.Interface
	profileStorage storage.ProfileStorage
}

func NewProfileService(logger logger.Interface, profileStorage storage.ProfileStorage) ProfileService {
	return &profileServiceImpl{
		logger:         logger,
		profileStorage: profileStorage,
	}
}

func (p *profileServiceImpl) CreateProfile(ctx context.Context, request *dto.CreateProfileRequest) error {
	p.logger.Infof("create profile for %s %s", request.Name, request.Surname)

	err := p.profileStorage.Create(ctx, request)
	if err != nil {
		p.logger.Errorf("create profile: %s", err.Error())
		return fmt.Errorf("create profile: %w", err)
	}

	return nil
}

func (p *profileServiceImpl) GetProfile(ctx context.Context, request *dto.GetProfileRequest) (*model.Profile, error) {
	p.logger.Infof("get profile with %s email", request.Email)

	profile, err := p.profileStorage.Get(ctx, request)
	if err != nil {
		p.logger.Errorf("get profile: %s", err.Error())
		return nil, fmt.Errorf("get profile: %w", err)
	}

	return profile, nil
}

func (p *profileServiceImpl) UpdateProfile(ctx context.Context, request *dto.UpdateProfileRequest) error {
	p.logger.Infof("update %s %s profile", request.Name, request.Surname)

	err := p.profileStorage.Update(ctx, request)
	if err != nil {
		p.logger.Errorf("update profile: %s", err.Error())
		return fmt.Errorf("update profile: %w", err)
	}

	return nil
}

func (p *profileServiceImpl) DeleteProfile(ctx context.Context, request *dto.DeleteProfileRequest) error {
	p.logger.Infof("delete profile with %s email", request.Email)

	err := p.profileStorage.Delete(ctx, request)
	if err != nil {
		p.logger.Errorf("delete profile: %s", err.Error())
		return fmt.Errorf("delete profile: %w", err)
	}

	return nil
}
