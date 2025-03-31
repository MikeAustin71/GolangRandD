package main

import (
	"fmt"
	"math/big"
)

func main()  {

	fmt.Println()
	bigPrecision := big.NewInt(int64(3))
	base10 := big.NewInt(int64(10))

	dividend := big.NewInt(int64(123456789))
	divisor := big.NewInt(0).Exp(base10, bigPrecision, nil)
	quotient, mod, modulus := DivideMod(dividend, divisor)

	PrintDivResult(dividend, divisor, quotient, mod, modulus)


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



