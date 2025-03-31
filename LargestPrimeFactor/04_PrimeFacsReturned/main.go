package main

import (
	"fmt"
	"math/big"
)

// Note: The prime factors of 13195 are 5, 7, 13 and 29.
// The largest prime factor of 600851475143 is: 6857

func main() {
	//n := big.NewInt(600851475143)
	n := big.NewInt(13195)

	zero := big.NewInt(0)
	largestFac, allFactors, nonPrimes, primes := findLargestPrimeFactor(n)

	if largestFac.Cmp(zero) == 0 {
		largestFac.Set(n)
	}

	fmt.Println("-----------------------------------------------------")
	fmt.Println("The Largest Prime Factor of", n, " is", largestFac)
	fmt.Println("-----------------------------------------------------")
	fmt.Println("All Factors: ")
	fmt.Println(allFactors)
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Non-Prime Factors: ")
	fmt.Println(nonPrimes)
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Prime Factors: ")
	fmt.Println(primes)
	fmt.Println("=====================================================")
}

func findLargestPrimeFactor(n *big.Int) (*big.Int, []*big.Int, []*big.Int, []*big.Int) {

	allFacs, nonPrimes, primes := findAllFacs(n)
	largestPrimeFac := findLargestFactor(primes)

	return largestPrimeFac, allFacs, nonPrimes, primes
}

func findLargestFactor(factors []*big.Int) *big.Int {

	limit := len(factors)

	if limit == 0 {
		return big.NewInt(0)
	}

	largestFac := big.NewInt(0)

	for i := 0; i < limit; i++ {
		if factors[i].Cmp(largestFac) > 0 {
			largestFac.Set(factors[i])
		}
	}

	return largestFac
}

func findAllFacs(n *big.Int) (allFacs []*big.Int, primeFacs []*big.Int, nonPrimeFacs []*big.Int) {

	allFacs = make([]*big.Int, 0, 200)
	primeFacs = make([]*big.Int, 0, 50)
	nonPrimeFacs = make([]*big.Int, 0, 100)

	temp := big.NewInt(1)
	increment := big.NewInt(1)
	quotient := big.NewInt(0)
	r := big.NewInt(0)
	mod := big.NewInt(0)
	zero := big.NewInt(0)

	for i := big.NewInt(2); temp.Mul(i, i).Cmp(n) < 0; i.Add(i, increment) {

		quotient, mod = temp.QuoRem(n, i, r)

		if mod.Cmp(zero) != 0 {
			continue
		}

		allFacs = append(allFacs, big.NewInt(0).Set(i))
		allFacs = append(allFacs, big.NewInt(0).Set(quotient))

		if isPrime(i) {
			primeFacs = append(primeFacs, big.NewInt(0).Set(i))
		} else {
			nonPrimeFacs = append(nonPrimeFacs, big.NewInt(0).Set(i))
		}

		if isPrime(quotient) {
			primeFacs = append(primeFacs, big.NewInt(0).Set(quotient))
		} else {
			nonPrimeFacs = append(nonPrimeFacs, big.NewInt(0).Set(quotient))
		}
	}

	return allFacs, nonPrimeFacs, primeFacs
}

func isPrime(n *big.Int) bool {

	temp := big.NewInt(0)
	increment := big.NewInt(1)
	zero := big.NewInt(0)
	r := big.NewInt(0)
	mod := big.NewInt(0)

	for i := big.NewInt(2); i.Cmp(n) < 0; i.Add(i, increment) {

		_, mod = temp.QuoRem(n, i, r)

		if mod.Cmp(zero) == 0 {
			return false
		}
	}

	return true
}
