package main

import (
	"fmt"
)

func main() {
	days := 0
	switch {
	case days == 0:
		fmt.Println("today")
	case days <= 3:
		fmt.Println("this week")
	default:
		fmt.Println("later")
	}
}
