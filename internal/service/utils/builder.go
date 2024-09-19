package utils

import (
	"time"

	"course/internal/service/dto"
)

type CheckpointBuilder struct {
	checkpointID int64
	documentID   int64
	passageType  int64
	time         *time.Time
}

func (c CheckpointBuilder) WithCheckpointID(checkpointID int64) CheckpointBuilder {
	c.checkpointID = checkpointID
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

func (c CheckpointBuilder) WithTime(time time.Time) CheckpointBuilder {
	c.time = &time
	return c
}

func (c CheckpointBuilder) ToCreateDTO() *dto.CreatePassageRequest {
	return &dto.CreatePassageRequest{
		CheckpointID: c.checkpointID,
		DocumentID:   c.documentID,
		Type:         c.passageType,
		Time:         c.time,
	}
}

func (c CheckpointBuilder) ToListDTO() *dto.ListPassagesRequest {
	return &dto.ListPassagesRequest{
		DocumentID: c.documentID,
	}
}

type CompanyBuilder struct {
	companyID int64
}

func (c CompanyBuilder) WithCompanyID(companyID int64) CompanyBuilder {
	c.companyID = companyID
	return c
}

func (c CompanyBuilder) ToGetDTO() *dto.GetCompanyRequest {
	return &dto.GetCompanyRequest{
		CompanyID: c.companyID,
	}
}
