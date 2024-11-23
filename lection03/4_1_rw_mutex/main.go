package main

import (
	"fmt"
	"sync"
)

var s string

func main() {
	mu := new(sync.RWMutex)

	go func() {
		for {
			mu.RLock()
			fmt.Println(s)
			mu.RUnlock()
		}
	}()

	for i := 0; i < 100000; i++ {
		mu.Lock()
		if i%2 == 0 {
			s = "123"
		} else {
			s = "ABCDE"
		}
		mu.Unlock()
	}
}

// ABCDE
// 12312
// ABCDE
// ABCDE
// ABC
// ABCDE
// ABCDE
// 123
// 12312
// 123
// 123
// 123
// 123
