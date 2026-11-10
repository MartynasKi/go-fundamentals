package main

import (
	"fmt"
	"slices"
)

func main() {
	labels := []string{"go", "draft", "notes"}
	published := slices.Delete(slices.Clone(labels), 1, 2)
	fmt.Println(labels)
	fmt.Println(published)
}
