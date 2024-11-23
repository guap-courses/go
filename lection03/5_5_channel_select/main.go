package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go writer1(ch1)
	go writer2(ch2)

	for i := 0; i < 6; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}
}

func writer1(ch chan string) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		ch <- "from channel 1"
	}
}

func writer2(ch chan string) {
	for i := 0; i < 3; i++ {
		time.Sleep(2 * time.Second)
		ch <- "from channel 2"
	}
}
