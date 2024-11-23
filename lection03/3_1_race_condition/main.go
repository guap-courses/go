package main

import (
	"fmt"
	"time"
)

var counter int

func increment() {
	for i := 0; i < 5000; i++ {
		counter++ // Потенциально опасная операция
	}
}

func main() {
	//runtime.GOMAXPROCS(1)

	go increment()
	go increment()

	time.Sleep(200 * time.Millisecond)
	fmt.Println("Final counter value:", counter)
}
