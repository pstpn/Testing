package service_test

import (
	"sync"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/internal/service"
	"course/internal/service/utils"
	"course/internal/storage/postgres"
	"course/pkg/jwt"
)

func TestRunner(t *testing.T) {
	db, ids := utils.NewTestStorage()
	defer utils.DropTestStorage(db)

	tm, err := jwt.NewManager("test")
	if err != nil {
		panic(err)
	}

	t.Parallel()

	wg := &sync.WaitGroup{}
	wg.Add(4)
	go func() {
		suite.RunSuite(t, &AuthSuite{
			authService: service.NewAuthService(
				utils.NewMockLogger(),
				postgres.NewEmployeeStorage(db),
				postgres.NewInfoCardStorage(db),
				tm,
				time.Hour,
				time.Hour,
			),
			companyID: ids["companyID"],
		})
		wg.Done()
	}()
	go func() {
		suite.RunSuite(t, &CheckpointSuite{})
		wg.Done()
	}()
	go func() {
		suite.RunSuite(t, &CompanySuite{
			companyService: service.NewCompanyService(
				utils.NewMockLogger(),
				postgres.NewCompanyStorage(db),
			),
			companyID: ids["companyID"],
		})
		wg.Done()
	}()
	go func() {
		suite.RunSuite(t, &DocumentSuite{})
		wg.Done()
	}()
	wg.Wait()
}
