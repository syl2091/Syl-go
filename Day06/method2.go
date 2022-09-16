package main

import "fmt"

type IntVector []int

func main() {
	fmt.Println(IntVector{1, 2, 3}.Sum())
}

func (v IntVector) Sum() (s int) {
	for _, value := range v {
		s += value
	}
	return s
}
