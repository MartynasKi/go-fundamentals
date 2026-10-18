# 010: Interfaces

Use a small interface and inspect a typed nil.

## What an interface value holds

A type satisfies an interface by providing its methods. At runtime an interface value has a dynamic type and a dynamic value. For the interface to equal nil, both must be absent. This matters when a function returns a pointer as an error.

## Run the examples

Run these commands from the repository root.

### [Write a label](examples/01-write-a-label/main.go)

Write a label to a buffer through io.Writer. The caller chooses where the bytes go.

```shell
go run ./010-interfaces/examples/01-write-a-label
```

### [Typed nil](examples/02-typed-nil/main.go)

Assign a nil pointer to an error, then explicitly clear the interface.

```shell
go run ./010-interfaces/examples/02-typed-nil
```

## Further reading

[Interfaces in the Go documentation](https://go.dev/doc/faq#nil_error)
