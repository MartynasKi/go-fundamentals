# 006: Maps

Distinguish a missing key from a stored zero and visit keys in order.

## Reading a missing key

A lookup returns the zero value of the element type when the key is absent. Ask for the second result when presence matters. Reading a nil map works too, but writing to one panics. Initialize it with a literal or make before adding entries.

## Run the examples

Run these commands from the repository root.

### [Stock lookup](examples/01-stock-lookup/main.go)

Check a sold-out item and an unknown item. Their counts match but their presence flags differ.

```shell
go run ./006-maps/examples/01-stock-lookup
```

### [Ordered keys](examples/02-ordered-keys/main.go)

Map iteration order is unspecified. Collect and sort the keys for a repeatable report.

```shell
go run ./006-maps/examples/02-ordered-keys
```

## Further reading

[Maps in the Go documentation](https://go.dev/blog/maps)
