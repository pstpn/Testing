package service

import (
	"context"
	"fmt"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/logger"
)

type EmployeeService interface {
	GetEmployee(ctx context.Context, request *dto.GetEmployeeRequest) (*model.Employee, error)
	DeleteEmployee(ctx context.Context, request *dto.DeleteEmployeeRequest) error
}

type employeeServiceImpl struct {
	logger          logger.Interface
	employeeStorage storage.EmployeeStorage
}

func NewEmployeeService(logger logger.Interface, employeeStorage storage.EmployeeStorage) EmployeeService {
	return &employeeServiceImpl{
		logger:          logger,
		employeeStorage: employeeStorage,
	}
}

func (e *employeeServiceImpl) GetEmployee(ctx context.Context, request *dto.GetEmployeeRequest) (*model.Employee, error) {
	e.logger.Infof("get employee by phone number %s", request.PhoneNumber)

	employee, err := e.employeeStorage.GetByPhone(ctx, request)
	if err != nil {
		e.logger.Errorf("get employee: %s", err.Error())
		return nil, fmt.Errorf("get employee: %w", err)
	}

	return employee, nil
}

func (e *employeeServiceImpl) DeleteEmployee(ctx context.Context, request *dto.DeleteEmployeeRequest) error {
	e.logger.Infof("delete employee by ID %d", request.EmployeeID)

	err := e.employeeStorage.Delete(ctx, request)
	if err != nil {
		e.logger.Errorf("delete employee: %s", err.Error())
		return fmt.Errorf("delete employee: %w", err)
	}

	return nil
}
