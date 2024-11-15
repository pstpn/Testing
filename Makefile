gen-swag:
	swag init --parseDependency --dir ./internal/controller/v1/http,./internal/controller --instanceName v1 -o ./docs/v1

fmt-swag:
	swag f

swagger: fmt-swag gen-swag

run:
	go run cmd/main.go

docker-build:
	docker build -t bee .

docker-run: docker-build
	docker run -d --name bee -p 8081:8081 bee:latest

docker-rm:
	docker stop bee
	docker rm bee
	docker image rm bee:latest

check:
	./metrics.sh

.PHONY: gen-swag fmt-swag swagger run docker-build docker-run docker-rm check