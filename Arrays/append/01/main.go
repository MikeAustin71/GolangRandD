package main

import (
	"fmt"
)

func main() {
	primes := []uint64{2}
	primes = append(primes, 3)
	fmt.Println("This is primes array: ", primes)

	var empty []uint64
	empty = append(empty, 2)
	fmt.Println("This is array empty: ", empty)
}
