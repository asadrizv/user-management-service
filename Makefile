.PHONY: test start start-build

unit-test:
	go test ./...

start:
	docker-compose up

start-build:
	docker-compose up --build

integration-test:
	sh test/integration/start.sh

