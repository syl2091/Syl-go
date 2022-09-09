package main

import "fmt"

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
func main() {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
}
