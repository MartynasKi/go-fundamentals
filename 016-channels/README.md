# 016: Channels

Pass values between goroutines and signal the end of a stream.

## When a send can finish

On an unbuffered channel a send waits for a receiver. A buffered channel can hold values until its buffer fills. Closing a channel lets receivers drain what is left. A range loop exits once the channel is closed and empty. Sending after close panics, so give the producer responsibility for closing it.

## Run the examples

Run these commands from the repository root.

### [Stream prices](examples/01-stream-prices/main.go)

Read prices until the producer closes the channel, then print the sum.

```shell
go run ./016-channels/examples/01-stream-prices
```

### [Drain buffer](examples/02-drain-buffer/main.go)

Receive twice from a closed channel holding one item. The second result distinguishes a buffered zero from the end.

```shell
go run ./016-channels/examples/02-drain-buffer
```

## Further reading

[Channels in the Go documentation](https://go.dev/tour/concurrency/4)
