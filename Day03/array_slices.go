package main

import "fmt"

func main() {
	var arr1 [6]int
	var slices []int = arr1[2:5]
	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slices); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slices[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slices))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slices))

	// grow the slice
	slices = slices[0:4]
	for i := 0; i < len(slices); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slices[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slices))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slices))

	// grow the slice beyond capacity
	//slice1 = slice1[0:7 ] // panic: runtime error: slice bound out of range
}
