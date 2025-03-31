package main

import (
	"fmt"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	//f := factorial(20)
	n := int64(20)

	c := incrementor(n)

	cSum := puller(c)

	for x := range cSum {
		fmt.Println(n, "Factorial =", x)
	}

}

/*
func factorial(n int64) int64 {
	total := int64(1)
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
*/

func incrementor(n int64) chan int64 {
	out := make(chan int64)

	go func() {
		for i := n; i > 0; i-- {
			out <- i
		}
		close(out)

	}()

	return out
}

func puller(c chan int64) chan int64 {
	out := make(chan int64)

	go func() {
		sum := int64(1)
		for n := range c {
			sum *= n
			// fmt.Println("n = ", n, "  - sum =", sum)
		}
		out <- sum
		close(out)
	}()

	return out
}
