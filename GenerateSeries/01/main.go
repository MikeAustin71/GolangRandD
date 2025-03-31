package main

import (
	"fmt"
	"math/big"
)

func main() {

	upperLimit := big.NewInt(10)
	series := generateSequence(upperLimit)
	fmt.Println("This is the series with upper limit = ", upperLimit)
	fmt.Println(series)
}

func generateSequence(upperLimit *big.Int) []*big.Int {

	series := make([]*big.Int, upperLimit.Int64())
	aryLimit := upperLimit.Int64()
	for i := int64(0); i < aryLimit; i++ {

		series[i] = big.NewInt(i + int64(1))
	}

	return series
}
