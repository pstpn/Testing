package dto

type CreateDocumentRequest struct {
	SerialNumber string
	InfoCardID   int64
	DocumentType int64
}

type GetDocumentByIDRequest struct {
	DocumentID int64
}

type GetDocumentByInfoCardIDRequest struct {
	InfoCardID int64
}

type ListEmployeeDocumentsRequest struct {
	EmployeePhoneNumber string
}

type DeleteDocumentRequest struct {
	DocumentID int64
}
