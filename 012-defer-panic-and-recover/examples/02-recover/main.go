package main

import (
	"fmt"
)

func process() {
	defer func() {
		if cause := recover(); cause != nil {
			fmt.Println("recovered:", cause)
		}
	}()
	panic("bad record")
}
func main() {
	process()
	fmt.Println("back in main")
}
