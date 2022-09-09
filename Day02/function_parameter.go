package main

import "fmt"

//函数作为参数

func add(a, b int) {
	fmt.Printf("the sum of %d and %d is %d", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}

func main() {
	callback(1, add)
}
