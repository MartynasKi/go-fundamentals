# 002: Variables and types

Explore zero values, named types and explicit conversions.

## What a declaration gives you

A variable has a type even when you leave out its initial value. Inside a function, `:=` infers the type from the expression. A named type can give a number a meaning of its own. Converting a floating-point value to an integer discards the fractional part.

## Run the examples

Run these commands from the repository root.

### [Zero values](examples/01-zero-values/main.go)

Inspect the starting values of a counter, a label and a flag. Quoting the label makes its empty string visible.

```shell
go run ./002-variables-and-types/examples/01-zero-values
```

### [Temperature](examples/02-temperature/main.go)

Convert a measured temperature to a whole number for display.

```shell
go run ./002-variables-and-types/examples/02-temperature
```

## Further reading

[Variables and types in the Go documentation](https://go.dev/ref/spec#Variables)
