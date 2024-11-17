package main

import (
	"fmt"
	"time"
)

func main() {

	//initialization()

	//emptyStruct()

	//embedding()

	//set()
	//
	//methods()
	//
	//embedding()
	//
	//pointerValue()
}

//type ИмяСтруктуры struct {
//	Поле1 Тип
//	Поле2 Тип
//}

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

func (p *Person) SetPassword(password string) {
	p.password = password
}

func initialization() {

	c1 := Person{Name: "Vasja", Age: 12, City: "SPB", password: "123456"}
	fmt.Println(c1)

	c2 := Person{Name: "Vasja", Age: 12}
	fmt.Println(c2)

	c3 := Person{}
	fmt.Println(c3)

	c4 := Person{"Vasja", 12, "SPB", "123456"}
	fmt.Println(c4)

	a := NewPerson("Vasja", 12, "SPB", "123456")

	fmt.Println(a)
}

func emptyStruct() {

	var _ struct{}

	_ = struct{}{}

	elements := make(map[string]struct{})
	elements["element"] = struct{}{}

	if _, exists := elements["element"]; exists {
		fmt.Println("Element exists in the elements")
	}

}

type PersonWithBirthday struct {
	Person

	Birthday time.Time
}

func embedding() {

	person := PersonWithBirthday{
		Person:   Person{Name: "Vasja"},
		Birthday: time.Now(),
	}

	person.Person.PrintName()

	person.SetPassword("123456")

	fmt.Println(person)
}
