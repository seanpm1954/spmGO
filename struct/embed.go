package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {

	t := truck{
		vehicle:   vehicle{2, "blue"},
		fourWheel: false,
	}
	s := sedan{
		vehicle: vehicle{4, "white"},
		luxury:  true,
	}

	fmt.Println(t, s)
	fmt.Println(t.doors, t.color, t.fourWheel)
}
