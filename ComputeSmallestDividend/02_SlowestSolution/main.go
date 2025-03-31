package main

import (
	"fmt"
	"math/big"
)

// Using brute force approach with math/big and big.Int. For series with
// upper limit greater than 20, this is very slow.

func main() {

	seriesMax := big.NewInt(20)

	series := GenerateSeries(seriesMax)

	smallestDividend, success := FindSmallestDividend(series)

	if success {
		fmt.Println("Success. The smallest dividend evenly divisible by the series 1 to ", seriesMax, " is - ", smallestDividend)

	} else {
		fmt.Println("Failure - Did Not Locate smallest dividend.")
	}

}

func GenerateSeries(upperLimit *big.Int) []*big.Int {

	series := make([]*big.Int, upperLimit.Int64())
	aryLimit := upperLimit.Int64()
	for i := int64(0); i < aryLimit; i++ {

		series[i] = big.NewInt(i + int64(1))
	}

	return series
}

func FindSmallestDividend(series []*big.Int) (*big.Int, bool) {

	zero := big.NewInt(0)

	if len(series) == 0 {
		return zero, false
	}

	increment := big.NewInt(series[len(series)-1].Int64())
	var limit = big.NewInt(10).Exp(big.NewInt(10), big.NewInt(99), nil)

	for i := big.NewInt(0).Set(series[len(series)-1]); i.Cmp(limit) <= 0; i.Add(i, increment) {

		if IsTestNumDivisibleBySeries(series, i) {
			return i, true
		}
	}

	return zero, false
}

func IsTestNumDivisibleBySeries(series []*big.Int, testNum *big.Int) bool {

	zero := big.NewInt(0)

	for _, v := range series {

		if big.NewInt(0).Mod(testNum, v).Cmp(zero) != 0 {
			return false
		}

	}

	return true
}

/*

* Source: https://projecteuler.net/problem=5

* Title: Smallest multiple - Problem 5

* Problem Description

-----------------------------------------------------------------------
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder. What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
-----------------------------------------------------------------------
Success, Correct Answer = 232792560 - Confirmed by Project Euler

*/
