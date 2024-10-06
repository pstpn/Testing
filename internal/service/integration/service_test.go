//go:build integration
// +build integration

package service_test

import (
	"sync"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

func TestRunner(t *testing.T) {
	//db, ids := utils.NewTestStorage()
	//defer utils.DropTestStorage(db)
	//
	//tm, err := jwt.NewManager("test")
	//if err != nil {
	//	panic(err)
	//}

	t.Parallel()

	wg := &sync.WaitGroup{}
	suits := []runner.TestSuite{
		//&AuthSuite{
		//	authService: service.NewAuthService(
		//		utils.NewMockLogger(),
		//		postgres.NewEmployeeStorage(db),
		//		postgres.NewInfoCardStorage(db),
		//		tm,
		//		time.Hour,
		//		time.Hour,
		//	),
		//	companyID: ids["companyID"],
		//},
		//&CheckpointSuite{},
		//&CompanySuite{
		//	companyService: service.NewCompanyService(
		//		utils.NewMockLogger(),
		//		postgres.NewCompanyStorage(db),
		//	),
		//	companyID: ids["companyID"],
		//},
		//&DocumentSuite{},
		//&EmployeeSuite{},
		//&FieldSuite{},
		//&InfoCardSuite{},
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
