package main

import (
	"fmt"
	"sync"
)

func main() {
	// Solution #1: Mutex
	var mu sync.Mutex
	count := 0

	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100_000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)

}
