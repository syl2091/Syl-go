package main

import (
	"fmt"
	"strconv"
)

//string()方法

type TwoInts struct {
	a, b int
}

func main() {
	two1 := new(TwoInts)
	two1.a = 10
	two1.b = 12
	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1)
	fmt.Printf("two1 is: %#v\n", two1)
}

func (tn *TwoInts) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}
