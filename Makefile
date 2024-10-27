run:
	docker build -t go_env:latest . -f env.dockerfile
	docker compose up -d

rerun-gatling:
	docker stop gatling && docker rm gatling
	docker compose up gatling

rm:
	docker compose down
	docker image rm testing-collector:latest go_env:latest

gatling-delta:
	java -jar ~/Downloads/gatling-report-6.1-capsule-fat.jar gatling/results/serverloadsimulation-20241027104451296/simulation.log gatling/results/serverloadsimulation-20241027104530219/simulation.log gatling/results/serverloadsimulation-20241027105037583/simulation.log -o test

.PHONY: s-bench run rerun-gatling rm gatling-delta