# 007: Structs

Group related fields and see what a struct copy shares.

## Copying a record

A struct literal can name the fields it sets. Omitted fields take their zero values. Assigning a struct copies every field, including any slice descriptors. The two records can then refer to the same backing array.

## Run the examples

Run these commands from the repository root.

### [Copy a record](examples/01-copy-a-record/main.go)

Rename a copied draft, then edit one of its tags. Only the tag change reaches the original.

```shell
go run ./007-structs/examples/01-copy-a-record
```

## Further reading

[Structs in the Go documentation](https://go.dev/ref/spec#Struct_types)
