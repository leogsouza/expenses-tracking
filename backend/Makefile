GO_CMD=go
GO_BUILD=${GO_CMD} build
BINARY_NAME=tracking
BIN_DIR=bin
clean:
	rm -rf bin/*
compose_up:
	docker compose up -d
compose_down:
	docker compose down
compose_start:
	docker compose start -d
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GO_BUILD} -o ${BIN_DIR}/${BINARY_NAME}
build-windows:
	GOOS=windows GOARCH=amd64 ${GO_BUILD} -o ${BIN_DIR}/${BINARY_NAME}.exe
build: build-linux build-windows

execute:
	./${BIN_DIR}/${BINARY_NAME}

run: build execute

test:
	go test -cover -race ./...