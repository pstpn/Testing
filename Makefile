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

.PHONY: test allure report