package e2e_test

import (
	"net/http"
	"testing"

	"github.com/cucumber/godog"
	"github.com/gavv/httpexpect/v2"
)

var (
	expectReset *httpexpect.Expect
)

func TestResetPassword(t *testing.T) {
	client := &http.Client{}
	expectReset = httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8111",
		Client:   client,
		Reporter: httpexpect.NewRequireReporter(nil),
	})

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeResetPasswordScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/reset.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run reset password feature tests")
	}
}

func resetPasswordWith2FA(ctx *godog.ScenarioContext) {
	var response *httpexpect.Response

	ctx.Step(`^User send "([^"]*)" request to "([^"]*)"$`, func(method, endpoint string) error {
		response = expectReset.Request(method, endpoint).
			WithJSON(map[string]string{
				"email":       "stepaha78@gmail.com",
				"oldPassword": "123123",
			}).
			Expect()
		return nil
	})

	ctx.Step(`^the response on /reset_password code should be (\d+)$`, func(statusCode int) error {
		response.Status(statusCode)
		return nil
	})

	ctx.Step(`^the response on /reset_password should match json:$`, func(expectedJSON *godog.DocString) error {
		response.JSON().Object().IsEqual(map[string]interface{}{
			"message": "Password reset code sent to email",
		})
		return nil
	})

	ctx.Step(`^user send "([^"]*)" request to "([^"]*)"$`, func(method, endpoint string) error {
		response = expectReset.Request(method, endpoint).
			WithJSON(map[string]string{
				"email":       "stepaha78@gmail.com",
				"code":        "123456",
				"newPassword": "123",
			}).Expect()
		return nil
	})

	ctx.Step(`^the response on /verify_reset_password code should be (\d+)$`, func(statusCode int) error {
		response.Status(statusCode)
		return nil
	})

	ctx.Step(`^the response on /verify_reset_password should match json:$`, func(expectedJSON *godog.DocString) error {
		response.JSON().Object().IsEqual(map[string]interface{}{
			"message": "Password changed successfully!",
		})
		return nil
	})
}

func InitializeResetPasswordScenario(ctx *godog.ScenarioContext) {
	resetPasswordWith2FA(ctx)
}
