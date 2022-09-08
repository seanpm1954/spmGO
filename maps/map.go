package main

import "fmt"

func main() {

	//m := map[string]int{
	//	"James":           42,
	//	"MIss Moneypenny": 27,
	//}
	//
	//fmt.Println(m)
	//delete(m, "James")
	//delete(m, "MIss Moneypenny")
	//fmt.Println(m)

	//	ARRAY

	//var a [5]int
	//for k, v := range a {
	//	a[k] = k + 10
	//	fmt.Printf("%d\t%d\n", v, k)
	//}
	//fmt.Printf("Map %d\t\tType %T\n", a, a)

	// slice

	//xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	//xi1 := xi[:5]
	//xi2 := xi[5:]
	//xi3 := xi[2:7]
	//xi4 := xi[1:6]
	//
	//fmt.Println(xi1)
	//fmt.Println(xi2)
	//fmt.Println(xi3)
	//fmt.Println(xi4)

	//x := []int{56, 57, 58, 59, 60}
	//xi = append(xi, 52)
	//fmt.Println(xi)

	//xi = append(xi[:3], xi[6:]...)
	//fmt.Println(xi)

	//y := make([]string, 50, 50)
	//states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	//
	//fmt.Println(len(y))
	//fmt.Println(cap(y))
	//for i := 0; i < len(states); i++ {
	//	fmt.Println(i+1, states[i])
	//}

	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	// fmt.Println(m)

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
	m["Sean"] = []string{`Taken, not stirred`, `Martinis`, `Women`}

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
	fmt.Println()
	fmt.Println()
	delete(m, "Sean")
	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}
