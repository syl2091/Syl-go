package main

import "fmt"

//指针方法和值方法都可以在指针或非指针上被调用
type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) append(val int) {
	*l = append(*l, val)
}
func main() {
	//值
	var lst List
	lst.append(1)
	fmt.Printf("%v (len: %d)", lst, lst.Len())

	//指针
	plst := new(List)
	plst.append(2)
	fmt.Printf("%v (len: %d)", plst, plst.Len())

}
