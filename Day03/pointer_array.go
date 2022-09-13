package main

import "fmt"

func main() {
	var arr [3]int
	f(arr)
	fp(&arr)
}

func f(a [3]int) {
	fmt.Println(a)
}

func fp(a *[3]int) {
	fmt.Println(a)
}
