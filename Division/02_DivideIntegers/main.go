package main

import (
	"fmt"
	"math/big"
)

func main()  {

	Ex01()

}

func Ex01() {
	/*
	dividend := big.NewInt(int64(15123))
	divisor := big.NewInt(int64(4800))
	*/
	dividend := big.NewInt(int64(-15123))
	divisor := big.NewInt(int64(-4800))
	quotient, mod, modulus := DivideMod(dividend, divisor)

	PrintDivResult(dividend, divisor, quotient, mod, modulus)

	dividend = big.NewInt(0).Set(mod)
	quotient, mod, modulus = DivideMod(dividend, divisor)
	PrintDivResult(dividend, divisor, quotient, mod, modulus)

}

func PrintDivResult (dividend *big.Int, divisor *big.Int, quotient *big.Int, mod *big.Int, modulus *big.Int ){
	fmt.Println()
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
