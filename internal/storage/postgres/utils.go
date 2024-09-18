package postgres

import (
	"fmt"
	"strings"
)

const (
	checkpointTable = "checkpoint"
	companyTable    = "company"
	documentTable   = "document"
	employeeTable   = "employee"
	fieldTable      = "field"
	infoCardTable   = "info_card"
	photoTable      = "photo"
	passageTable    = "passage"
)

const (
	idField                = "id"
	phoneNumberField       = "phone_number"
	nameField              = "name"
	cityField              = "city"
	serialNumberField      = "serial_number"
	infoCardIdField        = "info_card_id"
	fullNameField          = "full_name"
	companyIdField         = "company_id"
	postField              = "post"
	passwordField          = "password"
	refreshTokenField      = "refresh_token"
	tokenExpiredAtField    = "token_expired_at"
	dateOfBirthField       = "date_of_birth"
	documentIdField        = "document_id"
	typeField              = "type"
	valueField             = "value"
	createdEmployeeIDField = "created_employee_id"
	isConfirmedField       = "is_confirmed"
	createdDateField       = "created_date"
	checkpointIdField      = "checkpoint_id"
	timeField              = "time"
	keyField               = "key"
)

func fullColName(tableName, columnName string) string {
	return fmt.Sprintf("%s.%s", tableName, columnName)
}

func on(baseTable, targetTable, baseColumn, targetColumn string) string {
	return fmt.Sprintf("%s on %s.%s=%s.%s",
		targetTable,
		baseTable,
		baseColumn,
		targetTable,
		targetColumn,
	)
}

func returningCompanyColumns() string {
	return fmt.Sprintf("RETURNING %s", strings.Join([]string{
		idField,
		nameField,
		cityField,
	}, ","))
}

func returningDocumentColumns() string {
	return fmt.Sprintf("RETURNING %s", strings.Join([]string{
		idField,
		serialNumberField,
		infoCardIdField,
		typeField,
	}, ","))
}

func returningEmployeeColumns() string {
	return fmt.Sprintf("RETURNING %s", strings.Join([]string{
		idField,
		phoneNumberField,
		fullNameField,
		companyIdField,
		postField,
		passwordField,
		refreshTokenField,
		tokenExpiredAtField,
		dateOfBirthField,
	}, ","))
}

func returningFieldColumns() string {
	return fmt.Sprintf("RETURNING %s", strings.Join([]string{
		idField,
		documentIdField,
		typeField,
		valueField,
	}, ","))
}

func returningInfoCardColumns() string {
	return fmt.Sprintf("RETURNING %s", strings.Join([]string{
		idField,
		createdEmployeeIDField,
		isConfirmedField,
		createdDateField,
	}, ","))
}

func returningPassageColumns() string {
	return fmt.Sprintf("RETURNING %s", strings.Join([]string{
		idField,
		checkpointIdField,
		documentIdField,
		typeField,
		timeField,
	}, ","))
}

func returningCheckpointColumns() string {
	return fmt.Sprintf("RETURNING %s", strings.Join([]string{
		idField,
		phoneNumberField,
	}, ","))
}

func returningPhotoMetaColumns() string {
	return fmt.Sprintf("RETURNING %s", strings.Join([]string{
		idField,
		documentIdField,
		keyField,
	}, ","))
}
