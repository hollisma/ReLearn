package main

import (
	"fmt"
	"math"
)

func main() {	
	// loops and conditionals
	fmt.Println("---------Flow Control---------")
	// flowControl()
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