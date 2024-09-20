package utils

import (
	"time"

	"course/internal/service/dto"
	"course/pkg/storage/postgres"
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

type InfoCardBuilder struct {
	employeeID  int64
	infoCardID  int64
	isConfirmed bool
	createdDate *time.Time
}

func (i InfoCardBuilder) WithEmployeeID(employeeID int64) InfoCardBuilder {
	i.employeeID = employeeID
	return i
}

func (i InfoCardBuilder) WithIsConfirmed(isConfirmed bool) InfoCardBuilder {
	i.isConfirmed = isConfirmed
	return i
}

func (i InfoCardBuilder) WithCreatedDate(createdDate time.Time) InfoCardBuilder {
	i.createdDate = &createdDate
	return i
}

func (i InfoCardBuilder) WithInfoCardID(infoCardID int64) InfoCardBuilder {
	i.infoCardID = infoCardID
	return i
}

func (i InfoCardBuilder) ToCreateDTO() *dto.CreateInfoCardRequest {
	return &dto.CreateInfoCardRequest{
		EmployeeID:  i.employeeID,
		IsConfirmed: i.isConfirmed,
		CreatedDate: i.createdDate,
	}
}

func (i InfoCardBuilder) ToValidateDTO() *dto.ValidateInfoCardRequest {
	return &dto.ValidateInfoCardRequest{
		InfoCardID:  i.infoCardID,
		IsConfirmed: i.isConfirmed,
	}
}

func (i InfoCardBuilder) ToGetByIDDTO() *dto.GetInfoCardByIDRequest {
	return &dto.GetInfoCardByIDRequest{
		InfoCardID: i.infoCardID,
	}
}

func (i InfoCardBuilder) ToGetByEmployeeIDDTO() *dto.GetInfoCardByEmployeeIDRequest {
	return &dto.GetInfoCardByEmployeeIDRequest{
		EmployeeID: i.employeeID,
	}
}

func (i InfoCardBuilder) ToListDTO() *dto.ListInfoCardsRequest {
	return &dto.ListInfoCardsRequest{
		Pagination: &postgres.Pagination{
			PageNumber: -1,
			PageSize:   -1,
			Filter: postgres.FilterOptions{
				Pattern: "",
				Column:  "",
			},
			Sort: postgres.SortOptions{
				Direction: postgres.ASC,
				Columns:   []string{""},
			},
		},
	}
}

func (i InfoCardBuilder) ToIncorrectListDTO() *dto.ListInfoCardsRequest {
	return &dto.ListInfoCardsRequest{
		Pagination: &postgres.Pagination{
			PageNumber: -1,
			PageSize:   -1,
			Filter: postgres.FilterOptions{
				Pattern: "o214uo213",
				Column:  "4093204u0",
			},
			Sort: postgres.SortOptions{
				Direction: postgres.ASC,
				Columns:   []string{"12y3u2i3"},
			},
		},
	}
}

func (i InfoCardBuilder) ToDeleteDTO() *dto.DeleteInfoCardRequest {
	return &dto.DeleteInfoCardRequest{
		InfoCardID: i.infoCardID,
	}
}
