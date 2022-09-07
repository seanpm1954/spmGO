package main

import "fmt"

func main() {
	s := "Hello World"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}

	for i, v := range s {
		fmt.Println(i, v)
	}

	x := 2
	fmt.Printf("%d\t\t%b\n\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n\n\n", y, y)

	const (
		_  = iota
		kb = 1 << (iota * 10)
		mb = 1 << (iota * 10)
		gb = 1 << (iota * 10)
	)
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n\n", gb, gb)

	nkb := kb << 10
	fmt.Printf("%d\t\t%b\n", nkb, nkb)

}
