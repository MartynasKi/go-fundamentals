package main

import (
	"fmt"
)

func main() {
	counts := make(chan int, 1)
	counts <- 0
	close(counts)
	for range 2 {
		count, ok := <-counts
		fmt.Println(count, ok)
	}
}
