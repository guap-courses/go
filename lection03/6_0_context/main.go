package main

import (
	"context"
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
	ctx, cancel := context.WithCancel(context.Background())

	go writer(ch)

	go func() {
		for {
			select {
			case msg := <-ch:
				fmt.Println(msg)
			case _ = <-ctx.Done():
				fmt.Println("got signal from done chan. exiting")
				return
			}
		}
	}()

	time.Sleep(3 * time.Second)

	cancel()

	time.Sleep(100 * time.Millisecond)
}
