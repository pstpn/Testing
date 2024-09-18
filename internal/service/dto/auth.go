package dto

import (
	"time"
)

type RegisterEmployeeRequest struct {
	PhoneNumber    string
	FullName       string
	CompanyID      int64
	Post           int64
	Password       string
	RefreshToken   string
	TokenExpiredAt *time.Time
	DateOfBirth    *time.Time
}

type UpdateToken struct {
	EmployeeID     int64
	RefreshToken   string
	TokenExpiredAt *time.Time
}

type LoginEmployeeRequest struct {
	PhoneNumber string
	Password    string
}

type VerifyEmployeeAccessTokenRequest struct {
	AccessToken string
}

type RefreshEmployeeTokensRequest struct {
	InfoCardID   int64
	RefreshToken string
}
