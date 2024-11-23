package main

import (
	"fmt"
	"time"
)

func someWork(id string) {
	fmt.Println(id, ": начало работы", time.Now())

	// Имитация работы
	time.Sleep(time.Second)

	fmt.Println(id, ": завершение работы", time.Now())
}

func main() {
	go someWork("Work 1")
	go someWork("Work 2")
	go someWork("Work 3")

	// Задержка, чтобы основной поток не завершился сразу
	//time.Sleep(5 * time.Second)
}
