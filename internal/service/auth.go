package service

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/jwt"
	"course/pkg/logger"
)

type AuthService interface {
	RegisterEmployee(ctx context.Context, request *dto.RegisterEmployeeRequest) (*tokens, error)
	LoginEmployee(ctx context.Context, request *dto.LoginEmployeeRequest) (*tokens, error)
	VerifyEmployeeAccessToken(ctx context.Context, request *dto.VerifyEmployeeAccessTokenRequest) (*model.Payload, error)
	RefreshTokens(ctx context.Context, request *dto.RefreshEmployeeTokensRequest) (*tokens, error)
}

type authServiceImpl struct {
	logger          logger.Interface
	employeeStorage storage.EmployeeStorage
	infoCardStorage storage.InfoCardStorage

	tokenManager    jwt.TokenManager
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewAuthService(
	logger logger.Interface,
	employeeStorage storage.EmployeeStorage,
	infoCardStorage storage.InfoCardStorage,
	tokenManager jwt.TokenManager,
	accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration,
) AuthService {
	return &authServiceImpl{
		logger:          logger,
		employeeStorage: employeeStorage,
		infoCardStorage: infoCardStorage,

		tokenManager:    tokenManager,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

type tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	IsAdmin      bool
}

func (a *authServiceImpl) RegisterEmployee(ctx context.Context, request *dto.RegisterEmployeeRequest) (*tokens, error) {
	a.logger.Infof("register employee with phone %s", request.PhoneNumber)

	refreshToken, err := a.tokenManager.NewRefreshToken()
	if err != nil {
		a.logger.Errorf("create refresh token for employee: %s", err.Error())
		return nil, fmt.Errorf("create refresh token for employee: %w", err)
	}
	refreshTokenExpiredAt := time.Now().UTC().Add(a.refreshTokenTTL)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		a.logger.Errorf("hash password: %s", err.Error())
		return nil, fmt.Errorf("hash password: %w", err)
	}

	employee, err := a.employeeStorage.Register(ctx, &dto.RegisterEmployeeRequest{
		PhoneNumber:    request.PhoneNumber,
		FullName:       request.FullName,
		CompanyID:      request.CompanyID,
		Post:           request.Post,
		Password:       string(hashedPassword),
		RefreshToken:   refreshToken,
		TokenExpiredAt: &refreshTokenExpiredAt,
		DateOfBirth:    request.DateOfBirth,
	})
	if err != nil {
		a.logger.Errorf("create employee: %s", err.Error())
		return nil, fmt.Errorf("create employee: %w", err)
	}

	now := time.Now()
	infoCard, err := a.infoCardStorage.Create(ctx, &dto.CreateInfoCardRequest{
		EmployeeID:  employee.ID.Int(),
		IsConfirmed: false,
		CreatedDate: &now,
	})
	if err != nil {
		a.logger.Errorf("create employee info card: %s", err.Error())
		return nil, fmt.Errorf("create employee info card: %w", err)
	}

	accessToken, err := a.tokenManager.NewJWT(&model.Payload{InfoCardID: infoCard.ID.String()}, time.Now())
	if err != nil {
		a.logger.Errorf("create access token for employee: %s", err.Error())
		return nil, fmt.Errorf("create access token for employee: %w", err)
	}

	return &tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		IsAdmin:      employee.Post.IsAdmin(),
	}, nil
}

