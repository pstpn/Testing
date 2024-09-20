package postgres_test

import (
	"sync"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/internal/storage/postgres"
	"course/internal/storage/postgres/utils"
)

func TestRunner(t *testing.T) {
	db, ids := utils.NewTestStorage()
	defer utils.DropTestStorage(db)

	t.Parallel()

	wg := &sync.WaitGroup{}
	suits := []runner.TestSuite{
		&CompanyStorageSuite{
			companyStorage: postgres.NewCompanyStorage(db),
		},
		&CheckpointStorageSuite{
			checkpointStorage: postgres.NewCheckpointStorage(db),
			checkpointID:      ids["checkpointID"],
			documentID:        ids["documentID"],
		},
		&DocumentStorageSuite{
			documentStorage: postgres.NewDocumentStorage(db),
			infoCardID:      ids["infoCardID"],
			documentID:      ids["documentID"],
		},
		&EmployeeStorageSuite{
			employeeStorage: postgres.NewEmployeeStorage(db),
			employeeID:      ids["employeeID"],
			companyID:       ids["companyID"],
			infoCardID:      ids["infoCardID"],
		},
		&FieldStorageSuite{
			fieldStorage: postgres.NewFieldStorage(db),
			documentID:   ids["documentID"],
		},
		&InfoCardStorageSuite{
			infoCardStorage: postgres.NewInfoCardStorage(db),
			employeeID:      ids["employeeID"],
		},
	}
	wg.Add(len(suits))

	for _, s := range suits {
		go func() {
			suite.RunSuite(t, s)
			wg.Done()
		}()
	}

	wg.Wait()
}
