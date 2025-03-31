package main

import (
	"fmt"
	"math/big"
)

func main()  {

	fmt.Println()
	numStr := "123456654321"
	precision := 6
	signedAllDigits, isOk := big.NewInt(0).SetString(numStr, 10)

	if !isOk {
		fmt.Printf("Error from big.NewInt(0).SetString(numStr, 10). numStr='%v'", numStr)
		return
	}

	base10 := big.NewInt(10)
	bigPrecision := big.NewInt(int64(precision))

	scaleFactor := big.NewInt(0).Exp(base10,bigPrecision, nil)

	fmt.Println("signedAllDigits: ", signedAllDigits.Text(10))
	fmt.Println("     scalFactor: ", scaleFactor.Text(10))

	modM := big.NewInt(0)

	s2, s2Mod := big.NewInt(0).DivMod(signedAllDigits, scaleFactor, modM)
	fmt.Println("modM Value after Division: ", modM.Text(10))
	str := s2.Text(10)
	str += "."
	str += s2Mod.Text(10)

	fmt.Println("original numStr: ", numStr)
	fmt.Println("      precision: ", precision)
	fmt.Println("  Output String: ", str)




}

