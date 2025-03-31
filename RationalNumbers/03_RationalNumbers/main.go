package main

import (
	"fmt"
	"math/big"
)

func main()  {

	fmt.Println()

	dividend := big.NewInt(int64(15123))
	divisor := big.NewInt(int64(4800))
	quotient, mod, modulus := DivideMod(dividend, divisor)

	PrintDivResult(dividend, divisor, quotient, mod, modulus)

	rDividend := big.NewRat(1, 1).SetInt(dividend)
	rDivisor := big.NewRat(1, 1).SetInt(divisor)
	rQuotient := big.NewRat(1,1).Quo(rDividend,rDivisor)

	fmt.Println("RQuotient Float String: ", rQuotient.FloatString(10))

}

func PrintDivResult (dividend *big.Int, divisor *big.Int, quotient *big.Int, mod *big.Int, modulus *big.Int ){
	fmt.Println("Dividend:", dividend)
	fmt.Println("Divisor :", divisor)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Mod     :", mod)
	fmt.Println("Modulus :", modulus)

}

func DivideMod(dividend *big.Int, divisor *big.Int) (quotient *big.Int, mod *big.Int, modulus *big.Int) {

	modulus = big.NewInt(0)

	quotient, mod = big.NewInt(0).DivMod(dividend, divisor, modulus)

	return quotient, mod, modulus
}
