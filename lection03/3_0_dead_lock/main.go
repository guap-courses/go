package main

import (
	"fmt"
	"sync"
	"time"
)

func someWork(id string, wg sync.WaitGroup) {
	fmt.Println(id, ": начало работы", time.Now())

	// Имитация работы
	time.Sleep(time.Second)

	fmt.Println(id, ": завершение работы", time.Now())

	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go someWork("Work 1", wg)
	wg.Add(1)
	go someWork("Work 2", wg)

	wg.Wait()
}
