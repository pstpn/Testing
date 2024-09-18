package postgres

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"

	"course/internal/service/dto"
)

func Test_companyStorageImpl_Create(t *testing.T) {
	companyStorage := NewCompanyStorage(testDB)

	request := &dto.CreateCompanyRequest{
		Name: "test",
		City: "tetest",
	}

	company, err := companyStorage.Create(context.TODO(), request)
	require.NoError(t, err)
	require.NotEmpty(t, company)
	require.Equal(t, request.Name, company.Name)
	require.Equal(t, request.City, company.City)

	err = companyStorage.Delete(context.TODO(), &dto.DeleteCompanyRequest{CompanyID: company.ID.Int()})
	require.NoError(t, err)
}

func Test_companyStorageImpl_GetByID(t *testing.T) {
	companyStorage := NewCompanyStorage(testDB)

	company1, err := companyStorage.Create(context.TODO(), &dto.CreateCompanyRequest{
		Name: "test",
		City: "tetest",
	})
	require.NoError(t, err)
	require.NotEmpty(t, company1)

	company2, err := companyStorage.GetByID(context.TODO(), &dto.GetCompanyRequest{CompanyID: company1.ID.Int()})
	require.NoError(t, err)
	require.NotEmpty(t, company2)
	require.Equal(t, company1.ID, company2.ID)
	require.Equal(t, company1.Name, company2.Name)
	require.Equal(t, company1.City, company2.City)

	err = companyStorage.Delete(context.TODO(), &dto.DeleteCompanyRequest{CompanyID: company1.ID.Int()})
	require.NoError(t, err)
}

func Test_companyStorageImpl_Delete(t *testing.T) {
	companyStorage := NewCompanyStorage(testDB)

	company1, err := companyStorage.Create(context.TODO(), &dto.CreateCompanyRequest{
		Name: "test",
		City: "tetest",
	})
	require.NoError(t, err)
	require.NotEmpty(t, company1)

	err = companyStorage.Delete(context.TODO(), &dto.DeleteCompanyRequest{CompanyID: company1.ID.Int()})
	require.NoError(t, err)

	company2, err := companyStorage.GetByID(context.TODO(), &dto.GetCompanyRequest{CompanyID: company1.ID.Int()})
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, company2)

	err = companyStorage.Delete(context.TODO(), &dto.DeleteCompanyRequest{CompanyID: company1.ID.Int()})
	require.NoError(t, err)
}
