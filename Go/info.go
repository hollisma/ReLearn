package main

import (
	"fmt"
	"math"
)

// func name(var1, var2 type) return_type {}
func add(x, y int) int {
	return x + y
}

// func name(var1 type) (return1, return2 type) {}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// "naked" return. just returns specified stuff
	return
}

var (
	isDank bool   = true
	MaxInt uint64 = 1<<64 - 1
	str    string = string(54)
)

func main() {
	// vars and printing
	fmt.Println("------------Basics------------")
	// basics()
	
	// loops and conditionals
	fmt.Println("---------Flow Control---------")
	// flowControl()
	
	fmt.Println("------Pointers & Structs------")
	// pointersAndStructs()
	
	fmt.Println("------------Arrays------------")
	// arrays()
	
	fmt.Println("-------------Maps-------------")
	// maps()
	
	fmt.Println("----------Functional----------")
	functional()
	
	fmt.Println("------------Methods-----------")
	methods()	
}

func basics() {
	// Each method from a package starts with a capital char
	fmt.Println(isDank, MaxInt, str)
	fmt.Println("Hola!")

	// var var1, var2 type = stuff
	var a, b int = 1, 2
	fmt.Println(add(a, b))

	// equivalent to var statement
	c := 17
	fmt.Println(split(c))
}

func flowControl() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// While loop
	for sum < 900 {
		sum += sum
	}
	fmt.Printf("%g", sum)

	// Infitinte loop
	// for {}

	if 1 < 2 {
		fmt.Println("1 < 2")
	}

	// Can execute short statement. Only in scope until end of if
	if v := 1; v < 2 {
		fmt.Println("v < 2")
	} else {
		fmt.Println("v > 2")
	}

	// Stops when a case succeeds
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default: 
		fmt.Printf("%s. \n", os)
	}

	// defer runs the functions after the surrounding function has executed
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Printf("in defer: %g\n", i)
	}
	fmt.Println("done")
}

func pointersAndStructs() {
	// Pointers. No pointer arithmetic
	i := 42
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	// Structs
	type Vertex struct {
		X int
		Y int
	}
	v := Vertex{1, 2}
	fmt.Println(v)

	// Accessing field of struct
	v.X = 4
	fmt.Println(v.X)

	// Pointers to structs can access fields
	pV := &v
	pV.X = 5
	fmt.Println(v.X)

	// Struct Literals
	pL := &Vertex{Y: 4} // defaults X:0
	fmt.Println(*pL)
}

func arrays() {
	// Creating arrays
	// EQ to var primes [6]int
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slices
	primesLT10 := primes[:4]
	fmt.Println(primesLT10)

	// Slices merely reference
	primesLT10[3] = 1
	fmt.Println(primes)

	// Slice literal
	sliceLiteral := []bool{true, true, false}
	fmt.Println(sliceLiteral)

	// Slice length vs. capacity
	s := sliceLiteral[:1]
	fmt.Println(s)
	fmt.Println(len(s), cap(s))
	s = sliceLiteral[:2]
	fmt.Println(s)
	fmt.Println(len(s), cap(s))
	s = sliceLiteral[1:]
	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	// Nil slice = zero value
	var n []int
	fmt.Println(n, len(n), cap(n))
	
	// Slices with make
	length, capacity := 5, 8
	a := make ([]int, length, capacity)
	fmt.Println(a, len(a), cap(a))

	// Append
	a = append(a, 3)
	fmt.Println(a)

	// Range takes array and first value is index, second is element at that index
	var pow = []int{1, 2, 4, 8, 16}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	// Can skip with blank space
	// for _, v := range pow {}
}

func maps() {
	type Vertex struct {
		Lat, Long float64
	}

	// Maps are basically dictionaries in Python. [x] means keys are type x
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex {
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// Map literal
	var n = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google": {37.42202, -122.08408},
	}
	fmt.Println(n)

	delete(n, "Bell Labs")
	// Test is element is in key. ok is true if in, false if not in
	elem, ok := n["Amazon"]
	fmt.Println(elem, ok)
	elem, ok = n["Google"]
	fmt.Println(elem, ok)
}

func functional() {

	// Can store functions as vars
	// Closure
	// func compute(fn func(float64, float64) float64) float64 {
	// 	return fn(3, 4)
	// }

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// Closure
	// func adder() func(int) {
	// 	sum := 0
	// 	return func(x int) int {
	// 		sum += x
	// 		return sum
	// 	}
	// }

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i), 
			neg(-2*i),
		)
	}

	// Cool way to do fibonacci with closures. Not good code but the way it works is interesting
	// func fibonacci() func() int {
	// 	prev := 0
	// 	curr := 0
	// 	return func() int {
	// 		if prev == 0 {
	// 			prev = 1
	// 			return curr
	// 		}
	// 		curr += prev
	// 		prev = curr - prev
	// 		return curr
	// 	}
	// }

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func fibonacci() func() int {
	prev := 0
	curr := 0
	return func() int {
		if prev == 0 {
			prev = 1
			return curr
		}
		curr += prev
		prev = curr - prev
		return curr
	}
}

func methods() {
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
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))
}

// This is a cheat sheet made from the Golang tour @ https://tour.golang.org
