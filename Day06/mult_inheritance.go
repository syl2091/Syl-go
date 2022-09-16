package main

import "fmt"

//多重继承
type Camera struct {
}

func (c *Camera) TakeAPicture() string {
	return "Click"
}

type Phone struct{}

func (p *Phone) Call() string {

	return "ding ding"
}

type CameraPhone struct {
	Camera
	Phone
}

func main() {
	c := new(CameraPhone)
	fmt.Println(c.Call())
	fmt.Println(c.TakeAPicture())
}
