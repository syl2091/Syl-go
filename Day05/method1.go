package main

import "fmt"

type twoInts struct {
	a int
	b int
}

func (t *twoInts) AddThem() int {
	return t.a + t.b
}

func (t *twoInts) AddParam(param int) int {
	return t.a + t.b + param
}

func main() {
	t := twoInts{1, 2}
	fmt.Println(t.AddThem())
}
