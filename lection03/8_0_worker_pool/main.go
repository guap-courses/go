package main

import (
	"fmt"
	"sync"
)

const numJobs = 5

func getData(n int) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func worker(wg *sync.WaitGroup, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- j * 10
	}
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	inChan := getData(100)

	resultsChan := make(chan int, numJobs)

	for w := 0; w <= numJobs; w++ {
		wg.Add(1)
		go worker(wg, inChan, resultsChan)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	for res := range resultsChan {
		fmt.Println(res)
	}
}
