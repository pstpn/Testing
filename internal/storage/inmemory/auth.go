package inmemory

import (
	"context"
	"fmt"

	"bee/internal/model"
	"bee/internal/service/dto"
	"bee/internal/storage"
	"bee/pkg/storage/inmemory"
)

type authStorageImpl struct {
	s *inmemory.Storage
}

func NewAuthStorage(s *inmemory.Storage) storage.AuthStorage {
	return &authStorageImpl{s: s}
}

func (a *authStorageImpl) CreateUser(ctx context.Context, request *dto.RegisterUserRequest) error {
	err := a.s.Insert(request.Email, &model.User{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *authStorageImpl) GetUser(ctx context.Context, request *dto.GetUserRequest) (*model.User, error) {
	user, exist := a.s.Get(request.Email)
	if !exist {
		return nil, fmt.Errorf("unknown user")
	}

	switch u := user.(type) {
	case *model.User:
		return u, nil
	default:
		return nil, fmt.Errorf("incorrect user data")
	}
}
