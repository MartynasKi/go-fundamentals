# Go Fundamentals

First of all... Hello!

I’m Martynas Ki. 👋 I’ve spent a good part of my career writing PHP. Learning Go gave me a reason to revisit questions I usually take for granted. When I pass a slice to a function, for example, which changes can the caller see?

These are my notes and the little programs I use to check my understanding. I also share things I learn on my [Laraholic blog](https://laraholic.com/).

## Requirements

[Go](https://go.dev/dl/) 1.25.0 or newer. All examples use the standard library.

## How to run

Open a terminal in the repository root and run one example by its directory name.

```shell
go run ./001-getting-started/examples/01-hello
```

Each topic links to its source files and includes the commands to try.

## Topics

- [001: Getting started](001-getting-started/README.md) — Run a Go program and pass it an argument.
- [002: Variables and types](002-variables-and-types/README.md) — Explore zero values, named types and explicit conversions.
- [003: Control flow](003-control-flow/README.md) — Use loops, early exits and expressionless switches.
- [004: Functions](004-functions/README.md) — Return multiple values and keep state in a closure.
- [005: Arrays and slices](005-arrays-and-slices/README.md) — See when values are copied and when slices share storage.
- [006: Maps](006-maps/README.md) — Distinguish a missing key from a stored zero and visit keys in order.
- [007: Structs](007-structs/README.md) — Group related fields and see what a struct copy shares.
- [008: Pointers](008-pointers/README.md) — Follow an address and change a caller’s value.
