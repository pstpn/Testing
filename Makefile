test:
	rm -rf allure-results
	export ALLURE_OUTPUT_PATH="/Users/stepa/Study/testingpsa" && go test ./... --race --parallel 11
	cp environment.properties allure-results

allure:
	cp -R allure-reports/history allure-results
	rm -rf allure-reports
	allure generate allure-results -o allure-reports
	allure serve allure-results -p 4000

report: test allure

ci-unit:
	export ALLURE_OUTPUT_PATH="/app/" && go test ./... --race

ci-integration:
	export ALLURE_OUTPUT_PATH="/app/" && go test -tags=integration ./... --race

build-ci:
	docker build -t testing .

run-ci: build-ci
	docker run --name testing testing:latest && docker cp testing:/app/allure-results .

rm-ci:
	docker rm testing
	docker image rm testing:latest

.PHONY: test allure report ci-unit ci-integration build-ci run-ci rm-ci