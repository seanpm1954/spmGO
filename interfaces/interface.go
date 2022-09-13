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
	fmt.Println("The secret agent speak")
}
func (p person) speak() {
	fmt.Println("The person speak")
}

type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I was passed into bar()", h)
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
	p1 := person{
		first: "Bob",
		last:  "Saget",
	}

	//fmt.Println(sa1.first, sa1.last, sa1.ltk)
	sa1.speak()
	p.speak()
	p1.speak()
	(bar(sa1))
	(bar(p))
	(bar(p1))

}
