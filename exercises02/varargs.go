package main

import "fmt"

// 可变参数

func main() {
	println(1, 2, 3, 4, 5)
}

func println(list ...int) {
	for _, i := range list {
		fmt.Printf("%d\n", i)
	}
}
