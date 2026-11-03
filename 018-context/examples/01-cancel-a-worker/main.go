package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		defer close(done)
		<-ctx.Done()
		fmt.Println(ctx.Err())
	}()
	cancel()
	<-done
}
