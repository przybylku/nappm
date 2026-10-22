# nappm

A new project written in Go.

## Requirements

- Go 1.27 or newer

## Running

```sh
go run .
```

## Building

```sh
make build
# or
go build -o bin/nappm .
```

## Testing

```sh
make test
# or
go test ./...
```

## Node.js test server

A small dependency-free HTTP server is included for testing:

```sh
node server.js
```

It listens on `http://127.0.0.1:3000` by default. Use `PORT` and `HOST` to override:

```sh
PORT=8080 HOST=0.0.0.0 node server.js
```

Available endpoints:

- `/` — request details as JSON
- `/health` — `{"status":"ok"}`

## Project structure

```
.
├── main.go        # application entry point
├── main_test.go   # tests
├── server.js      # simple Node.js HTTP test server
├── Makefile       # common tasks
└── go.mod         # module definition
```
