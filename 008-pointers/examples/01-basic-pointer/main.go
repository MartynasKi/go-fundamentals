package main

import (
	"fmt"
)

func main() {
	stock := 5
	remaining := &stock
	fmt.Println(*remaining)
	*remaining = 4
	fmt.Println(stock)
}
