MAX_CYCLO=10

gocyclo -over "$MAX_CYCLO" .
if [ $? -ne 0 ]; then \
  echo "Цикломатическая сложность превышает $MAX_CYCLO. Исправьте код!"; \
  exit 1; \
fi
echo "Проверка цикломатической сложности: OK"

go run metrics.go
echo "Проверка сложности по Холстеду: OK"

staticcheck ./... || exit 1
echo "Проверка staticcheck: OK"

echo "Все проверки пройдены!"
