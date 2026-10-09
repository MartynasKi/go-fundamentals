package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	stock := map[string]int{"tea": 2, "coffee": 3}
	for _, item := range slices.Sorted(maps.Keys(stock)) {
		fmt.Println(item, stock[item])
	}
}
