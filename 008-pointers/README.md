# 008: Pointers

Follow an address and change a caller’s value.

## Passing an address

Go copies function arguments. A pointer argument is a copy of an address, so dereferencing it reaches the same variable the caller sees. A pointer has a nil zero value. Dereferencing nil panics.

## Run the examples

Run these commands from the repository root.

### [Basic pointer](examples/01-basic-pointer/main.go)

Use & to take an address and * to read and update the value there.

```shell
go run ./008-pointers/examples/01-basic-pointer
```

### [Pass by value](examples/02-pass-by-value/main.go)

Write zero through the pointer, then clear the local parameter. The caller still has a pointer to the updated count.

```shell
go run ./008-pointers/examples/02-pass-by-value
```

## Further reading

[Pointers in the Go documentation](https://go.dev/tour/moretypes/1)
