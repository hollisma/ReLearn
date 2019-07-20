package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("---------Flow Control---------\n")

	// For loop
	fmt.Println("-----------For loop-----------")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// While loop
	fmt.Println("----------While loop----------")
	for sum < 900 {
		sum += sum
	}
	fmt.Printf("%d\n", sum)

	// Infitinte loop
	// for {}

	fmt.Println("--------------If--------------")
	if 1 < 2 {
		fmt.Println("1 < 2")
	}

	// Can execute short statement. Only in scope until end of if
	fmt.Println("--------Short statement-------")
	if v := 1; v < 2 {
		fmt.Println("v < 2")
	} else {
		fmt.Println("v > 2")
	}

	// Stops when a case succeeds
	fmt.Println("------------Switch------------")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default: 
		fmt.Printf("%s. \n", os)
	}

	// defer runs the functions after the surrounding function has executed
	fmt.Println("-------------Defer------------")
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Printf("in defer: %d\n", i)
	}
	fmt.Println("done")
}

// This is a cheat sheet made from the Golang tour @ https://tour.golang.org
