package main

import (
	"fmt"
	"io"
	"os"
)

type PersonRepository interface {
	SavePerson(p Person)
}

type DatabaseRepository struct{}

func (m DatabaseRepository) SavePerson(p Person) {
	fmt.Println("person with name ", p.Name, " is saved to the database")
}

type CacheRepository struct{}

func (m CacheRepository) SavePerson(p Person) {
	fmt.Println("person with name ", p.Name, " is saved to the database")
}

type MockRepository struct{}

func (m MockRepository) SavePerson(p Person) {
	fmt.Println("do nothing")
}

type Service struct {
	r PersonRepository
}

func NewService(repository PersonRepository) Service {
	return Service{
		r: repository,
	}
}

func main() {

	dbRepo := DatabaseRepository{}
	_ = NewService(dbRepo)

	cacheRepo := CacheRepository{}
	_ = NewService(cacheRepo)

	mockRepo := MockRepository{}
	_ = NewService(mockRepo)

}

func compose() {
	var (
		_ io.ReadWriter
		_ io.WriteCloser
	)
}

func returningErrorFunc() error {
	return nil
}

func commonGotcha() bool {
	foo := func() error {
		var err *os.PathError // nil
		//err == nil // true
		return err
	}

	err := foo()
	if err != nil { // error != nil
		return false
	}

	return true
}
