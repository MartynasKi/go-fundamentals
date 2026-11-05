# 019: Memory and garbage collection

Observe how references keep data alive and what heap statistics measure.

## What stays reachable

A small slice can keep a large backing array reachable. Copying the portion you need lets the old array become eligible for collection once its other references are gone. The runtime decides when to collect. HeapAlloc is a process-wide byte count and includes runtime activity, so its exact value varies.

## Run the examples

Run these commands from the repository root.

### [Keep a prefix](examples/01-keep-a-prefix/main.go)

Clone a short prefix before reusing the input buffer. Compare the shared view with the independent copy.

```shell
go run ./019-memory-and-garbage-collection/examples/01-keep-a-prefix
```

### [Live buffer](examples/02-live-buffer/main.go)

Read HeapAlloc while a one-mebibyte buffer is live. KeepAlive keeps the buffer reachable through the measurement. The byte count varies.

```shell
go run ./019-memory-and-garbage-collection/examples/02-live-buffer
```

## Further reading

[Memory and garbage collection in the Go documentation](https://go.dev/doc/gc-guide)
