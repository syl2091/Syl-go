package main

import "fmt"

//指针或值作为接收者
type B struct {
	thing int
}

func (b *B) change() {
	b.thing = 1
}
func (b B) write() string { return fmt.Sprint(b) }

func main() {
	var b B
	b.change()
	b2 := new(B)
	b2.write()

}
