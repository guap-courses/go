package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	City     string
	password string // приватное поле
}

func NewPerson(name string, age int, city string, password string) *Person {
	if age < 0 {
		println("age must be greater than zero")
	}

	return &Person{
		Name:     name,
		Age:      age,
		City:     city,
		password: password,
	}
}

func (p Person) PrintName() {
	fmt.Println("Name", p.Name)
}

func (p Person) SetPassword(password string) {
	p.password = password
}
