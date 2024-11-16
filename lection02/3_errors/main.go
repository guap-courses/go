package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль") // возвращаем ошибку
	}
	return a / b, nil // ошибок нет
}

func errorHandling() {

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Результат:", result)

}

type custError struct {
	ErrorCode int
	ErrorMsg  string
	ErrorMsg2 string
}

func (e custError) Error() string {
	return e.ErrorMsg
}

var customErr error = custError{}

func errAsIs() {

	var ErrNotFound = custError{ErrorMsg: "not found", ErrorCode: 404}

	f := func() error {
		return fmt.Errorf("opening file error: %w", ErrNotFound)
	}

	err := f()

	_, ok := err.(custError)
	fmt.Println(ok) // ok = false

	ok = errors.Is(err, ErrNotFound)
	fmt.Println(ok) // ok = true

	ok = errors.Is(err, ErrNotFound)
	fmt.Println(ok)

	var ce custError
	ok = errors.As(err, &ce)
	fmt.Printf("%t, %v, %v", ok, ce.ErrorMsg, ce.ErrorCode)

}

func main() {
	//errorHandling()

	errAsIs()
}
