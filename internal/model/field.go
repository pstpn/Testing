package model

type FieldID int64

func ToFieldID(id int64) *FieldID {
	fieldID := FieldID(id)
	return &fieldID
}

func (f *FieldID) Int() int64 {
	return int64(*f)
}

type FieldType int64

const (
	DateOfRelease FieldType = iota
	IssuingAuthority
	PlaceOfIssue
	UnknownFieldType
)

func ToFieldTypeFromInt(field int64) *FieldType {
	fieldType := FieldType(field)
	return &fieldType
}

func ToFieldTypeFromString(field string) *FieldType {
	var fieldType FieldType
	switch field {
	case "Дата выдачи":
		fieldType = DateOfRelease
	case "Выдавший орган":
		fieldType = IssuingAuthority
	case "Место выдачи":
		fieldType = PlaceOfIssue
	default:
		fieldType = UnknownFieldType
	}

	return &fieldType
}

func (f *FieldType) String() string {
	switch *f {
	case DateOfRelease:
		return "Дата выдачи"
	case IssuingAuthority:
		return "Выдавший орган"
	case PlaceOfIssue:
		return "Место выдачи"
	default:
		return "Неизвестное поле"
	}
}

func (f *FieldType) Int() int64 {
	return int64(*f)
}

type Field struct {
	ID         *FieldID
	DocumentID *DocumentID
	Type       *FieldType
	Value      string
}

type KeyValueField struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func ModelToKeyValue(documentFields []*Field) []*KeyValueField {
	fields := make([]*KeyValueField, 0)
	for _, documentField := range documentFields {
		fields = append(fields, &KeyValueField{
			Type:  documentField.Type.String(),
			Value: documentField.Value,
		})
	}
	return fields
}
