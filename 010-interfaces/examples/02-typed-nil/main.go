package main

import (
	"fmt"
)

type ParseError struct{}

func (*ParseError) Error() string { return "invalid input" }
func main() {
	var detail *ParseError
	var err error = detail
	fmt.Println(detail == nil, err == nil)
	err = nil
	fmt.Println(err == nil)
}
