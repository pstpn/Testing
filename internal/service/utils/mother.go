package utils

import (
	"time"

	"course/internal/service/dto"
)

type AuthObjectMother struct {
	CompanyID int64
}

func (o AuthObjectMother) IncorrectCompanyIDRegisterEmployeeRequest() *dto.RegisterEmployeeRequest {
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

func (o AuthObjectMother) DefaultRegisterEmployeeRequest() *dto.RegisterEmployeeRequest {
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

func (o AuthObjectMother) IncorrectPhoneNumberLoginEmployeeRequest() *dto.LoginEmployeeRequest {
	return &dto.LoginEmployeeRequest{
		PhoneNumber: "7n8c8937d38d73",
		Password:    "21e12",
	}
}

func (o AuthObjectMother) IncorrectPasswordLoginEmployeeRequest() *dto.LoginEmployeeRequest {
	return &dto.LoginEmployeeRequest{
		PhoneNumber: "123",
		Password:    "89t3n82rdjy437dr",
	}
}

func (o AuthObjectMother) DefaultLoginEmployeeRequest() *dto.LoginEmployeeRequest {
	return &dto.LoginEmployeeRequest{
		PhoneNumber: "123",
		Password:    "21e12",
	}
}

type DocumentObjectMother struct {
	InfoCardID int64
	DocumentID int64
}

func (d DocumentObjectMother) IncorrectInfoCardIDCreateDocumentRequest() *dto.CreateDocumentRequest {
	return &dto.CreateDocumentRequest{
		SerialNumber: "0001",
		InfoCardID:   -1,
		DocumentType: 1,
	}
}

func (d DocumentObjectMother) DefaultCreateDocumentRequest() *dto.CreateDocumentRequest {
	return &dto.CreateDocumentRequest{
		SerialNumber: "0001",
		InfoCardID:   d.InfoCardID,
		DocumentType: 1,
	}
}

func (d DocumentObjectMother) IncorrectDocumentIDGetDocumentRequest() *dto.GetDocumentByIDRequest {
	return &dto.GetDocumentByIDRequest{
		DocumentID: -1,
	}
}

func (d DocumentObjectMother) DefaultGetDocumentRequest() *dto.GetDocumentByIDRequest {
	return &dto.GetDocumentByIDRequest{
		DocumentID: d.DocumentID,
	}
}

func (d DocumentObjectMother) IncorrectInfoCardIDGetDocumentByInfoCardIDRequest() *dto.GetDocumentByInfoCardIDRequest {
	return &dto.GetDocumentByInfoCardIDRequest{
		InfoCardID: -1,
	}
}

func (d DocumentObjectMother) DefaultGetDocumentByInfoCardIDRequest() *dto.GetDocumentByInfoCardIDRequest {
	return &dto.GetDocumentByInfoCardIDRequest{
		InfoCardID: d.InfoCardID,
	}
}

func (d DocumentObjectMother) IncorrectDocumentIDDeleteDocumentRequest() *dto.DeleteDocumentRequest {
	return &dto.DeleteDocumentRequest{
		DocumentID: -1,
	}
}

func (d DocumentObjectMother) DefaultDeleteDocumentRequest() *dto.DeleteDocumentRequest {
	return &dto.DeleteDocumentRequest{
		DocumentID: d.DocumentID,
	}
}

type EmployeeObjectMother struct {
	EmployeeID int64
}

func (e EmployeeObjectMother) IncorrectPhoneNumberGetEmployeeRequest() *dto.GetEmployeeRequest {
	return &dto.GetEmployeeRequest{
		PhoneNumber: "86^*(8723d",
	}
}

func (e EmployeeObjectMother) DefaultGetEmployeeRequest() *dto.GetEmployeeRequest {
	return &dto.GetEmployeeRequest{
		PhoneNumber: "123",
	}
}

func (e EmployeeObjectMother) IncorrectEmployeeIDDeleteEmployeeRequest() *dto.DeleteEmployeeRequest {
	return &dto.DeleteEmployeeRequest{
		EmployeeID: -123,
	}
}

func (e EmployeeObjectMother) DefaultDeleteEmployeeRequest() *dto.DeleteEmployeeRequest {
	return &dto.DeleteEmployeeRequest{
		EmployeeID: e.EmployeeID,
	}
}

type FieldObjectMother struct {
	DocumentID int64
	FieldID    int64
}

func (f FieldObjectMother) IncorrectDocumentIDCreateDocumentFieldRequest() *dto.CreateDocumentFieldRequest {
	return &dto.CreateDocumentFieldRequest{
		DocumentID: -111,
		Type:       1,
		Value:      "123",
	}
}

func (f FieldObjectMother) DefaultCreateDocumentFieldRequest1() *dto.CreateDocumentFieldRequest {
	return &dto.CreateDocumentFieldRequest{
		DocumentID: f.DocumentID,
		Type:       0,
		Value:      "123",
	}
}

func (f FieldObjectMother) DefaultCreateDocumentFieldRequest2() *dto.CreateDocumentFieldRequest {
	return &dto.CreateDocumentFieldRequest{
		DocumentID: f.DocumentID,
		Type:       1,
		Value:      "123",
	}
}

func (f FieldObjectMother) IncorrectDocumentIDGetDocumentFieldRequest() *dto.GetDocumentFieldRequest {
	return &dto.GetDocumentFieldRequest{
		DocumentID: -1,
		FieldType:  1,
	}
}

func (f FieldObjectMother) DefaultGetDocumentFieldRequest2() *dto.GetDocumentFieldRequest {
	return &dto.GetDocumentFieldRequest{
		DocumentID: f.DocumentID,
		FieldType:  1,
	}
}

func (f FieldObjectMother) IncorrectDocumentIDListDocumentFieldsRequest() *dto.ListDocumentFieldsRequest {
	return &dto.ListDocumentFieldsRequest{
		DocumentID: -1,
	}
}

func (f FieldObjectMother) DefaultListDocumentFieldsRequest() *dto.ListDocumentFieldsRequest {
	return &dto.ListDocumentFieldsRequest{
		DocumentID: f.DocumentID,
	}
}

func (f FieldObjectMother) IncorrectFieldIDDeleteDocumentFieldRequest() *dto.DeleteDocumentFieldRequest {
	return &dto.DeleteDocumentFieldRequest{
		FieldID: -1,
	}
}

func (f FieldObjectMother) DefaultDeleteDocumentFieldRequest() *dto.DeleteDocumentFieldRequest {
	return &dto.DeleteDocumentFieldRequest{
		FieldID: f.FieldID,
	}
}