package model

import (
	"strconv"
	"time"
)

type InfoCardID int64

func ToInfoCardID(id int64) *InfoCardID {
	infoCardID := InfoCardID(id)
	return &infoCardID
}

func (i *InfoCardID) Int() int64 {
	return int64(*i)
}

func (i *InfoCardID) String() string {
	return strconv.FormatInt(i.Int(), 10)
}

type InfoCard struct {
	ID                *InfoCardID
	CreatedEmployeeID *EmployeeID
	IsConfirmed       bool
	CreatedDate       *time.Time
}

type FullInfoCard struct {
	ID                *InfoCardID `json:"id"`
	CreatedEmployeeID *EmployeeID `json:"createdEmployeeID"`
	IsConfirmed       bool        `json:"isConfirmed"`
	CreatedDate       *time.Time  `json:"createdDate"`

	FullName    string     `json:"fullName"`
	PhoneNumber string     `json:"phoneNumber"`
	CompanyID   *CompanyID `json:"companyID"`
	Post        string     `json:"post"`
	DateOfBirth *time.Time `json:"dateOfBirth"`
}

type InfoCardWithPassages struct {
	Document
}
