# 011: Error handling

Wrap an error with context and find its cause.

## Following an error chain

Return an error alongside the result when an operation can fail. Check it before using the result. A %w verb in fmt.Errorf preserves the cause for errors.Is or errors.As. Formatting the same error with %v keeps its text but loses that link.

## Run the examples

Run these commands from the repository root.

### [Missing entry](examples/01-missing-entry/main.go)

Look up a missing setting in a map. The caller prints the failed key and recognizes the sentinel error.

```shell
go run ./011-error-handling/examples/01-missing-entry
```

### [Number error](examples/02-number-error/main.go)

Extract the numeric parsing error to inspect the input that failed.

```shell
go run ./011-error-handling/examples/02-number-error
```

## Further reading

[Error handling in the Go documentation](https://go.dev/blog/go1.13-errors)
