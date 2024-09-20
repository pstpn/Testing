package utils

import (
	"time"

	"course/internal/service/dto"
)

type CheckpointBuilder struct {
	checkpointID int64
	phoneNumber  string
	documentID   int64
	passageType  int64
	passageID    int64
	time         *time.Time
}

func (c CheckpointBuilder) WithCheckpointID(checkpointID int64) CheckpointBuilder {
	c.checkpointID = checkpointID
	return c
}

func (c CheckpointBuilder) WithPhoneNumber(phoneNumber string) CheckpointBuilder {
	c.phoneNumber = phoneNumber
	return c
}

func (c CheckpointBuilder) WithDocumentID(documentID int64) CheckpointBuilder {
	c.documentID = documentID
	return c
}

func (c CheckpointBuilder) WithPassageType(passageType int64) CheckpointBuilder {
	c.passageType = passageType
	return c
}

func (c CheckpointBuilder) WithPassageID(passageID int64) CheckpointBuilder {
	c.passageID = passageID
	return c
}

func (c CheckpointBuilder) WithTime(time time.Time) CheckpointBuilder {
	c.time = &time
	return c
}

func (c CheckpointBuilder) ToCreateCheckpointDTO() *dto.CreateCheckpointRequest {
	return &dto.CreateCheckpointRequest{
		PhoneNumber: c.phoneNumber,
	}
}

func (c CheckpointBuilder) ToCreateDTO() *dto.CreatePassageRequest {
	return &dto.CreatePassageRequest{
		CheckpointID: c.checkpointID,
		DocumentID:   c.documentID,
		Type:         c.passageType,
		Time:         c.time,
	}
}

func (c CheckpointBuilder) ToGetCheckpointDTO() *dto.GetCheckpointRequest {
	return &dto.GetCheckpointRequest{
		CheckpointID: c.checkpointID,
	}
}

func (c CheckpointBuilder) ToGetDTO() *dto.GetPassageRequest {
	return &dto.GetPassageRequest{
		PassageID: c.passageID,
	}
}

func (c CheckpointBuilder) ToListDTO() *dto.ListPassagesRequest {
	return &dto.ListPassagesRequest{
		DocumentID: c.documentID,
	}
}

func (c CheckpointBuilder) ToDeleteCheckpointDTO() *dto.DeleteCheckpointRequest {
	return &dto.DeleteCheckpointRequest{
		CheckpointID: c.checkpointID,
	}
}

func (c CheckpointBuilder) ToDeleteDTO() *dto.DeletePassageRequest {
	return &dto.DeletePassageRequest{
		PassageID: c.passageID,
	}
}

type CompanyBuilder struct {
	companyID int64
	name      string
	city      string
}

func (c CompanyBuilder) WithCompanyID(companyID int64) CompanyBuilder {
	c.companyID = companyID
	return c
}

func (c CompanyBuilder) WithName(name string) CompanyBuilder {
	c.name = name
	return c
}

func (c CompanyBuilder) WithCity(city string) CompanyBuilder {
	c.city = city
	return c
}

func (c CompanyBuilder) ToGetDTO() *dto.GetCompanyRequest {
	return &dto.GetCompanyRequest{
		CompanyID: c.companyID,
	}
}

func (c CompanyBuilder) ToCreateDTO() *dto.CreateCompanyRequest {
	return &dto.CreateCompanyRequest{
		Name: c.name,
		City: c.city,
	}
}

func (c CompanyBuilder) ToDeleteDTO() *dto.DeleteCompanyRequest {
	return &dto.DeleteCompanyRequest{
		CompanyID: c.companyID,
	}
}
