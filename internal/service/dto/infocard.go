package dto

import (
	"time"

	"course/pkg/storage/postgres"
)

type CreateInfoCardRequest struct {
	EmployeeID  int64
	IsConfirmed bool
	CreatedDate *time.Time
}

type ValidateInfoCardRequest struct {
	InfoCardID  int64
	IsConfirmed bool
}

type GetInfoCardByIDRequest struct {
	InfoCardID int64
}

type GetInfoCardByEmployeeIDRequest struct {
	EmployeeID int64
}

type ListInfoCardsRequest struct {
	Pagination *postgres.Pagination
}

type DeleteInfoCardRequest struct {
	InfoCardID int64
}
