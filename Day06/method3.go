package main

import (
	"fmt"
	"math"
)

//内嵌类型的方法和继承
type Point struct {
	x, y float64
}

func (p *Point) ads() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamePoint struct {
	Point
	name string
}

func main() {
	n := &NamePoint{Point{3, 4}, "lege"}
	fmt.Println(n.ads())
}
