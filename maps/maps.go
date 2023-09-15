package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := make(map[string]Vertex)

	m["India"] = Vertex{1.5, 3.2}

	fmt.Println(m)
	fmt.Println(m["India"])

	atlas := map[string]Vertex{
		"Bell Labs":  Vertex{5, 4},
		"Apple Corp": Vertex{15, 24},
	}

	fmt.Println(atlas)

	value, present := atlas["Zambia"]
	fmt.Println(value, present)

	delete(atlas, "Bell Labs")
	fmt.Println(atlas)
}
