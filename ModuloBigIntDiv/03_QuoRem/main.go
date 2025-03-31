package main

import (
	"fmt"
	"math/big"
)

func main() {

	dividend := big.NewInt(13)
	divisor := big.NewInt(2)
	quotient, mod, r := QuoRem(dividend, divisor)

	fmt.Println("Dividend:", dividend)
	fmt.Println("Divisor :", divisor)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Mod     :", mod)
	fmt.Println("r       :", r)

}

func QuoRem(dividend *big.Int, divisor *big.Int) (quotient *big.Int, mod *big.Int, r *big.Int) {

	quotient = big.NewInt(0)
	mod = big.NewInt(0)
	r = big.NewInt(0)
	temp := big.NewInt(0)

	quotient, mod = temp.QuoRem(dividend, divisor, r)

	return quotient, mod, r
}
