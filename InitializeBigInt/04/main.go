package main

import (
	"fmt"
	"math/big"
)

func main() {
	series := getSeries(6)

	fmt.Println(series)
}

func getSeries(seriesSize int) []big.Int {

	series := make([]big.Int, seriesSize-1)

	for i := 2; i <= seriesSize; i++ {
		// x := big.NewInt(int64(i))
		// series[i-2] = *x
		series[i-2] = *big.NewInt(int64(i))
	}

	return series
}
