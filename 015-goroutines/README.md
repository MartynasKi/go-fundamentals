# 015: Goroutines

Start work concurrently and wait for it to finish.

## Who waits for the worker?

A go statement starts a function in a new goroutine. Returning from main ends the program, even if other goroutines still have work to do. A WaitGroup lets the caller wait for completion. Sleeping for a guessed duration makes the result depend on scheduling.

## Run the examples

Run these commands from the repository root.

### [Wait for workers](examples/01-wait-for-workers/main.go)

Give each worker its own slice element. Read the completed results in input order after Wait returns.

```shell
go run ./015-goroutines/examples/01-wait-for-workers
```

## Further reading

[Goroutines in the Go documentation](https://pkg.go.dev/sync#WaitGroup.Go)
