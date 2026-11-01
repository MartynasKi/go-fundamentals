# 017: Select and synchronization

Check whether a channel is ready and protect shared updates.

## Choosing between communication and locking

A select executes one ready channel operation. If several are ready, source order gives none of them priority. A default case runs immediately when none is ready. For shared state, a mutex can protect a whole read-modify-write operation. A WaitGroup alone only tells you when the workers finish.

## Run the examples

Run these commands from the repository root.

### [Ready or empty](examples/01-ready-or-empty/main.go)

Poll a buffered inbox first while it has a message and again when it is empty.

```shell
go run ./017-select-and-synchronization/examples/01-ready-or-empty
```

### [Shared count](examples/02-shared-count/main.go)

Two workers increment the same counter under one mutex. Read the total after both finish.

```shell
go run ./017-select-and-synchronization/examples/02-shared-count
```

## Further reading

[Select and synchronization in the Go documentation](https://go.dev/ref/spec#Select_statements)
