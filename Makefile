build:
	GOOS=js GOARCH=wasm go build -o ./client/main.wasm ./wasm/*.go

clean:
	rm ./client/main.wasm
