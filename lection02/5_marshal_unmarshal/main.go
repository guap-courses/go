package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type config struct {
	Name   string
	Active bool
	Age    uint
}

func main() {

	//fileJson()

	//jsonTag()

	customMarshal()
}

func fileJson() {
	type User struct {
		Name   string
		Active bool
		Age    uint
	}

	data, _ := json.Marshal(User{Name: "Vasya", Active: false, Age: 30})
	fmt.Println(string(data))

	var u User
	_ = json.Unmarshal(data, &u)
	fmt.Println(u)

}
func jsonTag() {
	type UserJ struct {
		Name string `json:"name"`

		Password string `json:"-"`

		Registered bool  `json:"registered,omitempty"`
		Active     *bool `json:"active,omitempty"`

		age uint
	}

	f := false
	data, _ := json.Marshal(UserJ{
		Name:     "Vasya",
		Password: "admin",

		Registered: false,
		Active:     &f,
		age:        30,
	})
	fmt.Println(string(data))
}

type Date struct {
	time.Time
}

var dFormat = "2006-01-02"

func (d Date) MarshalJSON() ([]byte, error) {
	s := d.Format(dFormat)
	return []byte(fmt.Sprintf(`"%s"`, s)), nil
}

//func (d *Date) UnmarshalJSON(data []byte) error {
//	s := bytes.Trim(data, "\"")
//	t, err := time.Parse(dFormat, string(s))
//	if err != nil {
//		return err
//	}
//	d.Time = t
//	return nil
//}

func customMarshal() {
	type DateTime struct {
		D Date      `json:"date"`
		T time.Time `json:"time"`
	}

	now := time.Now()
	data, _ := json.Marshal(DateTime{
		D: Date{Time: now},
		T: now,
	})

	fmt.Println(string(data))

	var dt DateTime
	_ = json.Unmarshal(data, &dt)

	fmt.Printf("%+v\n", dt)
}
