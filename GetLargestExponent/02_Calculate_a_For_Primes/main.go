package main

import (
	"fmt"
	"math/big"
)

func main() {

	upperLimit := int64(3025)

	fmt.Println("----------------------------------------------")
	fmt.Println("Number series 1 - ", upperLimit)

	primes := generatePrimes(upperLimit)
	series := generateNumberSeries(upperLimit)
	primeexponents := calculateMaxExponentsForPrimeSeries(primes, series)

	fmt.Println("Primes -------------------------------------")
	for i := 0; i < len(primes); i++ {
		fmt.Println("Prime: ", primes[i],
			"  Prime Max Exponent: ", primeexponents[i])
	}
	fmt.Println("--------------------------------------------")
	largestPrimeIdx := findMaxValueIndexInSeries(primes)
	base := primes[largestPrimeIdx]
	exp := primeexponents[largestPrimeIdx]
	fmt.Println("Largest Prime (base):", base)
	fmt.Println("base exponent :", exp)
	value := big.NewInt(1)
	value = value.Exp(base, exp, nil)
	fmt.Println("Largest value base^exponent :", value)
	smallestDividend := computeSmallestDividend(primes, primeexponents)
	fmt.Println("--------------------------------------------")
	fmt.Println("Smallest Dividend:", smallestDividend)
}

func generateNumberSeries(upperLimit int64) []*big.Int {

	bintArray := make([]*big.Int, upperLimit)

	for i := int64(1); i <= upperLimit; i++ {
		bintArray[i-1] = big.NewInt(i)
	}

	return bintArray
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

func calculateMaxExponentsForPrimeSeries(primes []*big.Int,
	series []*big.Int) []*big.Int {

	upperLimit := findMaximumValueInSeries(series)
	primeExps := make([]*big.Int, len(primes))
	limit := len(primes)
	for i := 0; i < limit; i++ {

		primeExps[i] = findLargestExponent(primes[i], upperLimit)

	}

	return primeExps
}

func findMaxValueIndexInSeries(series []*big.Int) int {

	limit := len(series)
	maxVal := big.NewInt(0)
	maxIdx := 0

	for i := 0; i < limit; i++ {

		if series[i].Cmp(maxVal) > 0 {
			maxVal.Set(big.NewInt(0).Set(series[i]))
			maxIdx = i
		}

	}

	return maxIdx
}

func findMaximumValueInSeries(series []*big.Int) *big.Int {

	limit := len(series)
	maxVal := big.NewInt(0)

	for i := 0; i < limit; i++ {

		if series[i].Cmp(maxVal) > 0 {
			maxVal.Set(big.NewInt(0).Set(series[i]))
		}

	}

	return maxVal
}

func findLargestExponent(base *big.Int, dividend *big.Int) *big.Int {

	exp := big.NewInt(0)
	increment := big.NewInt(1)

	for value := big.NewInt(0).Set(base); value.Cmp(dividend) <= 0; exp.Add(exp, increment) {

		value = value.Mul(value, base)

	}

	return exp
}

func computeSmallestDividend(primes []*big.Int,
	primeExponents []*big.Int) *big.Int {

	result := big.NewInt(1)
	limit := len(primes)
	num := big.NewInt(1)

	for i := 0; i < limit; i++ {

		result = result.Mul(result, num.Exp(primes[i], primeExponents[i], nil))
	}

	return result
}
