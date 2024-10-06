//go:build integration

package service_test

import (
	"sync"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/internal/service"
	"course/internal/service/utils"
	"course/internal/storage/postgres"
	"course/pkg/jwt"
)

func TestRunner(t *testing.T) {
	db, ctr, ids := utils.NewTestStorage()
	defer utils.DropTestStorage(db, ctr)

	tm, err := jwt.NewManager("test")
	if err != nil {
		panic(err)
	}

	t.Parallel()

	wg := &sync.WaitGroup{}
	suits := []runner.TestSuite{
		&AuthSuite{
			authService: service.NewAuthService(
				utils.NewMockLogger(),
				postgres.NewEmployeeStorage(db),
				postgres.NewInfoCardStorage(db),
				tm,
				time.Hour,
				time.Hour,
			),
			companyID: ids["companyID"],
		},
		&CompanySuite{
			companyService: service.NewCompanyService(
				utils.NewMockLogger(),
				postgres.NewCompanyStorage(db),
			),
			companyID: ids["companyID"],
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