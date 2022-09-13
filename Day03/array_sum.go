package main

import "fmt"

//数组传递给函数  传递数组指针
func main() {
	array := [3]float64{7.0, 8.5, 9.1}
	x := sum(&array) // Note the explicit address-of operator
	// to pass a pointer to the array
	fmt.Printf("The sum of the array is: %f", x)
}

func sum(a *[3]float64) (sum float64) {
	for _, f2 := range a {
		sum += f2
	}
	return
}
