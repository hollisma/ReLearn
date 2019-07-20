package main

import "fmt"

func main() {
	fmt.Println("------Pointers & Structs------\n")

	// Pointers. No pointer arithmetic
	fmt.Println("-----------Pointers-----------")
	i := 42
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	// Structs
	fmt.Println("------------Structs-----------")
	type Vertex struct {
		X int
		Y int
	}
	v := Vertex{1, 2}
	fmt.Println(v)

	// Accessing field of struct
	fmt.Println("------Access struct field-----")
	v.X = 4
	fmt.Println(v.X)

	// Pointers to structs can access fields
	fmt.Println("--------Pointer access--------")
	pV := &v
	pV.X = 5
	fmt.Println(v.X)

	// Struct Literals
	fmt.Println("--------Struct Literal--------")
	pL := &Vertex{Y: 4} // defaults X:0
	fmt.Println(*pL)
}

// This is a cheat sheet made from the Golang tour @ https://tour.golang.org
