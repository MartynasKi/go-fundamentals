package main

import (
	"fmt"
)

func main() {
	step := "opening"
	defer fmt.Println("argument:", step)
	defer func() { fmt.Println("closure:", step) }()
	step = "finished"
}
