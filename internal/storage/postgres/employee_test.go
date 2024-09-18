package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"

	"course/internal/model"
	"course/internal/service/dto"
)

func Test_employeeStorageImpl_Register(t *testing.T) {
	employeeStorage := NewEmployeeStorage(testDB)

	tm, _ := time.Parse(time.RFC3339, "2006-01-02T00:00:00Z")

	request := &dto.RegisterEmployeeRequest{
		PhoneNumber: "333",
		FullName:    "123",
		CompanyID:   ids["companyID"],
		Post:        0,
		Password: &model.Password{
			Value:    "432",
			IsHashed: true,
		},
		DateOfBirth: &tm,
	}

	employee, err := employeeStorage.Register(context.TODO(), request)
	require.NoError(t, err)
	require.NotEmpty(t, employee)
	require.Equal(t, request.PhoneNumber, employee.PhoneNumber)
	require.Equal(t, request.FullName, employee.FullName)
	require.Equal(t, model.ToCompanyID(request.CompanyID), employee.CompanyID)
	require.Equal(t, model.ToPostTypeFromInt(request.Post), employee.Post)
	require.Equal(t, request.Password.IsHashed, employee.Password.IsHashed)
	require.Equal(t, request.Password.Value, employee.Password.Value)
	require.Equal(t, request.DateOfBirth, employee.DateOfBirth)

	err = employeeStorage.Delete(context.TODO(), &dto.DeleteEmployeeRequest{EmployeeID: employee.ID.Int()})
	require.NoError(t, err)
}

func Test_employeeStorageImpl_GetByID(t *testing.T) {
	employeeStorage := NewEmployeeStorage(testDB)

	tm, _ := time.Parse(time.RFC3339, "2006-01-02T00:00:00Z")

	employee1, err := employeeStorage.Register(context.TODO(), &dto.RegisterEmployeeRequest{
		PhoneNumber: "321",
		FullName:    "123",
		CompanyID:   ids["companyID"],
		Post:        0,
		Password: &model.Password{
			Value:    "432",
			IsHashed: true,
		},
		DateOfBirth: &tm,
	})
	require.NoError(t, err)
	require.NotEmpty(t, employee1)

	employee2, err := employeeStorage.GetByPhone(context.TODO(), &dto.GetEmployeeRequest{
		PhoneNumber: employee1.PhoneNumber,
	})
	require.NoError(t, err)
	require.NotEmpty(t, employee2)
	require.Equal(t, employee1.ID, employee2.ID)
	require.Equal(t, employee1.FullName, employee2.FullName)
	require.Equal(t, employee1.CompanyID, employee2.CompanyID)
	require.Equal(t, employee1.Post, employee2.Post)
	require.Equal(t, employee1.Password.Value, employee2.Password.Value)
	require.Equal(t, employee1.Password.IsHashed, employee2.Password.IsHashed)
	require.Equal(t, employee1.DateOfBirth, employee2.DateOfBirth)

	err = employeeStorage.Delete(context.TODO(), &dto.DeleteEmployeeRequest{EmployeeID: employee1.ID.Int()})
	require.NoError(t, err)
}

func Test_employeeStorageImpl_Delete(t *testing.T) {
	employeeStorage := NewEmployeeStorage(testDB)

	tm, _ := time.Parse(time.RFC3339, "2006-01-02T00:00:00Z")

	employee1, err := employeeStorage.Register(context.TODO(), &dto.RegisterEmployeeRequest{
		PhoneNumber: "444",
		FullName:    "123",
		CompanyID:   ids["companyID"],
		Post:        0,
		Password: &model.Password{
			Value:    "432",
			IsHashed: true,
		},
		DateOfBirth: &tm,
	})
	require.NoError(t, err)
	require.NotEmpty(t, employee1)

	err = employeeStorage.Delete(context.TODO(), &dto.DeleteEmployeeRequest{EmployeeID: employee1.ID.Int()})
	require.NoError(t, err)

	employee2, err := employeeStorage.GetByPhone(context.TODO(), &dto.GetEmployeeRequest{
		PhoneNumber: employee1.PhoneNumber,
	})
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, employee2)

	err = employeeStorage.Delete(context.TODO(), &dto.DeleteEmployeeRequest{EmployeeID: employee1.ID.Int()})
	require.NoError(t, err)
}
