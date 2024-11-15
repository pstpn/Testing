package storage

import (
	"context"

	"bee/internal/model"
	"bee/internal/service/dto"
)

type AuthStorage interface {
	CreateUser(ctx context.Context, request *dto.RegisterUserRequest) error
	GetUser(ctx context.Context, request *dto.GetUserRequest) (*model.User, error)
}
