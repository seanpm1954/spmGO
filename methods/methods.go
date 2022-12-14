package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}
func (p person) speak() {
	fmt.Println("No License to kill")
}

func main() {

	sa1 := secretAgent{
		person: person{"James", "Bond"},
		ltk:    true,
	}
	p := person{
		first: "Sean",
		last:  "Maloney",
	}
	fmt.Println(sa1.first, sa1.last, sa1.ltk)
	sa1.speak()
	p.speak()
}
