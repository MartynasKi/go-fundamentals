package main

import (
	"fmt"
)

func main() {
	inbox := make(chan string, 1)
	inbox <- "receipt"
	for range 2 {
		select {
		case message := <-inbox:
			fmt.Println(message)
		default:
			fmt.Println("empty")
		}
	}
}
