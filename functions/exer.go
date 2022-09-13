package main

import "fmt"

func main() {
	//fmt.Println(foo())
	//fmt.Println(bar())

	xi := []int{1, 2, 3, 4, 5}
	defer fmt.Println("defered", foo(xi...))
	fmt.Println(bar(xi))
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(i []int) int {
	sum := 0
	for _, i := range i {
		sum += i
	}
	return sum
}
