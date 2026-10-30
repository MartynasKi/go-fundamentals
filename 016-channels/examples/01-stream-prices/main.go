package main

import (
	"fmt"
)

func main() {
	prices := make(chan int)
	go func() {
		defer close(prices)
		for _, price := range []int{4, 6, 2} {
			prices <- price
		}
	}()
	total := 0
	for price := range prices {
		total += price
	}
	fmt.Println(total)
}
