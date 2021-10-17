package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		arr int
		mu  sync.Mutex
		wg  sync.WaitGroup
	)

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			arr += 1
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Println(arr)

}
