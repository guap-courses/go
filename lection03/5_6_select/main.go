package main

import (
	"fmt"
	"time"
)

func writer(ch chan int) {
	for i := 0; ; i++ {
		time.Sleep(100 * time.Millisecond)
		ch <- i
	}
}

func main() {
	ch := make(chan int)
	done := make(chan struct{})

	go writer(ch)

	go func() {
		time.Sleep(3 * time.Second)
		done <- struct{}{}
	}()

	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		case _ = <-done:
			fmt.Println("got signal from done chan. exiting")
			return
		}
	}
}
