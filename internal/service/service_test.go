//go:build unit

package service_test

import (
	"sync"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

func TestRunner(t *testing.T) {
	t.Parallel()

	wg := &sync.WaitGroup{}
	suits := []runner.TestSuite{
		&AuthSuite{},
		&CheckpointSuite{},
		&CompanySuite{},
		&DocumentSuite{},
		&EmployeeSuite{},
		&FieldSuite{},
		&InfoCardSuite{},
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
