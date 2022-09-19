package main

import (
	"fmt"
	"reflect"
)

//通过反射修改值

func main() {
	f := 3.4
	v := reflect.ValueOf(f)
	fmt.Println("settability of v:", v.CanSet())
	v = reflect.ValueOf(&f) // Note: take the address of x.
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())
	v = v.Elem()
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(3.1415) // this works!
	fmt.Println(v.Interface())
	fmt.Println(v)
}
