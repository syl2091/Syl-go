package main

import "fmt"

var num int = 10
var numx2, numx3 int

func main() {
	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()
}

func getX2AndX3_2(n int) (x2 int, x3 int) {
	x2 = 2 * n
	x3 = 3 * n
	// return x2, x3
	return
}

func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

func getX2AndX3(n int) (int, int) {
	return 2 * n, 3 * n
}
