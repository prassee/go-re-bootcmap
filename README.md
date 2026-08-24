# go-re-bootcamp

A bootcamp project for practicing Go fundamentals: structs, generics, interfaces, concurrency, and atomics.

## Structure

- [main.go](main.go) — entry point and scratch examples (structs, atomics, logging)
- [logger/](logger/) — `Logger` interface with `ConsoleLogger` and `FileLogger` implementations
- [conc/](conc/) — goroutines, channels, and a worker-pool example (`CoOrdinate`)
- [utils/](utils/) — generic helpers (`GIMax`, `Box[T]`)

## Requirements

- Go 1.27+

## Run

```sh
go run .
```
