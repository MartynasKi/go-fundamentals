package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	names := []string{"ada", "lin"}
	labels := make([]string, len(names))
	var workers sync.WaitGroup
	for i, name := range names {
		workers.Go(func() { labels[i] = strings.ToUpper(name) })
	}
	workers.Wait()
	fmt.Println(labels)
}
