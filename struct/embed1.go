package main

import "fmt"

func main() {

	t := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "sean",
		friends: map[string]int{
			"bob": 555,
			"sam": 777,
		},
		favDrinks: []string{
			"Grand Marnier",
			"Jameson",
		},
	}
	fmt.Println(t.first)
	for k, v := range t.friends {
		fmt.Println("Friend ", k, v)
	}
	for _, v := range t.favDrinks {
		fmt.Println("Drinks ", v)
	}

}
