package inmemory

import (
	"context"
	"fmt"

	"bee/internal/model"
	"bee/internal/service/dto"
	"bee/internal/storage"
	"bee/pkg/storage/inmemory"
)

type profileStorageImpl struct {
	s *inmemory.Storage
}

func NewProfileStorage(s *inmemory.Storage) storage.ProfileStorage {
	return &profileStorageImpl{s: s}
}

func (p *profileStorageImpl) Create(ctx context.Context, request *dto.CreateProfileRequest) error {
	err := p.s.Insert(request.Email, &model.Profile{
		Email:   request.Email,
		Name:    request.Name,
		Surname: request.Surname,
		City:    request.City,
	})
	if err != nil {
		return err
	}

	return nil
}

func (p *profileStorageImpl) Get(ctx context.Context, request *dto.GetProfileRequest) (*model.Profile, error) {
	profile, exist := p.s.Get(request.Email)
	if !exist {
		return nil, fmt.Errorf("profile not found")
	}

	switch pr := profile.(type) {
	case *model.Profile:
		return pr, nil
	default:
		return nil, fmt.Errorf("incorrect profile data")
	}
}

func (p *profileStorageImpl) Update(ctx context.Context, request *dto.UpdateProfileRequest) error {
	p.s.Update(request.Email, &model.Profile{
		Email:   request.Email,
		Name:    request.Name,
		Surname: request.Surname,
		City:    request.City,
	})

	return nil
}

func (p *profileStorageImpl) Delete(ctx context.Context, request *dto.DeleteProfileRequest) error {
	p.s.Delete(request.Email)

	return nil
}
