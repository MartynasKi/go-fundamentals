package main

import (
	"fmt"
	"os"
)

func main() {
	name := "reader"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Println("Hello,", name)
}
