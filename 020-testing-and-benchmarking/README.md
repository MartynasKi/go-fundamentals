# 020: Testing and benchmarking

Check behavior with table tests and measure one operation.

## Two different questions

A table test runs the same assertion against several inputs. A named subtest tells you which case failed. A benchmark repeats an operation with testing.B.Loop and reports time per operation. CPU, operating system and other work on the machine affect timing. Compare repeated runs before drawing conclusions.

## Run the examples

Run these commands from the repository root.

### [Word count](examples/01-word-count/count_test.go)

Count whitespace-separated words. Run the named cases for empty input, mixed whitespace and punctuation, then benchmark that same function.

[count.go](examples/01-word-count/count.go)

```shell
go test -v ./020-testing-and-benchmarking/examples/01-word-count
go test -run ^$ -bench BenchmarkCount -benchmem ./020-testing-and-benchmarking/examples/01-word-count
```

## Further reading

[Testing and benchmarking in the Go documentation](https://pkg.go.dev/testing)
