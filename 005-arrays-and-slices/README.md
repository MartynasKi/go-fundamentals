# 005: Arrays and slices

See when values are copied and when slices share storage.

## What a slice carries

An array assignment copies its elements. A slice describes a portion of an underlying array, with a length and a capacity. Appending within capacity can overwrite an element visible through another slice. Once the new length exceeds capacity, append allocates a new underlying array. Keep the slice that append returns.

## Run the examples

Run these commands from the repository root.

### [Array copy](examples/01-array-copy/main.go)

Edit an array copy and then edit a slice of the original.

```shell
go run ./005-arrays-and-slices/examples/01-array-copy
```

### [Append and capacity](examples/02-append-and-capacity/main.go)

Compare two views of the same array. Limiting capacity with a third slice index forces the second append to allocate.

```shell
go run ./005-arrays-and-slices/examples/02-append-and-capacity
```

## Further reading

[Arrays and slices in the Go documentation](https://go.dev/blog/slices-intro)
