package main

import (
	"fmt"
	"math/big"
)

func main()  {

	base10 := big.NewInt(10)

	exponent := big.NewInt(40)

	scaleFactor := big.NewInt(0).Exp(base10, exponent, nil)

	numstr := "1234567890123456789012345678901234567890"

	bNum, isOk := big.NewInt(0).SetString(numstr, 10)

	if !isOk {
		fmt.Printf("Big Int SetString failed to convert numstr. numstr='%v' \n", numstr)
		return
	}


	rDividend := big.NewRat(1, 1).SetInt(bNum)
	rDivisor := big.NewRat(1, 1).SetInt(scaleFactor)
	rQuotient := big.NewRat(1,1).Quo(rDividend,rDivisor)
	fmt.Println("         numstr: ", numstr)
	fmt.Println("  bNum.Text(10): ", bNum.Text(10))
	fmt.Println("           bNum: ", bNum.String())
	fmt.Println("   Scale Factor: ", scaleFactor)
	fmt.Println("      rQuotient: ", rQuotient.FloatString(40))
	bFloat := big.NewFloat(0.0).SetRat(rQuotient)
	fmt.Println("  bFloat String: ", bFloat.String())
	fmt.Println("bFloat Text(40): ", bFloat.Text('f', 40))

	bFloat.SetPrec(40)
	fmt.Println("After Setting Precision = 40")
	fmt.Println("  bFloat String: ", bFloat.String())
	fmt.Println("bFloat Text(40): ", bFloat.Text('g', 40))
	accuracy := bFloat.Acc()
	fmt.Println("bFloat Accuracy: ", accuracy)





	fmt.Println()
}

