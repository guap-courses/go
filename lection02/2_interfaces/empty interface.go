package main

import (
	"fmt"
)

func printValue(v interface{}) {
	fmt.Println(v)
}

func emptyInterface() {
	printValue(42)      // передаем int
	printValue("Hello") // передаем string
	printValue(3.14)    // передаем float64

}

func typeAssertion() {

	var x interface{} = 24

	v, ok := x.(int)   // true, так как x - int
	fmt.Println(v, ok) // 24, true

	var y interface{} = "Hello"

	v, ok = y.(int)    // false, так как x не int
	fmt.Println(v, ok) // 0, false

	var b interface{} = "Hello"

	// assertion with type switch syntax
	switch b.(type) {
	case string:
		fmt.Println("b is string", b)
	case int:
		fmt.Println("b is Exchange", b)
	case error:
		fmt.Println("b is error", b)
	default:
		fmt.Println("b is unde", b)
	}
}

func main() {

	emptyInterface()

	typeAssertion()
}
