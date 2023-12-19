WASM_DIR := cmd/wasm
OUTPUT_DIR := assets
OUTPUT := hello_world.wasm


build: cmd/wasm/*.go
	GOOS=js GOARCH=wasm go build -o ${OUTPUT_DIR}/${OUTPUT} ${WASM_DIR}/*

run: build
	go run cmd/server/main.go