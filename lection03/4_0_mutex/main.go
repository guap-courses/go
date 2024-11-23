package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var mutex sync.Mutex

func increment() {

	for i := 0; i < 5000; i++ {

		mutex.Lock()
		counter++
		mutex.Unlock()

	}

}

func main() {
	go increment()
	go increment()

	time.Sleep(200 * time.Millisecond)
	fmt.Println("Final counter value:", counter)

}
