# 009: Methods

Choose a receiver that matches the operation.

## The receiver is an argument

A value receiver works on a copy. A pointer receiver can update the original value. When a variable is addressable, Go lets you call its pointer method with ordinary dot syntax. That convenience does not add pointer methods to the value type’s method set.

## Run the examples

Run these commands from the repository root.

### [Scoreboard](examples/01-scoreboard/main.go)

Preview a bonus on a copy, then record it on the scoreboard.

```shell
go run ./009-methods/examples/01-scoreboard
```

## Further reading

[Methods in the Go documentation](https://go.dev/ref/spec#Method_sets)
