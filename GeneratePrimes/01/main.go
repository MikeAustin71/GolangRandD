package main

import (
	"fmt"
	"math"
)

func main() {
	upperLimit := uint64(5)
	primes := generatePrimes(upperLimit)
	fmt.Println("This is primes for upperlimit = ", upperLimit)
	fmt.Println("Primes: ", primes)
	result := computeSmallestDivisor(primes, upperLimit)
	fmt.Println("-----------------------------------")
	fmt.Println("result = ", result)
}

func computeSmallestDivisor(primes []uint64, upperLimit uint64) (result uint64) {

	result = 1

	for i := 0; i < len(primes); i++ {
		a := math.Floor(math.Log(float64(upperLimit)) / math.Log(float64(primes[i])))
		fmt.Println("i = ", i, "  a = ", a)
		result = result * uint64(math.Pow(float64(primes[i]), float64(a)))
	}

	return result
}

func generatePrimes(upperLimit uint64) []uint64 {
	var primes []uint64
	var isPrime bool
	var j int
	primes = append(primes, 2)

	for i := uint64(3); i <= upperLimit; i += 2 {
		j = 0
		isPrime = true
		for y := primes[j] * primes[j]; y <= i && j < len(primes); j++ {
			if i%primes[j] == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}
