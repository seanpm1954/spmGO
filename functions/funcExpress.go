package main

import "fmt"

func main() {
	test(f)
}

func test(f func()) {
	f()
}

var f = func() {
	fmt.Println("the meaning of life is")
}
