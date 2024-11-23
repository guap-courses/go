package main

import (
	"fmt"
)

func main() {
	bufferedChan := make(chan string, 10)

	for i := 0; i <= 10; i++ {
		bufferedChan <- fmt.Sprintf("Hello #%d", i)
	}
	close(bufferedChan)

	for v := range bufferedChan {
		fmt.Println(v)
	}

	fmt.Println("channel closed")
}
