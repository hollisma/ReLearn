package main

import (
	"fmt"
	"math"
)

func main() {	
	// loops and conditionals
	fmt.Println("---------Flow Control---------")
	// flowControl()
	
	fmt.Println("------Pointers & Structs------")
	// pointersAndStructs()
	
	fmt.Println("------------Arrays------------")
	// arrays()
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