package main

import "fmt"

// a SLICE allows you to group VALUES of he same TYPE
func main() {
	// x:= type{values} COMPOSITE LITERAL
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	for _, num := range x {
		fmt.Println(num)
	}
}
