package main

import (
	"fmt"
)

func indexOf[T comparable](items []T, want T) int {
	for i, item := range items {
		if item == want {
			return i
		}
	}
	return -1
}
func main() {
	fmt.Println(indexOf([]string{"tea", "coffee"}, "coffee"))
	fmt.Println(indexOf([]int{4, 7}, 9))
}
