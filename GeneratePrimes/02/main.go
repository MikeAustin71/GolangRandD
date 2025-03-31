package main

import (
	"fmt"
	"math/big"
)

func main() {
	upperLimit := int64(10)
	primes := generatePrimes(upperLimit)
	fmt.Println("Primes of arry upper limit = ", upperLimit)
	fmt.Println(primes)
}

func generatePrimes(upperLimit int64) []*big.Int {
	var primes []*big.Int
	var isPrime bool
	var j int
	aryLimit := big.NewInt(upperLimit)
	primes = append(primes, big.NewInt(2))
	increment := big.NewInt(2)
	zero := big.NewInt(0)

	for i := big.NewInt(3); i.Cmp(aryLimit) <= 0; i.Add(i, increment) {
		j = 0
		isPrime = true
		y := big.NewInt(1)

		for y = y.Mul(primes[j], primes[j]); y.Cmp(i) <= 0 && j < len(primes); j++ {
			modulo := big.NewInt(1)
			modulo = modulo.Mod(i, primes[j])
			if modulo.Cmp(zero) == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			prime := big.NewInt(0)
			prime = prime.Add(prime, i)
			primes = append(primes, prime)
		}
	}

	return primes
}

/*
func generateSequence(upperLimit int64) []*big.Int {

	bintArray := make([]*big.Int, upperLimit)

	for i := int64(1); i <= upperLimit; i++ {
		bintArray[i-1] = big.NewInt(i)
	}

	return bintArray
}
*/
