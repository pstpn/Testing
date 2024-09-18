package dto

type GetEmployeeRequest struct {
	PhoneNumber string
}

type GetEmployeeByInfoCardIDRequest struct {
	InfoCardID int64
}

type DeleteEmployeeRequest struct {
	EmployeeID int64
}
