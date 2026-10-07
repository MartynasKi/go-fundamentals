package main

import (
	"fmt"
	"strings"
)

func main() {
	for _, label := range []string{"env:dev", "draft"} {
		key, value, found := strings.Cut(label, ":")
		fmt.Printf("%q %q %t\n", key, value, found)
	}
}
