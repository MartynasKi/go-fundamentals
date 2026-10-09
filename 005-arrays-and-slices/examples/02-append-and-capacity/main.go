package main

import (
	"fmt"
)

func main() {
	items := []int{1, 2, 3}
	first := items[:1]
	fmt.Println(len(first), cap(first))
	first = append(first, 8)
	fmt.Println(first, items)
	limited := items[:1:1]
	limited = append(limited, 9)
	fmt.Println(limited, items)
}
