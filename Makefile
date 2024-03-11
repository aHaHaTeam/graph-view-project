all: build run

build:
	GOOS=js GOARCH=wasm go build -o ./client/static/main.wasm ./wasm/*.go

run:
	go run server/main.go

clean:
	rm ./client/main.wasm
