package main

import (
	"fmt"
	"math/big"
)

func main() {
	var series [4]big.Int
	a := big.NewInt(-1)
	b := big.NewInt(3)
	c := big.NewInt(5)
	d := big.NewInt(9)

	series[0] = *a
	series[1] = *b
	series[2] = *c
	series[3] = *d

	fmt.Println(series)
}

/*
Printed Output
[{true [1]} {false [3]} {false [5]} {false [9]}]
*/
