# 998: Tips and tricks

Handle UTF-8 text and keep slice deletions from surprising a caller.

## Two details worth remembering

len counts bytes in a string. Ranging over it yields Unicode code points and their starting byte offsets. A displayed character can contain multiple code points. Slices have a different trap: helpers that delete elements can modify the input backing array. Clone first when a caller needs the original contents.

## Run the examples

Run these commands from the repository root.

### [Byte offsets](examples/01-byte-offsets/main.go)

Walk through café and compare byte length with rune count. The accented letter takes two bytes in UTF-8.

```shell
go run ./998-tips-and-tricks/examples/01-byte-offsets
```

### [Delete from copy](examples/02-delete-from-copy/main.go)

Remove a draft label from a cloned slice and print both lists.

```shell
go run ./998-tips-and-tricks/examples/02-delete-from-copy
```

## Further reading

[Tips and tricks in the Go documentation](https://go.dev/blog/strings)
