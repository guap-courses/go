package main

import (
	"fmt"
)

func writer() <-chan string {
	ch := make(chan string)

	go func() {
		for i := 0; i <= 9; i++ {
			ch <- fmt.Sprintf("Hello #%d", i)
		}

		close(ch)
	}()

	return ch
}

func main() {

	ch := writer()

	for v := range ch {
		fmt.Println(v)
	}
}
