package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func qube(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n * n
		}
		close(out)
	}()
	return out
}

func reader(in <-chan int) {
	for n := range in {
		fmt.Println(n)
	}
}

func main() {
	c := gen(2, 3)

	out := sq(c)
	out2 := qube(out)

	reader(out2)
}
