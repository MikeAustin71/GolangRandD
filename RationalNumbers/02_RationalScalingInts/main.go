package main

import (
	"fmt"
	"math/big"
)

func main()  {

	fmt.Println()

	dividend := big.NewInt(int64(15123))
	divisor := big.NewInt(int64(10000))



	rDividend := big.NewRat(1, 1).SetInt(dividend)
	rDivisor := big.NewRat(1, 1).SetInt(divisor)
	rQuotient := big.NewRat(1,1).Quo(rDividend,rDivisor)

	fmt.Println("RQuotient Float String: ", rQuotient.FloatString(10))

}

