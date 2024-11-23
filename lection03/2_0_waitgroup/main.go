package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	someWork := func(id string) {
		fmt.Println(id, ": начало работы", time.Now())

		// Имитация работы
		time.Sleep(time.Second)

		fmt.Println(id, ": завершение работы", time.Now())
		wg.Done()
	}

	wg.Add(1)
	go someWork("Work 1")
	wg.Add(1)
	go someWork("Work 2")

	wg.Wait()
}
