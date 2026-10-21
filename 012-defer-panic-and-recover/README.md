# 012: Defer, panic and recover

Trace deferred calls during a return and a panic.

## When deferred calls run

A defer evaluates the call arguments immediately and runs the call when the surrounding function exits. Deferred calls run in reverse registration order. During a panic, a deferred function can call recover directly to stop the panic in that goroutine. Execution then returns to the caller of the recovering function.

## Run the examples

Run these commands from the repository root.

### [Deferred arguments](examples/01-deferred-arguments/main.go)

Compare a deferred argument with a closure that reads the variable at exit.

```shell
go run ./012-defer-panic-and-recover/examples/01-deferred-arguments
```

### [Recover](examples/02-recover/main.go)

See where control resumes after a deferred function recovers a panic.

```shell
go run ./012-defer-panic-and-recover/examples/02-recover
```

## Further reading

[Defer, panic and recover in the Go documentation](https://go.dev/blog/defer-panic-and-recover)
