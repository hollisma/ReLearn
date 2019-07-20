package main

import (
	"fmt"
	"math"
)
	
// Go doesn't have classes, but you can define methods by adding a receiver argument to the function
type Vertex struct {
	X, Y float64
}

// Receiver arg is (v Vertex)
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y+v.Y)
}

// No receiver
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	fmt.Println("----------Functional----------\n")
	v := Vertex{3, 4}
	fmt.Println("---------Receiver fun---------")
	fmt.Println(v.Abs())
	fmt.Println("---------Regular func---------")
	fmt.Println(Abs(v))
}

// This is a cheat sheet made from the Golang tour @ https://tour.golang.org
