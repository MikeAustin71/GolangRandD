package main

import (
	"fmt"
	"math/big"
)

func main()  {

	numStr:= "123456.07890"
	absNumStr := "12345607890"
	precision := uint(5)
	bigPrecision := big.NewInt(int64(precision))
	base10 := big.NewInt(int64(10))
	scaleFactor := big.NewInt(0).Exp(base10, bigPrecision, nil)
	absBigInt, _ := big.NewInt(0).SetString(absNumStr, 10)
	rationalNum := big.NewRat(1,1).SetFrac(absBigInt,scaleFactor)

	ratNumerator := rationalNum.Num()
	ratDenominator := rationalNum.Denom()

	fmt.Println("      absNumStr: ", absNumStr)
	fmt.Println("   ratNumerator: ", ratNumerator.String())
	fmt.Println("---------------------------------------------------------")
	fmt.Println("    scaleFactor: ", scaleFactor.String())
	fmt.Println(" ratDenominator: ", ratDenominator.String())
	fmt.Println("---------------------------------------------------------")
	fmt.Println("         numStr: ", numStr)
	fmt.Println("      precision: ", precision)
	fmt.Println( "   FloatString: ", rationalNum.FloatString(int(precision)))

}
