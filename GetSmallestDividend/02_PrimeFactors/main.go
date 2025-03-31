package main

import (
	"fmt"
	"math/big"
)

func main() {

	upperLimit := big.NewInt(20)
	series := generateSequence(upperLimit)
	primes := generatePrimes(upperLimit)
	primeExps := calculateMaxExponentsForPrimeSeries(primes, upperLimit)

	fmt.Println("----------------------------------------------")
	fmt.Println("Number series 1 - ", upperLimit)

	fmt.Println("Primes Factors and Prime Exponents --------------------")
	for i := 0; i < len(primes); i++ {
		fmt.Println("Prime: ", primes[i],
			"  Prime Max Exponent: ", primeExps[i])
	}
	fmt.Println("--------------------------------------------")

	result := computeSmallestDividend(primes, primeExps)

	fmt.Println("Smallest Dividend Is: ", result)

	if isGood := IsResultDivisibleBySeries(series, result); isGood {
		fmt.Println("Result Confirmed ! Result is divisible by each number in the series")
	} else {
		fmt.Print("Failure !!!! - Result is NOT divisible by each number in the series")
	}

}

func generateSequence(upperLimit *big.Int) []*big.Int {

	series := make([]*big.Int, upperLimit.Int64())
	aryLimit := upperLimit.Int64()
	for i := int64(0); i < aryLimit; i++ {

		series[i] = big.NewInt(i + int64(1))
	}

	return series
}

func generatePrimes(upperLimit *big.Int) []*big.Int {
	var primes []*big.Int
	var isPrime bool
	var j int
	aryLimit := upperLimit
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
	upperLimit *big.Int) []*big.Int {

	primeExps := make([]*big.Int, len(primes))
	limit := len(primes)
	for i := 0; i < limit; i++ {

		primeExps[i] = findLargestExponent(primes[i], upperLimit)

	}

	return primeExps
}

func findLargestExponent(base *big.Int, upperLimit *big.Int) *big.Int {

	exp := big.NewInt(0)
	increment := big.NewInt(1)

	for value := big.NewInt(0).Set(base); value.Cmp(upperLimit) <= 0; exp.Add(exp, increment) {

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

func IsResultDivisibleBySeries(series []*big.Int, testNum *big.Int) bool {

	zero := big.NewInt(0)

	for _, v := range series {

		dividend := big.NewInt(0)
		modulo := dividend.Mod(testNum, v)

		if modulo.Cmp(zero) != 0 {
			return false
		}

	}

	return true
}
