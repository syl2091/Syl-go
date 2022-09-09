package main

import "fmt"

func main() {
	s := "good bye"
	p := &s
	*p = "ciao"
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(s)
}
