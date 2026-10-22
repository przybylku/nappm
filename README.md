# nappm

Nowy projekt w języku Go.

## Wymagania

- Go 1.27 lub nowszy

## Uruchomienie

```sh
go run .
```

## Budowanie

```sh
make build
# lub
go build -o bin/nappm .
```

## Testy

```sh
make test
# lub
go test ./...
```

## Struktura

```
.
├── main.go        # punkt wejścia aplikacji
├── main_test.go   # testy
├── Makefile       # typowe zadania
└── go.mod         # definicja modułu
```
