all: build run

build:
	GOOS=js GOARCH=wasm go build -o ./client/static/main.wasm ./wasm/*.go

run:
	go run server/main.go

test:
	go test -C ./ ./server/test/api_test.go

clean:
	rm ./client/main.wasm
