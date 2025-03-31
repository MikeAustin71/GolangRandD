package main

import (
	"fmt"
	"math/big"
)

// Playing with math/big

func main() {
	series := getSeries(20)

	fmt.Println(series)

	smallestDividend, success := findSmallestDividend(series)

	/*
		if success {
			fmt.Println("Success. The smallest dividend evenly divisible by the series 1 to 20 is - ", smallestDividend)

		} else {
			fmt.Println("Failure - Did Not Locate smallest dividend.")
		}
	*/
	fmt.Println(smallestDividend, "  ", success)
}

func getSeries(seriesSize int) []big.Int {

	series := make([]big.Int, seriesSize-1)

	for i := 2; i <= seriesSize; i++ {
		series[i-2] = *big.NewInt(int64(i))
	}

	return series
}

func findSmallestDividend(series []big.Int) (big.Int, bool) {

	zero := big.NewInt(0)

	if len(series) == 0 {
		return *zero, false
	}

	increment := big.NewInt(1)
	var limit = big.NewInt(30)
	// limit.Exp(big.NewInt(10), big.NewInt(99), nil)

	for i := series[len(series)-1]; i.Cmp(limit) < 0; i.Add(&i, increment) {

		fmt.Println("in loop, this is i ", i)
	}

	return *zero, false
}

func isTestNumDivisbleBySeries(series []big.Int, testNum *big.Int) bool {

	dividend := big.NewInt(0)
	zero := big.NewInt(0)

	for _, v := range series {

		modulo := dividend.Mod(testNum, &v)

		if modulo.Cmp(zero) != 0 {
			return false
		}

	}

	return true
}
