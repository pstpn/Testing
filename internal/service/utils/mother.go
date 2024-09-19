package utils

import (
	"time"

	"course/internal/service/dto"
)

type ObjectMother struct {
	CompanyID int64
}

func (o ObjectMother) IncorrectCompanyIDRegisterEmployeeRequest() *dto.RegisterEmployeeRequest {
	expiredAt := time.Now().Add(1500 * time.Hour)
	dateOfBirth := time.Now().Add(-1500 * time.Hour)
	return &dto.RegisterEmployeeRequest{
		PhoneNumber:    "+55555555555",
		FullName:       "Stepa Stepan",
		CompanyID:      -1,
		Post:           1,
		Password:       "123",
		RefreshToken:   "456789",
		TokenExpiredAt: &expiredAt,
		DateOfBirth:    &dateOfBirth,
	}
}

func (o ObjectMother) DefaultRegisterEmployeeRequest() *dto.RegisterEmployeeRequest {
	expiredAt := time.Now().Add(1500 * time.Hour)
	dateOfBirth := time.Now().Add(-1500 * time.Hour)
	return &dto.RegisterEmployeeRequest{
		PhoneNumber:    "+55555555555",
		FullName:       "Stepa Stepan",
		CompanyID:      o.CompanyID,
		Post:           1,
		Password:       "123",
		RefreshToken:   "456789",
		TokenExpiredAt: &expiredAt,
		DateOfBirth:    &dateOfBirth,
	}
}

func (o ObjectMother) IncorrectPhoneNumberLoginEmployeeRequest() *dto.LoginEmployeeRequest {
	return &dto.LoginEmployeeRequest{
		PhoneNumber: "7n8c8937d38d73",
		Password:    "21e12",
	}
}

func (o ObjectMother) IncorrectPasswordLoginEmployeeRequest() *dto.LoginEmployeeRequest {
	return &dto.LoginEmployeeRequest{
		PhoneNumber: "123",
		Password:    "89t3n82rdjy437dr",
	}
}

func (o ObjectMother) DefaultLoginEmployeeRequest() *dto.LoginEmployeeRequest {
	return &dto.LoginEmployeeRequest{
		PhoneNumber: "123",
		Password:    "21e12",
	}
}
