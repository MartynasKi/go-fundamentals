package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var workers sync.WaitGroup
	total := 0
	for range 2 {
		workers.Go(func() {
			for range 100 {
				mu.Lock()
				total++
				mu.Unlock()
			}
		})
	}
	workers.Wait()
	fmt.Println(total)
}
