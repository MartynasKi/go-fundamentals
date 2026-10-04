package main

import (
	"fmt"
)

type Celsius float64

func main() {
	reading := Celsius(21.9)
	fmt.Println(reading, int(reading))
}
