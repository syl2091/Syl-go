package main

import "fmt"

const c = "C"

var v int = 5

type T struct {
}

func init() {

}

func Func1() {

}

func main() {
	//var a int
	var t T
	Func1()
	Method1(t)
	// ...
	//fmt.Println(a)
	fmt.Println(a, b, d)
}

func Method1(t T) {
	fmt.Println(t)
}

const (
	a = iota
	b
	d
)
