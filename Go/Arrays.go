package main

import (
	"fmt"
)

func main() {
	fmt.Println("------------Arrays------------\n")

	fmt.Println("----------Make array----------")
	// Creating arrays
	// EQ to var primes [6]int
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slices
	fmt.Println("------------Slices------------")
	primesLT10 := primes[:4]
	fmt.Println(primesLT10)

	// Slices merely reference
	fmt.Println("-----------Slice Ref----------")
	primesLT10[3] = 1
	fmt.Println(primes)

	// Slice literal
	fmt.Println("---------Slice Literal--------")
	sliceLiteral := []bool{true, true, false}
	fmt.Println(sliceLiteral)

	// Slice length vs. capacity
	fmt.Println("------Slice len and cap-------")
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
	fmt.Println("----------Nil slice-----------")
	var n []int
	fmt.Println(n, len(n), cap(n))
	
	// Slices with make
	fmt.Println("-------------Make-------------")
	length, capacity := 5, 8
	a := make ([]int, length, capacity)
	fmt.Println(a, len(a), cap(a))

	// Append
	fmt.Println("------------Append------------")
	a = append(a, 3)
	fmt.Println(a)

	// Range takes array and first value is index, second is element at that index
	fmt.Println("------------Range-------------")
	var pow = []int{1, 2, 4, 8, 16}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	// Can skip with blank space
	// for _, v := range pow {}
}

// This is a cheat sheet made from the Golang tour @ https://tour.golang.org
