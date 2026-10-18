package main

import (
	"bytes"
	"fmt"
	"io"
)

func label(w io.Writer, item string) error {
	_, err := fmt.Fprintln(w, "Item:", item)
	return err
}
func main() {
	var buf bytes.Buffer
	if err := label(&buf, "tea"); err != nil {
		panic(err)
	}
	fmt.Print(buf.String())
}
