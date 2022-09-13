package main

import "fmt"

func main() {
	ints := make([]int, 4)
	for i, value := range ints {
		fmt.Printf("%d is %d", i, value)
	}
}
