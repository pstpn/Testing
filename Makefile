s-bench:
	go test -bench=. ./... > s-bench.txt

run:
	docker build -t go_env:latest . -f env.dockerfile
	docker compose up

rm:
	docker compose down
	#docker image rm

.PHONY: s-bench run rm