package main

import (
	"errors"
	"fmt"
)

// T sldkfjsdf
type T struct {
	F int
}

func main() {
	pripriCeses, err := getMap()

	if err != nil {
		return
	} else {
		fmt.Println(pripriCeses)
	}
	// _ = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	for k, v := range pripriCeses {
		fmt.Println(k, v)
	}
	t3 := T{}
	_ = t3

	ii3 := []int{}
	_ = ii3

	m3 := map[int]int{}
	_ = m3

	t4 := &T{}
	_ = t4

	ii4 := &[]int{}
	_ = ii4

	m4 := &map[int]int{}
	_ = m4
}

func getMap() (map[string]int, error) {
	return map[string]int{
		"apple":  12,
		"banana": 34,
		"cherry": 22,
	}, errors.New("error")
}
