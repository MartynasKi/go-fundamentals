package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	_, err := strconv.Atoi("four")
	var detail *strconv.NumError
	if errors.As(err, &detail) {
		fmt.Println(detail.Func, detail.Num)
	}
}
