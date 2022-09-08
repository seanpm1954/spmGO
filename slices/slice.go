package main

import "fmt"

// a SLICE allows you to group VALUES of he same TYPE
func main() {
	// x:= type{values} COMPOSITE LITERAL
	x := make([]int, 10, 100)
	fmt.Println(x)
	//fmt.Println(x)
	//for _, v := range x {
	//	fmt.Println(v)
	//}

	//	SLICE a SLICE
	//fmt.Println(x[2])
	//fmt.Println(x)
	//fmt.Println(x[1:])  // position 1 to end
	//fmt.Println(x[1:3]) // position 1 up to but NOT including 3 (last arg in 1:3)
	//fmt.Println()
	//for i := 0; i < len(x); i++ {
	//	fmt.Println(x[i])
	//}

	//	APPEND a SLICE
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{123, 456, 987}
	fmt.Println(y)
	x = append(x, y...)

	fmt.Println(x)

	//	DELETE from a SLICE
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
