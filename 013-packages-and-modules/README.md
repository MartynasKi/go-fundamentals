# 013: Packages and modules

Import a local package through the root module path.

## From import path to directory

A module groups packages under the path declared in go.mod. Within this module, append the package directory to that path to import it. The package name qualifies its exported identifiers. Here, the capital H in Hello makes the function accessible from main. Putting greeting in its own directory keeps the two package declarations separate.

## Run the examples

Run these commands from the repository root.

### [Import a package](examples/01-import-a-package/main.go)

Call an exported function in a neighboring package. Follow the import path to see where it lives.

[greeting/greeting.go](examples/01-import-a-package/greeting/greeting.go)

```shell
go run ./013-packages-and-modules/examples/01-import-a-package
```

## Further reading

[Packages and modules in the Go documentation](https://go.dev/doc/code)
