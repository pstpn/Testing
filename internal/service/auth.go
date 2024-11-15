package service

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"bee/internal/service/dto"
	"bee/internal/storage"
	"bee/pkg/logger"
)

type AuthService interface {
	RegisterUser(ctx context.Context, request *dto.RegisterUserRequest) error
	LoginUser(ctx context.Context, request *dto.LoginUserRequest) (bool, error)
}

type authServiceImpl struct {
	logger      logger.Interface
	authStorage storage.AuthStorage
}

func NewAuthService(logger logger.Interface, authStorage storage.AuthStorage) AuthService {
	return &authServiceImpl{
		logger:      logger,
		authStorage: authStorage,
	}
}

func (a *authServiceImpl) RegisterUser(ctx context.Context, request *dto.RegisterUserRequest) error {
	a.logger.Infof("register user with email %s", request.Email)

	encrypted, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		a.logger.Errorf("generate password: %s", err.Error())
		return fmt.Errorf("generate password: %w", err)
	}

	err = a.authStorage.CreateUser(ctx, &dto.RegisterUserRequest{
		Email:    request.Email,
		Password: string(encrypted),
	})
	if err != nil {
		a.logger.Errorf("register user: %s", err.Error())
		return fmt.Errorf("register user: %w", err)
	}

	return nil
}

func (a *authServiceImpl) LoginUser(ctx context.Context, request *dto.LoginUserRequest) (bool, error) {
	a.logger.Infof("login user with email %s", request.Email)

	user, err := a.authStorage.GetUser(ctx, &dto.GetUserRequest{Email: request.Email})
	if err != nil {
		a.logger.Errorf("get user: %s", err.Error())
		return false, fmt.Errorf("get user: %w", err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		a.logger.Errorf("incorrect password for %s: %s", request.Email, err.Error())
		return false, fmt.Errorf("incorrect password: %w", err)
	}

	return true, nil
}
