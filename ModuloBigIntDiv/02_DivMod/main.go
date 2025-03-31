package main

import (
	"fmt"
	"math/big"
)

func main() {

	dividend := big.NewInt(15)
	divisor := big.NewInt(4)
	quotient, mod, modulus := DivideMod(dividend, divisor)

	fmt.Println("Dividend:", dividend)
	fmt.Println("Divisor :", divisor)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Mod     :", mod)
	fmt.Println("Modulus :", modulus)

}

func DivideMod(dividend *big.Int, divisor *big.Int) (quotient *big.Int, mod *big.Int, modulus *big.Int) {

	quotient = big.NewInt(0)
	mod = big.NewInt(0)
	modulus = big.NewInt(0)

	quotient, mod = big.NewInt(0).DivMod(dividend, divisor, modulus)

	return quotient, mod, modulus
}
