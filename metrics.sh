MAX_CYCLO=10

gocyclo -over "$MAX_CYCLO" .
if [ $? -ne 0 ]; then \
  echo "Цикломатическая сложность превышает $MAX_CYCLO. Исправьте код!\n"; \
  exit 1; \
fi
echo "Проверка цикломатической сложности: OK\n"

go run metrics.go
echo "\nПроверка сложности по Холстеду: OK"

staticcheck ./... || exit 1
echo "\nПроверка staticcheck: OK"

echo "\nВсе проверки пройдены!"
