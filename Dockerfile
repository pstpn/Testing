FROM go_env:latest

COPY . .

CMD ["go", "run", "prometheus/core.go"]