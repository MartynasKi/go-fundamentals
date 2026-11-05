package main

import (
	"fmt"
	"runtime"
)

func main() {
	buffer := make([]byte, 1<<20)
	runtime.GC()
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	fmt.Printf("heap bytes: %d\n", stats.HeapAlloc)
	runtime.KeepAlive(buffer)
}
