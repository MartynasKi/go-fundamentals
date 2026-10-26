# 014: Generics

Write a function for several types with one constraint.

## Which operations are available

A type parameter stands for a type chosen at the call site. Its constraint determines the operations the function body can use. The comparable constraint permits equality checks, so the same search works on strings and integers. A slice itself cannot be compared to another slice for equality.

## Run the examples

Run these commands from the repository root.

### [Find index](examples/01-find-index/main.go)

Find the first match in two kinds of slice. The compiler infers the element type from each call.

```shell
go run ./014-generics/examples/01-find-index
```

## Further reading

[Generics in the Go documentation](https://go.dev/doc/tutorial/generics)
