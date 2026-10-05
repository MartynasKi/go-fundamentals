package main

import (
	"fmt"
)

func main() {
	for _, name := range []string{"Ada", "", "Lin"} {
		if name == "" {
			continue
		}
		fmt.Println(name)
	}
}
