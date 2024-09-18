package dto

type CreateDocumentFieldRequest struct {
	DocumentID int64
	Type       int64
	Value      string
}

type GetDocumentFieldRequest struct {
	DocumentID int64
	FieldType  int64
}

type ListDocumentFieldsRequest struct {
	DocumentID int64
}

type DeleteDocumentFieldRequest struct {
	FieldID int64
}
