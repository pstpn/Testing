# <span style="color:#C0BFEC">🦔 ***Тестирование***</span>

## <span style="color:#C0BFEC">📑 ***Описание:*** </span>

Лабораторные работы по курсу "Тестирование"

## <span style="color:#C0BFEC">⚙️ ***Описание Makefile:*** </span>

```makefile
# Запуск всех unit тестов проекта
test:
	rm -rf allure-results
	export ALLURE_OUTPUT_PATH="/Users/stepa/Study/testingpsa" && go test ./... --race --parallel 11
	cp environment.properties allure-results

# Создание allure отчета по результатам теста
allure:
	cp -R allure-reports/history allure-results
	rm -rf allure-reports
	allure generate allure-results -o allure-reports
	allure serve allure-results -p 4000

# Запуск всех unit тестов проекта с последующим созданием allure отчета
report: test allure

.PHONY: test allure report
```

## <span style="color:#C0BFEC">🏃🏻‍♂️ ***Запуск:*** </span>

1) Поменять в `MakeFile` путь до выходной директории allure:
```makefile
export ALLURE_OUTPUT_PATH=<ВАШ_ПУТЬ> && go test ./... --race --parallel 11
```
2) Выполнить в папке с `Makefile`:
```shell
make report -i
```