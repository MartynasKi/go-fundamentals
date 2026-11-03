# 018: Context

Cancel a worker and observe its exit.

## Cancellation is a signal

A context carries cancellation and deadlines across calls. Calling cancel closes Done. The worker must observe that signal and return on its own. Waiting for a separate completion signal lets the caller know the worker finished its cleanup.

## Run the examples

Run these commands from the repository root.

### [Cancel a worker](examples/01-cancel-a-worker/main.go)

Start a worker, cancel its context and wait for it to report the reason it exited.

```shell
go run ./018-context/examples/01-cancel-a-worker
```

## Further reading

[Context in the Go documentation](https://go.dev/blog/context)
