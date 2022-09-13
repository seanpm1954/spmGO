package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func main() {
	sq := square{15}
	sq.area()
	//fmt.Println("square area", sq)

	ci := circle{12.345}
	ci.area()
	//fmt.Println("circle area", ci)

	info(sq)
	info(ci)
	//anon func and assign func
	f := func(x int) {
		fmt.Println("anon", x)
	}
	f(42)

	//return func
	rf := rFunc()
	fmt.Println(rf())

}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {

	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println(s.area())
}

func rFunc() func() int {
	return func() int {
		return 42
	}
}
