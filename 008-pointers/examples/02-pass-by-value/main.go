package main

import (
	"fmt"
)

func reset(p *int) {
	*p = 0
	p = nil
	fmt.Println(p == nil)
}
func main() {
	count := 6
	p := &count
	reset(p)
	fmt.Println(p == nil, count)
}
