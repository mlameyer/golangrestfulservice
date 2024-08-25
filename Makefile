.PHONY: clean test security build run

APP_NAME = carrierapi
BUILD_DIR = $(PATH)/build

hello:
	echo "hello"

clean:
	rm -rf golangrestfulservice/service/src/carrier-service/build


security:
	gosec -quiet ./...

test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build:
	go build service/src/carrier-service/main.go service/src/carrier-service/environment.go
