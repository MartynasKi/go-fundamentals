package main

import (
	"fmt"
)

type Article struct {
	Title string
	Tags  []string
}

func main() {
	draft := Article{Title: "Go notes", Tags: []string{"draft"}}
	copyOfDraft := draft
	copyOfDraft.Title = "Slice notes"
	copyOfDraft.Tags[0] = "go"
	fmt.Println(draft)
	fmt.Println(copyOfDraft)
}
