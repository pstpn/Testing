package dto

type CreateCompanyRequest struct {
	Name string
	City string
}

type GetCompanyRequest struct {
	CompanyID int64
}

type DeleteCompanyRequest struct {
	CompanyID int64
}
