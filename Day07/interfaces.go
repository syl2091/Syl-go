package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (s *Square) Area() float32 {
	return s.side * s.side
}

func main() {
	s := new(Square)
	s.side = 5
	var s2 Shaper
	s2 = s
	fmt.Printf("The square has area: %f\n", s2.Area())
}
