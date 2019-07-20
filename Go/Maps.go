package main

import "fmt"

func main() {
	fmt.Println("-------------Maps-------------\n")

	// How to create a struct. Same as in C
	type Vertex struct {
		Lat, Long float64
	}

	// Maps are basically dictionaries in Python. [x] means keys are type x
	fmt.Println("-----------Make map-----------")
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex {
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// Map literal
	fmt.Println("---------Map literal----------")
	var n = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google": {37.42202, -122.08408},
	}
	fmt.Println(n)

	fmt.Println("-------Delete, test in--------")
	delete(n, "Bell Labs")
	// Test is element is in key. ok is true if in, false if not in
	elem, ok := n["Amazon"]
	fmt.Println(elem, ok)
	elem, ok = n["Google"]
	fmt.Println(elem, ok)
}

// This is a cheat sheet made from the Golang tour @ https://tour.golang.org
