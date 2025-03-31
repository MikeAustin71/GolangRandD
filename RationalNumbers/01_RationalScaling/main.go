package main

import (
	"fmt"
	"math/big"
)

func main()  {

	base10 := big.NewInt(10)

	exponent := big.NewInt(15)

	scaleFactor := big.NewInt(0).Exp(base10, exponent, nil)

	numstr := "12345678901234567890"

	bNum, isOk := big.NewInt(0).SetString(numstr, 10)

	if !isOk {
		fmt.Printf("Big Int SetString failed to convert numstr. numstr='%v' \n", numstr)
		return
	}


	rDividend := big.NewRat(1, 1).SetInt(bNum)
	rDivisor := big.NewRat(1, 1).SetInt(scaleFactor)
	rQuotient := big.NewRat(1,1).Quo(rDividend,rDivisor)
	fmt.Println("          bNum: ", bNum.String())
	fmt.Println("  Scale Factor: ", scaleFactor)
	fmt.Println("     rQuotient: ", rQuotient.FloatString(42))

	f20str := rQuotient.FloatString(20)
	bFloat, isOk := big.NewFloat(0.0).SetString(f20str)

	if !isOk {
		fmt.Printf("Big Float SetString failed to convert f20str. f20str='%v' \n", f20str)
		return
	}

	fmt.Println("bFloat  String: ", bFloat.String())
	fmt.Println("bFloat Text-16: ", bFloat.Text('f', 16))
}

