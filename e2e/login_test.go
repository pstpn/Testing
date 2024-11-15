package e2e_test

import (
	"net/http"
	"testing"

	"github.com/cucumber/godog"
	"github.com/gavv/httpexpect/v2"
)

var (
	expectLogin *httpexpect.Expect
)

func TestLogin(t *testing.T) {
	client := &http.Client{}
	expectLogin = httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8084",
		Client:   client,
		Reporter: httpexpect.NewRequireReporter(nil),
	})

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeLoginScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/login.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run login feature tests")
	}
}

func loginWith2FA(ctx *godog.ScenarioContext) {
	var response *httpexpect.Response

	ctx.Step(`^User send "([^"]*)" request to "([^"]*)"$`, func(method, endpoint string) error {
		response = expectLogin.Request(method, endpoint).
			WithJSON(map[string]string{
				"email":    "stepaha78@gmail.com",
				"Password": "123123",
			}).
			Expect()
		return nil
	})

	ctx.Step(`^the response on /login code should be (\d+)$`, func(statusCode int) error {
		response.Status(statusCode)
		return nil
	})

	ctx.Step(`^the response on /login should match json:$`, func(expectedJSON *godog.DocString) error {
		response.JSON().Object().IsEqual(map[string]interface{}{
			"message": "Verification code sent to email",
		})
		return nil
	})

	ctx.Step(`^user send "([^"]*)" request to "([^"]*)"$`, func(method, endpoint string) error {
		response = expectLogin.Request(method, endpoint).
			WithJSON(map[string]string{
				"email": "stepaha78@gmail.com",
				"code":  "123456",
			}).Expect()
		return nil
	})

	ctx.Step(`^the response on /verify code should be (\d+)$`, func(statusCode int) error {
		response.Status(statusCode)
		return nil
	})

	ctx.Step(`^the response on /verify should match json:$`, func(expectedJSON *godog.DocString) error {
		response.JSON().Object().IsEqual(map[string]interface{}{
			"message": "Auth success!",
		})
		return nil
	})
}

func InitializeLoginScenario(ctx *godog.ScenarioContext) {
	loginWith2FA(ctx)
}
