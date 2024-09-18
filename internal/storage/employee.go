package storage

import (
	"context"

	"course/internal/model"
	"course/internal/service/dto"
)

//go:generate mockery --name=EmployeeStorage
type EmployeeStorage interface {
	Register(ctx context.Context, request *dto.RegisterEmployeeRequest) (*model.Employee, error)
	UpdateRefreshToken(ctx context.Context, request *dto.UpdateToken) error
	GetByPhone(ctx context.Context, request *dto.GetEmployeeRequest) (*model.Employee, error)
	GetByInfoCardID(ctx context.Context, request *dto.GetEmployeeByInfoCardIDRequest) (*model.Employee, error)
	Delete(ctx context.Context, request *dto.DeleteEmployeeRequest) error
}
