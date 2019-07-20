package main

import (
	"fmt"
	"math"
)

// Stuff for MyFloat
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Stuff for Vertex
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Interface
type Abser interface {
	Abs() float64
}
// Types implement interfaces implicitly. No need to specify
// Under the hood, interfaces map values to types. 
// The value is the instance and the type is which method needs to be called

// An interface is a set of method signatures (just says which methods the interface has, doesn't implement them)
func main() {
	fmt.Println("----------Interface-----------\n")
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	fmt.Println("-----------Example------------")
	a = f  // works because MyFloat implements Abser
	fmt.Println(a.Abs())
	a = &v  // works because *Vertex implements Abser
	fmt.Println(a.Abs())
	// a = &v  Doesn't work

	// So that this main doesn't get too clustered
	nilMain()
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func nilMain() {
	fmt.Println("\n\n--------Interface p2----------\n")
	fmt.Println("---------Nil values-----------")
	
	var i I
	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hola!"}
	describe(i)
	i.M()
	// Having no concrete method (e.g. T.M()) is a run-time error)
	
	fmt.Println("-------Empty Interface--------")
	var i2 interface{}
	describeGeneral(i2)
	i2 = 42
	describeGeneral(i2)
	i2 = "string"
	describeGeneral(i2)

	fmt.Println("-------Type Assertions--------")
	var i3 interface{} = "hello"

	s := i3.(string)
	fmt.Println(s)
	s, ok := i3.(string)
	fmt.Println(s, ok)
	f, ok := i3.(float64)
	fmt.Println(f, ok)
	// f = i3.(float64)  // causes a panic because no 'ok' to check

	fmt.Println("--------Type switches---------")
	do(21)
	do("hello")
	do(true)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeGeneral(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}



// This is a cheat sheet made from the Golang tour @ https://tour.golang.org
