package main

import "fmt"

type person struct {
	first   string
	last    string
	flavors []string
}

func main() {
	p1 := person{
		first: "sean",
		last:  "maloney",
		flavors: []string{"choc",
			"martini",
			"rum and coke",
		},
	}

	p2 := person{
		first: "james",
		last:  "Bond",
		flavors: []string{"vanilla",
			"ltk",
			"black",
		},
	}

	fmt.Println(p1.first, p1.last)
	for _, v := range p1.flavors {
		fmt.Println(v)
	}
	fmt.Println(p2.first, p2.last)
	for _, v := range p2.flavors {
		fmt.Println(v)
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	fmt.Println(m)

	for i, v := range m {
		fmt.Println(i)
		for _, y := range v.flavors {
			fmt.Println(y)
		}
	}

}
