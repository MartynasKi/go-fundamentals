package main

import (
	"fmt"
)

func main() {
	stock := map[string]int{"tea": 0}
	for _, item := range []string{"tea", "coffee"} {
		count, exists := stock[item]
		fmt.Println(item, count, exists)
	}
}
