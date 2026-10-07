# 004: Functions

Return multiple values and keep state in a closure.

## What a function can return

A caller can receive several results from one call. A function can also return another function. If the inner function refers to a local variable, it keeps access to that variable after the outer call returns. Each new call to the outer function gets its own local state.

## Run the examples

Run these commands from the repository root.

### [Cut label](examples/01-cut-label/main.go)

Split a label at the first colon and check whether the separator was present.

```shell
go run ./004-functions/examples/01-cut-label
```

### [Counter](examples/02-counter/main.go)

Call one counter twice, then start another counter.

```shell
go run ./004-functions/examples/02-counter
```

## Further reading

[Functions in the Go documentation](https://go.dev/ref/spec#Function_literals)
