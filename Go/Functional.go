package main

import (
	"fmt"
	"math"
)

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

func main() {
	fmt.Println("----------Functional----------\n")

	// Can store functions as vars
	fmt.Println("---------Func as var----------")
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	fmt.Println("-----------Closure------------")
	// Closure
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i), 
			neg(-2*i),
		)
	}

	fmt.Println("----------Fibonacci-----------")
	// Cool way to do fibonacci with closures. Not good code but the way it works is interesting
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}