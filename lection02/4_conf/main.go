package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	//userInput()

	//env()
	//
	//flags()

	//filesOpen()

	filesCreate()
}

func userInput() {

	fmt.Println("Enter login:")
	var login string
	_, _ = fmt.Scan(&login) // same as fmt.Fscan(os.Stdin, &login)

	fmt.Println("Enter animal and random integer:")
	var (
		animal string
		number int
	)
	_, _ = fmt.Scanf("%s %d", &animal, &number)

	fmt.Printf("user %s likes %s and his number is %d", login, animal, number)
}

func env() {

	var environment = "PROD"

	s, ok := os.LookupEnv(environment)
	fmt.Println(s, ok)

	s = os.Getenv(environment)
	fmt.Println(s)

	//all := os.Environ()
	//fmt.Println(all)

}

func flags() {

	var verbose bool
	flag.BoolVar(&verbose, "v", false, "add expanded logs")

	var configFilePath = flag.String("config", "", "path to config file")

	flag.Parse()

	if verbose {
		fmt.Println("verbose is enabled")
	}

	if configFilePath != nil {
		fmt.Printf("path to file: %s\n", *configFilePath)
	}
}

func filesOpen() {

	_, err := os.Open("input.txt")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("file is not exists")
	}

	f2, _ := os.Open("lection02/4_conf/input2.txt")
	data, _ := io.ReadAll(f2)
	defer f2.Close()

	fmt.Println(string(data))

}

func filesCreate() {

	f, _ := os.Create("file.txt") // create or truncate
	defer f.Close()

	_, _ = f.WriteString("Текст, который нужно записать\n")

}
