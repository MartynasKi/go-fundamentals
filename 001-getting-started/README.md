# 001: Getting started

Run a Go program and pass it an argument.

## Where execution starts

An executable belongs to package `main`. After package initialization, Go calls its `main` function. Run the example from the repository root so the relative path points to the right package. Arguments after that path belong to your program.

## Run the examples

Run these commands from the repository root.

### [Hello](examples/01-hello/main.go)

Print a greeting using the standard library.

```shell
go run ./001-getting-started/examples/01-hello
```

### [Command line](examples/02-command-line/main.go)

Pass a name on the command line. The first element of os.Args is the executable name, so the first user argument is at index 1.

```shell
go run ./001-getting-started/examples/02-command-line Martynas
```

## Further reading

[Getting started in the Go documentation](https://go.dev/doc/tutorial/getting-started)
