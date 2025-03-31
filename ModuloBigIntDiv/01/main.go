package main

import (
	"fmt"
	"math/big"
)

func main() {

	increment := big.NewInt(1)
	limit := big.NewInt(100)

	for i := big.NewInt(1); i.Cmp(limit) <= 0; i.Add(i, increment) {

		if modulo, isEven := isModEvenBigNum(i); isEven {
			fmt.Println("The value of modulo is ", modulo, " ", i, " is an even number!")
		} else {
			fmt.Println("The value of modulo is ", modulo, " ", i, " is an odd number!")
		}

	}
}

func isModEvenBigNum(dividend *big.Int) (*big.Int, bool) {

	modulo := big.NewInt(0)
	divisor := big.NewInt(2)
	modValue := big.NewInt(0)

	modulo.Mod(dividend, divisor)

	if modulo.Cmp(modValue) == 0 {
		return modulo, true
	}

	return modulo, false
}
