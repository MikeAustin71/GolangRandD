package main

import (
	"fmt"
	"math/big"
)

func main() {
	series := getSeries()

	fmt.Println(series)
}

func getSeries() []big.Int {

	MAXSERIESNUM := 10

	series := make([]big.Int, MAXSERIESNUM-1)

	for i := 2; i <= MAXSERIESNUM; i++ {
		// x := big.NewInt(int64(i))
		// series[i-2] = *x
		series[i-2] = *big.NewInt(int64(i))
	}

	return series
}
