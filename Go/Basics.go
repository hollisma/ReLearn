// Basic vars and functions

package main

import (
	"fmt"
)

// Declaring variables
var (
	isDank bool   = true
	MaxInt uint64 = 1<<64 - 1
	str    string = string(54)
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

func main() {
	// vars and printing
	fmt.Println("------------Basics------------\n")

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

// This is a cheat sheet made from the Golang tour @ https://tour.golang.org
