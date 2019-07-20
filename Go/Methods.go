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
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
// Can only have receivers on types defined in this package (file)

// No receiver
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// If you want to modify receiver(struct) then use pointer receiver
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

// Scale as regular function
func Scale(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}
// Advantage of pointer receiver is that you can modify value and so you don't need to include it with each method call

func main() {
	fmt.Println("----------Functional----------\n")
	v := Vertex{3, 4}
	fmt.Println("---------Receiver fun---------")
	fmt.Println(v.Abs())
	fmt.Println("---------Regular func---------")
	fmt.Println(Abs(v))
	fmt.Println("-------Pointer receiver-------")
	v.Scale(10)
	fmt.Println(v.Abs())
	(&v).Scale(2)  // Also works
	fmt.Println(v.Abs())
	Scale(&v, 5)  // Scale(v, 5) won't work
	fmt.Println(v.Abs())
}

// This is a cheat sheet made from the Golang tour @ https://tour.golang.org
