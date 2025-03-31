package main

import (
	"fmt"
	"math/big"
)

/* Project Euler - Problem 3
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
*/

// Brute force solution

func main() {
	//n := big.NewInt(600851475143)
	n := big.NewInt(12)
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
	largestFac := big.NewInt(0)
	var factors []*big.Int
	factors = append(factors, big.NewInt(0))
	factors = append(factors, big.NewInt(0))

	var allFactors []*big.Int
	var primeFacs []*big.Int
	var nonPrimeFacs []*big.Int
	temp := big.NewInt(1)
	increment := big.NewInt(1)
	mod := big.NewInt(0)
	zero := big.NewInt(0)

	for i := big.NewInt(2); temp.Mul(i, i).Cmp(n) < 0; i.Add(i, increment) {

		if mod.Mod(n, i).Cmp(zero) == 0 { // It is a divisor
			factors[0] = big.NewInt(1).Set(i)
			allFactors = append(allFactors, big.NewInt(1).Set(i))
			factors[1] = big.NewInt(1).Div(n, i)
			allFactors = append(allFactors, big.NewInt(1).Div(n, i))

			for k := 0; k < 2; k++ {
				isPrime := true

				for j := big.NewInt(2); temp.Mul(j, j).Cmp(factors[k]) < 0; j.Add(j, increment) {
					if mod.Mod(factors[k], j).Cmp(zero) == 0 {
						isPrime = false
						nonPrimeFacs = append(nonPrimeFacs, factors[k])
						break
					}
				}

				if isPrime {
					primeFacs = append(primeFacs, factors[k])
				}

				if isPrime && factors[k].Cmp(largestFac) > 0 {
					largestFac.Set(factors[k])
				}
			}
		}
	}

	return largestFac, allFactors, nonPrimeFacs, primeFacs
}
