package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	City     string
	password string // приватное поле
}

func (p Person) PrintName() {
	fmt.Println("Name", p.Name)
}

func (p Person) SetPassword(password string) {
	p.password = password
}
