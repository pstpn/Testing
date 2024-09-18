package model

type CompanyID int64

func ToCompanyID(id int64) *CompanyID {
	companyID := CompanyID(id)
	return &companyID
}

func (c *CompanyID) Int() int64 {
	return int64(*c)
}

type Company struct {
	ID   *CompanyID
	Name string
	City string
}
