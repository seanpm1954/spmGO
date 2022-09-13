package main

import "fmt"

func main() {
	foo()
	func() {
		fmt.Println("anon")
	}()

	func(x int) {
		fmt.Println("anon", x)
	}(42)
}
func foo() {
	fmt.Println("foo ran")
}
