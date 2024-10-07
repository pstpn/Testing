package model

type DocumentType int64

const (
	Passport DocumentType = iota
	DrivingLicense
	UnknownDocument
)

func ToDocumentTypeFromInt(document int64) *DocumentType {
	documentType := DocumentType(document)
	return &documentType
}

func ToDocumentTypeFromString(document string) *DocumentType {
	var documentType DocumentType
	switch document {
	case "Паспорт":
		documentType = Passport
	case "Водительские права":
		documentType = DrivingLicense
	default:
		documentType = UnknownDocument
	}

	return &documentType
}

func (d *DocumentType) String() string {
	switch *d {
	case Passport:
		return "Паспорт"
	case DrivingLicense:
		return "Водительские права"
	default:
		return "Неизвестный документ"
	}
}

func (d *DocumentType) Int() int64 {
	return int64(*d)
}

type DocumentID int64

func ToDocumentID(id int64) *DocumentID {
	documentID := DocumentID(id)
	return &documentID
}

func (d *DocumentID) Int() int64 {
	return int64(*d)
}

type Document struct {
	ID           *DocumentID
	SerialNumber string
	InfoCardID   *InfoCardID
	Type         *DocumentType
}

type DataDocument struct {
	DocumentType string `json:"documentType"`
	SerialNumber string `json:"serialNumber"`
}

type FullDocument struct {
	Data   *DataDocument    `json:"data"`
	Fields []*KeyValueField `json:"fields"`
}
