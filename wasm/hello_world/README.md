# Hello world to Javscript from Golang using **WebAssembly**

Golang does support compiling go code to web assembly natively

### You need to have the following js code:
```
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" assets/
```

### Compile golang code to wasm
```
GOOS=js GOARCH=wasm go build -o  assets/hello_world.wasm

or

make build
```

### Run the project
```
go run cmd/server/main.go

or

make run
```