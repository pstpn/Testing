go-tests-template:
	gotests -all -w <path>

mockery-template:
	mockery --all --inpackage

go-test-command:
	go test -v ./...

test:
	go test ./internal/service/auth_test.go --race

allure:
	allure serve internal/service/allure-results -p 4000

.PHONY: go-tests-template mockery-template go-test-command test allure