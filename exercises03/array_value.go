package main

import "fmt"

//数组内存拷贝
func main() {
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i * 2
	}

	arr2 := arr
	arr[0] = 10
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d is %d\n", i, arr[i])
	}

	fmt.Println()

	for i := 0; i < len(arr2); i++ {
		fmt.Printf("%d is %d\n", i, arr2[i])
	}

}
