package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	word := "café"
	for offset, letter := range word {
		fmt.Printf("%d %c\n", offset, letter)
	}
	fmt.Println(len(word), utf8.RuneCountInString(word))
}