func (a *authServiceImpl) LoginEmployee(ctx context.Context, request *dto.LoginEmployeeRequest) (*tokens, error) {
	a.logger.Infof("login employee with phone %s", request.PhoneNumber)

	employee, err := a.employeeStorage.GetByPhone(ctx, &dto.GetEmployeeRequest{PhoneNumber: request.PhoneNumber})
	if err != nil {
		a.logger.Errorf("get employee: %s", err.Error())
		return nil, fmt.Errorf("get employee: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(request.Password))
	if err != nil {
		return nil, fmt.Errorf("wrong password: %w", err)
	}

	infoCard, err := a.infoCardStorage.GetByEmployeeID(ctx, &dto.GetInfoCardByEmployeeIDRequest{EmployeeID: employee.ID.Int()})
	if err != nil {
		a.logger.Errorf("get info card by employee: %s", err.Error())
		return nil, fmt.Errorf("get info card by employee: %w", err)
	}

	refreshToken, err := a.tokenManager.NewRefreshToken()
	if err != nil {
		a.logger.Errorf("create refresh token for employee: %s", err.Error())
		return nil, fmt.Errorf("create refresh token for employee: %w", err)
	}
	refreshTokenExpiredAt := time.Now().UTC().Add(a.refreshTokenTTL)

	err = a.employeeStorage.UpdateRefreshToken(ctx, &dto.UpdateToken{
		EmployeeID:     employee.ID.Int(),
		RefreshToken:   refreshToken,
		TokenExpiredAt: &refreshTokenExpiredAt,
	})
	if err != nil {
		a.logger.Errorf("update employee refresh token: %s", err.Error())
		return nil, fmt.Errorf("update employee refresh token: %w", err)
	}

	accessToken, err := a.tokenManager.NewJWT(&model.Payload{InfoCardID: infoCard.ID.String()}, time.Now())
	if err != nil {
		a.logger.Errorf("update access token for employee: %s", err.Error())
		return nil, fmt.Errorf("update access token for employee: %w", err)
	}

	return &tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		IsAdmin:      employee.Post.IsAdmin(),
	}, nil
}

func (a *authServiceImpl) VerifyEmployeeAccessToken(ctx context.Context, request *dto.VerifyEmployeeAccessTokenRequest) (*model.Payload, error) {
	a.logger.Infof("verify employee access token")

	payload := &model.Payload{}
	err := a.tokenManager.Parse(request.AccessToken, a.accessTokenTTL, payload)
	if err != nil {
		// Unwrap, because we need to check `github.com/golang-jwt/jwt/v5` library errors
		// using `errors.Is()` method
		return payload, err
	}

	return payload, nil
}

func (a *authServiceImpl) RefreshTokens(ctx context.Context, request *dto.RefreshEmployeeTokensRequest) (*tokens, error) {
	a.logger.Infof("refresh tokens for employee with %d infoCard id", request.InfoCardID)

	employee, err := a.employeeStorage.GetByInfoCardID(ctx, &dto.GetEmployeeByInfoCardIDRequest{InfoCardID: request.InfoCardID})
	if err != nil {
		a.logger.Errorf("get employee: %s", err.Error())
		return nil, fmt.Errorf("get employee: %w", err)
	}
	if employee.RefreshToken != request.RefreshToken {
		a.logger.Errorf("invalid refresh token")
		return nil, fmt.Errorf("invalid refresh token")
	}
	if err = a.tokenManager.RefreshTokenExpired(employee.TokenExpiredAt); err != nil {
		a.logger.Warnf("refresh token expired")
		// Unwrap, because we need to check `github.com/golang-jwt/jwt/v5` library errors
		// using `errors.Is()` method
		return nil, err
	}

	refreshToken, err := a.tokenManager.NewRefreshToken()
	if err != nil {
		a.logger.Errorf("create refresh token for employee: %s", err.Error())
		return nil, fmt.Errorf("create refresh token for employee: %w", err)
	}
	tokenExpiredAt := time.Now().UTC().Add(a.refreshTokenTTL)

	err = a.employeeStorage.UpdateRefreshToken(ctx, &dto.UpdateToken{
		EmployeeID:     employee.ID.Int(),
		RefreshToken:   refreshToken,
		TokenExpiredAt: &tokenExpiredAt,
	})
	if err != nil {
		a.logger.Errorf("update employee refresh token: %s", err.Error())
		return nil, fmt.Errorf("update employee refresh token: %w", err)
	}

	accessToken, err := a.tokenManager.NewJWT(&model.Payload{InfoCardID: strconv.Itoa(int(request.InfoCardID))}, time.Now())
	if err != nil {
		a.logger.Errorf("update access token for employee: %s", err.Error())
		return nil, fmt.Errorf("update access token for employee: %w", err)
	}

	return &tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		IsAdmin:      employee.Post.IsAdmin(),
	}, nil
}
