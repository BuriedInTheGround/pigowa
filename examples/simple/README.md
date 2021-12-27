# example: mouse events

## Quickstart

Use the following command to build the WASM binary file.
```shell
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```

From the root directory of the repo, execute the following command to start an
HTTP server on this directory.
```shell
go run ./server -dir ./examples/simple
```

Open your browser and go to `localhost:8080`. Enjoy!
