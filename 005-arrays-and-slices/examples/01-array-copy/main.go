package main

import (
	"fmt"
)

func main() {
	original := [3]int{10, 20, 30}
	copied := original
	copied[0] = 99
	view := original[1:]
	view[0] = 42
	fmt.Println(original)
	fmt.Println(copied)
}
