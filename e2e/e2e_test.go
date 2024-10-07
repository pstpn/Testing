//go:build e2e

package e2e_test

import (
	"net/http"
	"sync"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/e2e"
)

type E2ESuite struct {
	suite.Suite

	e httpexpect.Expect
}

func (s *E2ESuite) BeforeAll(t provider.T) {
	s.e = *httpexpect.WithConfig(httpexpect.Config{
		Client:   &http.Client{},
		BaseURL:  "http://localhost:8044/api/v2",
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
}

type registerRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	CompanyID   int64  `json:"companyID"`
	Post        string `json:"post"`
	Password    string `json:"password"`
	DateOfBirth string `json:"dateOfBirth"`
}

type fillProfileRequest struct {
	DocumentSerialNumber string `json:"serialNumber"`
	DocumentType         string `json:"documentType"`
	DocumentFields       string `json:"documentFields"`
}

func (s *E2ESuite) Test_E2E(t provider.T) {
	t.Title("[E2E] E2E test")
	t.Tags("e2e")
	t.Parallel()
	t.WithNewStep("E2E test", func(sCtx provider.StepCtx) {
		registerReq := &registerRequest{
			PhoneNumber: "+7112356",
			Name:        "Stepa",
			Surname:     "Stepik",
			CompanyID:   1,
			Post:        "Сотрудник",
			Password:    "123",
			DateOfBirth: "31.03.2004",
		}

		// Assert register response
		accessToken := s.e.POST("/register").
			WithJSON(registerReq).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			NotEmpty().
			ContainsKey("accessToken").
			ContainsKey("refreshToken").
			ContainsKey("isAdmin").
			HasValue("isAdmin", false).
			Value("accessToken").Raw().(string)

		fillProfileReq := &fillProfileRequest{
			DocumentSerialNumber: "123",
			DocumentType:         "Паспорт",
			DocumentFields:       "Дата выдачи,31.03.2010;Выдавший орган,ГУ ГУ ГУ",
		}

		// Assert fill profile response
		s.e.POST("/profile").
			WithHeader("Authorization", "Bearer "+accessToken).
			WithJSON(fillProfileReq).
			Expect().
			Status(http.StatusCreated)

		// Assert get profile response
		s.e.GET("/profile").
			WithHeader("Authorization", "Bearer "+accessToken).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			NotEmpty().
			ContainsKey("isConfirmed").
			ContainsKey("createdAt").
			ContainsKey("documentType").
			ContainsKey("serialNumber").
			ContainsKey("documentFields").
			HasValue("isConfirmed", false).
			HasValue("documentType", "Паспорт").
			HasValue("serialNumber", "123")
	})
}

func TestRunner(t *testing.T) {
	db, _ := e2e.NewTestStorage()
	t.Cleanup(func() {
		e2e.DropTestStorage(db)
	})

	t.Parallel()

	wg := &sync.WaitGroup{}
	suits := []runner.TestSuite{
		&E2ESuite{},
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
