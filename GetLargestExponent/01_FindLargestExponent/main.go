package main

import (
	"fmt"
	"math/big"
)

func main() {

	base := big.NewInt(2)
	dividend := big.NewInt(15)
	exp := findLargestExponent(base, dividend)
	fmt.Println("----------------------------------------------")
	fmt.Println("For base: ", base, " and max dividend: ", dividend)
	fmt.Println("Largest exponent is :", exp)
	value := big.NewInt(1)
	value = value.Exp(base, exp, nil)
	fmt.Println("Largest value base^exponent :", value)

}

func findLargestExponent(base *big.Int, dividend *big.Int) *big.Int {

	exp := big.NewInt(0)
	increment := big.NewInt(1)

	for value := big.NewInt(0).Set(base); value.Cmp(dividend) <= 0; exp.Add(exp, increment) {

		value = value.Mul(value, base)

	}

	return exp
}
