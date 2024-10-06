test:
	rm -rf allure-results
	export ALLURE_OUTPUT_PATH="/Users/stepa/Study/Testing" && go test -tags=unit ./... --race --parallel 11
	cp environment.properties allure-results

allure:
	cp -R allure-reports/history allure-results
	rm -rf allure-reports
	allure generate allure-results -o allure-reports
	allure serve allure-results -p 4000

report: test allure

ci-unit:
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
 	export ALLURE_OUTPUT_FOLDER="unit-allure" && \
 	export DB_INIT_PATH="/home/runner/work/Testing/Testing/sql/init/init.sql" && \
 	go test -tags=unit ./... --race

local-unit:
	export ALLURE_OUTPUT_PATH="/Users/stepa/Study/Testing" && \
 	export DB_INIT_PATH="/Users/stepa/Study/Testing/sql/init/init.sql" && \
 	go test -tags=unit ./... --race

ci-integration:
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
	export ALLURE_OUTPUT_FOLDER="integration-allure" && \
 	export DB_INIT_PATH="/home/runner/work/Testing/Testing/sql/init/init.sql" && \
	go test -tags=integration ./... --race

local-integration:
	export ALLURE_OUTPUT_PATH="/Users/stepa/Study/Testing" && \
 	export DB_INIT_PATH="/Users/stepa/Study/Testing/sql/init/init.sql" && \
	go test -tags=integration ./... --race

ci-concat-reports:
	ls
	mkdir allure-results
	cp -R unit-allure/ allure-results/
	cp -R integration-allure/ allure-results/
	cp environment.properties allure-results
	ls allure-results

.PHONY: test allure report ci-unit local-unit ci-integration local-integration ci-concat-reports