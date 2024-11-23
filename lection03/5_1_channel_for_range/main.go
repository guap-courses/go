package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {
		for i := 0; i <= 9; i++ {
			ch <- fmt.Sprintf("Hello #%d", i)
		}

		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("channel closed")
}
