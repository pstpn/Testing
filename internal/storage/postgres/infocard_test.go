package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"

	"course/internal/model"
	"course/internal/service/dto"
	"course/pkg/storage/postgres"
)

func Test_infoCardStorageImpl_Create(t *testing.T) {
	infoCardStorage := NewInfoCardStorage(testDB)

	tm, _ := time.Parse(time.RFC3339, "2006-01-02T00:00:00Z")

	request := &dto.CreateInfoCardRequest{
		EmployeeID:  ids["employeeID"],
		IsConfirmed: false,
		CreatedDate: &tm,
	}

	infoCard, err := infoCardStorage.Create(context.TODO(), request)
	require.NoError(t, err)
	require.NotEmpty(t, infoCard)
	require.Equal(t, model.ToEmployeeID(request.EmployeeID), infoCard.CreatedEmployeeID)
	require.Equal(t, request.IsConfirmed, infoCard.IsConfirmed)
	require.Equal(t, request.CreatedDate, infoCard.CreatedDate)

	err = infoCardStorage.Delete(context.TODO(), &dto.DeleteInfoCardRequest{InfoCardID: infoCard.ID.Int()})
	require.NoError(t, err)
}

func Test_infoCardStorageImpl_Validate(t *testing.T) {
	infoCardStorage := NewInfoCardStorage(testDB)

	tm, _ := time.Parse(time.RFC3339, "2006-01-02T00:00:00Z")

	infoCard1, err := infoCardStorage.Create(context.TODO(), &dto.CreateInfoCardRequest{
		EmployeeID:  ids["employeeID"],
		IsConfirmed: false,
		CreatedDate: &tm,
	})
	require.NoError(t, err)
	require.NotEmpty(t, infoCard1)
	require.False(t, infoCard1.IsConfirmed)

	request := &dto.ValidateInfoCardRequest{
		InfoCardID:  infoCard1.ID.Int(),
		IsConfirmed: true,
	}

	err = infoCardStorage.Validate(context.TODO(), request)
	require.NoError(t, err)

	infoCard2, err := infoCardStorage.GetByID(context.TODO(), &dto.GetInfoCardByIDRequest{InfoCardID: request.InfoCardID})
	require.NoError(t, err)
	require.NotEmpty(t, infoCard2)
	require.True(t, infoCard2.IsConfirmed)

	err = infoCardStorage.Delete(context.TODO(), &dto.DeleteInfoCardRequest{InfoCardID: infoCard1.ID.Int()})
	require.NoError(t, err)
}

func Test_infoCardStorageImpl_GetByID(t *testing.T) {
	infoCardStorage := NewInfoCardStorage(testDB)

	tm, _ := time.Parse(time.RFC3339, "2006-01-02T00:00:00Z")

	infoCard1, err := infoCardStorage.Create(context.TODO(), &dto.CreateInfoCardRequest{
		EmployeeID:  ids["employeeID"],
		IsConfirmed: false,
		CreatedDate: &tm,
	})
	require.NoError(t, err)
	require.NotEmpty(t, infoCard1)

	infoCard2, err := infoCardStorage.GetByID(context.TODO(), &dto.GetInfoCardByIDRequest{InfoCardID: infoCard1.ID.Int()})
	require.NoError(t, err)
	require.NotEmpty(t, infoCard2)
	require.Equal(t, infoCard1.ID, infoCard2.ID)
	require.Equal(t, infoCard1.CreatedEmployeeID, infoCard2.CreatedEmployeeID)
	require.Equal(t, infoCard1.IsConfirmed, infoCard2.IsConfirmed)
	require.Equal(t, infoCard1.CreatedDate, infoCard2.CreatedDate)

	err = infoCardStorage.Delete(context.TODO(), &dto.DeleteInfoCardRequest{InfoCardID: infoCard1.ID.Int()})
	require.NoError(t, err)
}

func Test_infoCardStorageImpl_List(t *testing.T) {
	infoCardStorage := NewInfoCardStorage(testDB)

	tm, _ := time.Parse(time.RFC3339, "2006-01-02T00:00:00Z")

	var infoCards []*model.InfoCard
	for range 10 {
		infoCard, err := infoCardStorage.Create(context.TODO(), &dto.CreateInfoCardRequest{
			EmployeeID:  ids["employeeID"],
			IsConfirmed: false,
			CreatedDate: &tm,
		})
		require.NoError(t, err)
		require.NotEmpty(t, infoCard)

		infoCards = append(infoCards, infoCard)
	}

	listInfoCards, err := infoCardStorage.List(context.TODO(), &dto.ListInfoCardsRequest{
		Pagination: &postgres.Pagination{
			PageNumber: 2,
			PageSize:   2,
			Filter: postgres.FilterOptions{
				Pattern: "2",
				Column:  fullColName(employeeTable, phoneNumberField),
			},
			Sort: postgres.SortOptions{
				Direction: postgres.DESC,
				Columns:   []string{infoCardTable, idField},
			},
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, listInfoCards)
	require.Equal(t, 2, len(listInfoCards))

	for _, infoCard := range infoCards {
		err = infoCardStorage.Delete(context.TODO(), &dto.DeleteInfoCardRequest{InfoCardID: infoCard.ID.Int()})
		require.NoError(t, err)
	}
}

func Test_infoCardStorageImpl_Delete(t *testing.T) {
	infoCardStorage := NewInfoCardStorage(testDB)

	tm, _ := time.Parse(time.RFC3339, "2006-01-02T00:00:00Z")

	infoCard1, err := infoCardStorage.Create(context.TODO(), &dto.CreateInfoCardRequest{
		EmployeeID:  ids["employeeID"],
		IsConfirmed: false,
		CreatedDate: &tm,
	})
	require.NoError(t, err)
	require.NotEmpty(t, infoCard1)

	err = infoCardStorage.Delete(context.TODO(), &dto.DeleteInfoCardRequest{InfoCardID: infoCard1.ID.Int()})
	require.NoError(t, err)

	infoCard2, err := infoCardStorage.GetByID(context.TODO(), &dto.GetInfoCardByIDRequest{InfoCardID: infoCard1.ID.Int()})
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, infoCard2)

	err = infoCardStorage.Delete(context.TODO(), &dto.DeleteInfoCardRequest{InfoCardID: infoCard1.ID.Int()})
	require.NoError(t, err)
}
