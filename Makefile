APP_NAME=ekak_backend_sleman

.PHONY: build run clean test swagger build-image format dev check-format

build:
	go build .

format:
	goimports -w .

run:
	go run .

dev: format run

test:
	go test -race ./...

check-format:
	@test -z "$$(gofmt -l .)" || (echo "gofmt required:" && gofmt -l . && exit 1)
	@test -z "$$(goimports -l .)" || (echo "goimports required:" && goimports -l . && exit 1)

swagger:
	swag init

clean:
	rm -f ./$(APP_NAME)

build-image:
	@docker build . -t $(APP_NAME)
