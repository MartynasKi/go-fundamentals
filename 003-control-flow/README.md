# 003: Control flow

Use loops, early exits and expressionless switches.

## How a loop chooses its next step

A range loop can give you both an index and a value. Use `continue` to skip one iteration or `break` to leave the loop. In an expressionless switch, Go selects the first true case, so put narrower conditions first.

## Run the examples

Run these commands from the repository root.

### [Skip empty](examples/01-skip-empty/main.go)

Print only the nonempty names from a list.

```shell
go run ./003-control-flow/examples/01-skip-empty
```

### [Delivery window](examples/02-delivery-window/main.go)

Choose one delivery message even when two conditions are true.

```shell
go run ./003-control-flow/examples/02-delivery-window
```

## Further reading

[Control flow in the Go documentation](https://go.dev/ref/spec#For_statements)
