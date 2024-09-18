package postgres

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/storage/postgres"
)

const connURL = "postgresql://postgres:admin@localhost:5432/tests"

var testDB *postgres.Postgres
var ids map[string]int64

func TestMain(m *testing.M) {
	testDB = NewTestStorage()

	code := m.Run()
	DropTestStorage(testDB, ids)
	testDB.Close()

	os.Exit(code)
}

func NewTestStorage() *postgres.Postgres {
	conn, err := postgres.New(connURL)
	if err != nil {
		panic(err)
	}

	ids = map[string]int64{}
	ids["companyID"] = initTestCompanyStorage(NewCompanyStorage(conn))
	ids["employeeID"] = initTestEmployeeStorage(NewEmployeeStorage(conn))
	ids["infoCardID"] = initTestInfoCardStorage(NewInfoCardStorage(conn))
	ids["documentID"] = initTestDocumentStorage(NewDocumentStorage(conn))
	ids["checkpointID"] = initTestCheckpointStorage(NewCheckpointStorage(conn))
	ids["photoID"] = initTestPhotoMetaStorage(NewPhotoMetaStorage(conn))

	return conn
}

func DropTestStorage(testDB *postgres.Postgres, ids map[string]int64) {
	err := NewPhotoMetaStorage(testDB).DeleteKey(context.TODO(), &dto.DeletePhotoRequest{DocumentID: ids["photoID"]})
	if err != nil {
		panic(err)
	}
	err = NewCheckpointStorage(testDB).DeleteCheckpoint(context.TODO(), &dto.DeleteCheckpointRequest{CheckpointID: ids["checkpointID"]})
	if err != nil {
		panic(err)
	}
	err = NewDocumentStorage(testDB).Delete(context.TODO(), &dto.DeleteDocumentRequest{DocumentID: ids["documentID"]})
	if err != nil {
		panic(err)
	}
	err = NewInfoCardStorage(testDB).Delete(context.TODO(), &dto.DeleteInfoCardRequest{InfoCardID: ids["infoCardID"]})
	if err != nil {
		panic(err)
	}
	err = NewEmployeeStorage(testDB).Delete(context.TODO(), &dto.DeleteEmployeeRequest{EmployeeID: ids["employeeID"]})
	if err != nil {
		panic(err)
	}
	err = NewCompanyStorage(testDB).Delete(context.TODO(), &dto.DeleteCompanyRequest{CompanyID: ids["companyID"]})
	if err != nil {
		panic(err)
	}
}

func initTestCompanyStorage(storage storage.CompanyStorage) int64 {
	company, err := storage.Create(context.TODO(), &dto.CreateCompanyRequest{
		Name: "Test",
		City: "Test",
	})
	if err != nil && !strings.Contains(err.Error(), "constraint") {
		panic(err)
	}

	return company.ID.Int()
}

func initTestEmployeeStorage(storage storage.EmployeeStorage) int64 {
	tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")

	employee, err := storage.Register(context.TODO(), &dto.RegisterEmployeeRequest{
		PhoneNumber: "123",
		FullName:    "123",
		CompanyID:   ids["companyID"],
		Post:        1,
		Password: &model.Password{
			Value:    "123",
			IsHashed: true,
		},
		DateOfBirth: &tm,
	})
	if err != nil {
		panic(err)
	}

	return employee.ID.Int()
}

func initTestInfoCardStorage(storage storage.InfoCardStorage) int64 {
	tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")

	infoCard, err := storage.Create(context.TODO(), &dto.CreateInfoCardRequest{
		EmployeeID:  ids["employeeID"],
		IsConfirmed: true,
		CreatedDate: &tm,
	})
	if err != nil && !strings.Contains(err.Error(), "constraint") {
		panic(err)
	}

	return infoCard.ID.Int()
}

func initTestDocumentStorage(storage storage.DocumentStorage) int64 {
	document, err := storage.Create(context.TODO(), &dto.CreateDocumentRequest{
		SerialNumber: "123",
		InfoCardID:   ids["infoCardID"],
		DocumentType: 1,
	})
	if err != nil && !strings.Contains(err.Error(), "constraint") {
		panic(err)
	}

	return document.ID.Int()
}

func initTestCheckpointStorage(storage storage.CheckpointStorage) int64 {
	checkpoint, err := storage.CreateCheckpoint(context.TODO(), &dto.CreateCheckpointRequest{
		PhoneNumber: "123123",
	})
	if err != nil && !strings.Contains(err.Error(), "constraint") {
		panic(err)
	}

	return checkpoint.ID.Int()
}

func initTestPhotoMetaStorage(storage storage.PhotoMetaStorage) int64 {
	photoMeta, err := storage.SaveKey(context.TODO(), &dto.CreatePhotoKeyRequest{
		DocumentID: model.ToDocumentID(ids["documentID"]),
		Key:        model.ToPhotoKey("123321"),
	})
	if err != nil && !strings.Contains(err.Error(), "constraint") {
		panic(err)
	}

	return photoMeta.ID.Int()
}
