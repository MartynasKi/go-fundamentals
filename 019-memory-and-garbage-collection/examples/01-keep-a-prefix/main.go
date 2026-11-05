package main

import (
	"bytes"
	"fmt"
)

func main() {
	buffer := []byte("draft: reusable input")
	view := buffer[:5]
	saved := bytes.Clone(view)
	copy(buffer, "final")
	fmt.Printf("%s / %s\n", view, saved)
}
