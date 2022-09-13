package main

import (
	"fmt"
)

func main() {
	//ii := []int{1, 2, 3, 4}
	//fmt.Println(sum(ii...))

	//xx := even(sum, ii...)
	//fmt.Println(xx)

	//xx1 := odd(sum, ii...)
	//fmt.Println(xx1)

	fmt.Println(fact(4))

}

func sum(xi ...int) int {
	i := 0
	for _, v := range xi {
		i += v
	}
	return i
}

func fact(x int) int {
	total := 1
	for ; x > 0; x-- {
		total *= x
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)
}

func bar() func() int {
	return func() int {
		return 42
	}
}
