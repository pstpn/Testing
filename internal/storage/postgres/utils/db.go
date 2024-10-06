package utils

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"

	"github.com/testcontainers/testcontainers-go"
	container "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	"course/internal/service/dto"
	"course/internal/storage"
	postgres2 "course/internal/storage/postgres"
	"course/pkg/storage/postgres"
)

const (
	pgImage     = "docker.io/postgres:16-alpine"
	initSQLPath = "/Users/stepa/Study/Testing/sql/init/init.sql"
	dbName      = "tests"
	dbUsername  = "postgres"
	dbPassword  = "admin"
)

var ids map[string]int64

func NewTestStorage() (*postgres.Postgres, *container.PostgresContainer, map[string]int64) {
	ctr, err := container.Run(
		context.TODO(),
		pgImage,
		container.WithInitScripts(initSQLPath),
		container.WithDatabase(dbName),
		container.WithUsername(dbUsername),
		container.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		panic(err)
	}

	connString, err := ctr.ConnectionString(context.TODO(), "sslmode=disable")
	if err != nil {
		panic(err)
	}

	conn, err := postgres.New(connString)
	if err != nil {
		panic(err)
	}

	ids = map[string]int64{}
	ids["companyID"] = initTestCompanyStorage(postgres2.NewCompanyStorage(conn))
	ids["employeeID"] = initTestEmployeeStorage(postgres2.NewEmployeeStorage(conn))
	ids["infoCardID"] = initTestInfoCardStorage(postgres2.NewInfoCardStorage(conn))
	ids["documentID"] = initTestDocumentStorage(postgres2.NewDocumentStorage(conn))
	ids["checkpointID"] = initTestCheckpointStorage(postgres2.NewCheckpointStorage(conn))

	return conn, ctr, ids
}

func DropTestStorage(testDB *postgres.Postgres, ctr *container.PostgresContainer) {
	defer func() {
		testDB.Close()
		ctr.Terminate(context.TODO())
	}()

	err := postgres2.NewCheckpointStorage(testDB).DeleteCheckpoint(context.TODO(), &dto.DeleteCheckpointRequest{CheckpointID: ids["checkpointID"]})
	if err != nil {
		panic(err)
	}
	err = postgres2.NewDocumentStorage(testDB).Delete(context.TODO(), &dto.DeleteDocumentRequest{DocumentID: ids["documentID"]})
	if err != nil {
		panic(err)
	}
	err = postgres2.NewInfoCardStorage(testDB).Delete(context.TODO(), &dto.DeleteInfoCardRequest{InfoCardID: ids["infoCardID"]})
	if err != nil {
		panic(err)
	}
	err = postgres2.NewEmployeeStorage(testDB).Delete(context.TODO(), &dto.DeleteEmployeeRequest{EmployeeID: ids["employeeID"]})
	if err != nil {
		panic(err)
	}
	err = postgres2.NewCompanyStorage(testDB).Delete(context.TODO(), &dto.DeleteCompanyRequest{CompanyID: ids["companyID"]})
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
	expiredAt, _ := time.Parse(time.RFC3339, "2008-01-02T00:00:00Z")
	tm, _ := time.Parse(time.RFC3339, "2010-01-02T00:00:00Z")
	pass, _ := bcrypt.GenerateFromPassword([]byte("21e12"), bcrypt.DefaultCost)

	employee, err := storage.Register(context.TODO(), &dto.RegisterEmployeeRequest{
		PhoneNumber:    "500500500",
		FullName:       "123",
		CompanyID:      ids["companyID"],
		Post:           1,
		Password:       string(pass),
		RefreshToken:   "974998",
		TokenExpiredAt: &expiredAt,
		DateOfBirth:    &tm,
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
		SerialNumber: "123923",
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
